package vmRouter

import (
	vmHandler "github.com/easysoft/zagent/cmd/agent-vm/router/handler"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api *iris.Application

	JobCtrl *vmHandler.JobCtrl `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{api: app}
	return router
}

func (r *Router) App() {
	iris.LimitRequestBodySize(serverConf.Inst.Options.UploadMaxSize)
	r.api.UseRouter(_httpUtils.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		v1 := app.Party("/v1")
		{
			v1.PartyFunc("/job", func(client iris.Party) {
				client.Post("/add", r.JobCtrl.Add).Name = "创建任务"
			})
		}
	}
}
