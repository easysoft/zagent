package hostRouter

import (
	hostHandler "github.com/easysoft/zagent/cmd/agent-host/router/handler"
	"github.com/easysoft/zagent/cmd/agent/router/handler"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api *iris.Application

	ArithCtrl *handler.ArithCtrl   `inject:""`
	JobCtrl   *hostHandler.JobCtrl `inject:""`

	KvmCtrl    *hostHandler.KvmCtrl    `inject:""`
	VmWareCtrl *hostHandler.VmWareCtrl `inject:""`
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
			v1.PartyFunc("/kvm", func(client iris.Party) {
				client.Post("/create", r.KvmCtrl.Create).Name = "创建虚机"
				client.Post("/destroy", r.KvmCtrl.Destroy).Name = "摧毁虚机"
				client.Post("/boot", r.KvmCtrl.Boot).Name = "启动虚机"
				client.Post("/shutdown", r.KvmCtrl.Shutdown).Name = "关闭虚机"
				client.Post("/reboot", r.KvmCtrl.Reboot).Name = "重启虚机"
				client.Post("/suspend", r.KvmCtrl.Suspend).Name = "暂停虚机"
				client.Post("/resume", r.KvmCtrl.Resume).Name = "恢复虚机"

			})
			v1.PartyFunc("/vmware", func(client iris.Party) {
				client.Post("/create", r.VmWareCtrl.Create).Name = "创建虚机"
				client.Post("/destroy", r.VmWareCtrl.Destroy).Name = "摧毁虚机"
			})
		}
	}
}
