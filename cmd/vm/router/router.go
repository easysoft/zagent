package vmRouter

import (
	vmHandler "github.com/easysoft/zv/cmd/vm/router/handler"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api      *iris.Application
	TestCtrl *vmHandler.TestCtrl `inject:""`

	JobCtrl *vmHandler.JobCtrl `inject:""`
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
			//v1.Use(core.Auth())

			v1.Get("/test", r.TestCtrl.Test).Name = "测试"

			v1.PartyFunc("/job", func(client iris.Party) {
				client.Post("/add", r.JobCtrl.Add).Name = "创建任务"
			})
		}
	}
}
