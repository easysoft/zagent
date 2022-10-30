package vmRouter

import (
	vmHandler "github.com/easysoft/zagent/cmd/vm/router/handler"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
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
