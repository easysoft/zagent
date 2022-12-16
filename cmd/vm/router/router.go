package vmRouter

import (
	vmHandler "github.com/easysoft/zagent/cmd/vm/router/handler"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api         *iris.Application
	ServiceCtrl *vmHandler.ServiceCtrl `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{api: app}
	return router
}

func (r *Router) App() {
	iris.LimitRequestBodySize(consts.UploadMaxSize)
	r.api.UseRouter(_httpUtils.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		v1 := app.Party("/v1")
		{
			// v1.Use(core.Auth())

			v1.PartyFunc("/service", func(client iris.Party) {
				client.Post("/check", r.ServiceCtrl.Check).Name = "检测节点服务状态"
				client.Post("/setup", r.ServiceCtrl.Setup).Name = "安装测试服务"
			})
		}
	}
}
