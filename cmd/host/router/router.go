package hostRouter

import (
	hostHandler "github.com/easysoft/zv/cmd/host/router/handler"
	serverConf "github.com/easysoft/zv/internal/server/conf"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api *iris.Application

	TestCtrl     *hostHandler.TestCtrl     `inject:""`
	DownloadCtrl *hostHandler.DownloadCtrl `inject:""`
	CheckCtrl    *hostHandler.ServiceCtrl  `inject:""`
	JobCtrl      *hostHandler.JobCtrl      `inject:""`

	KvmCtrl        *hostHandler.KvmCtrl        `inject:""`
	VirtualBoxCtrl *hostHandler.VirtualBoxCtrl `inject:""`
	VmWareCtrl     *hostHandler.VmWareCtrl     `inject:""`
	VirtualCtrl    *hostHandler.VirtualCtrl    `inject:""`
	MultiPassCtrl  *hostHandler.MultiPassCtrl  `inject:""`
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
			//v1.Use(core.Auth())

			v1.PartyFunc("/resource", func(client iris.Party) {
				client.Post("/download", r.DownloadCtrl.Download).Name = "下载文件"
			})

			v1.PartyFunc("/service", func(client iris.Party) {
				client.Post("/check", r.CheckCtrl.CheckStatus).Name = "检测服务状态"
			})

			v1.PartyFunc("/job", func(client iris.Party) {
				client.Post("/add", r.JobCtrl.Add).Name = "创建任务"
			})

			v1.PartyFunc("/virtual", func(client iris.Party) {
				client.Post("/notifyHost", r.VirtualCtrl.NotifyHost).Name = "虚拟机用Mac地址请求安全码"

				client.Post("/addVmPortMap", r.VirtualCtrl.AddVmPortMap).Name = "映射虚机端口到宿主机"
				client.Post("/removeVmPortMap", r.VirtualCtrl.RemoveVmPortMap).Name = "映射虚机端口到宿主机"
				client.Get("/getVncToken", r.VirtualCtrl.GetToken).Name = "获取VNC的Token"
			})

			v1.PartyFunc("/kvm", func(client iris.Party) {
				client.Get("/listTmpl", r.KvmCtrl.ListTmpl).Name = "克隆虚机"
				client.Post("/create", r.KvmCtrl.Create).Name = "创建虚机"
				client.Post("/clone", r.KvmCtrl.Clone).Name = "克隆虚机"
				client.Post("/{name:string}/destroy", r.KvmCtrl.Destroy).Name = "摧毁虚机"
				client.Post("/{name:string}/reboot", r.KvmCtrl.Reboot).Name = "重启虚机"
				client.Post("/{name:string}/suspend", r.KvmCtrl.Suspend).Name = "暂停虚机"
				client.Post("/{name:string}/resume", r.KvmCtrl.Resume).Name = "恢复虚机"
			})

			v1.PartyFunc("/virtualbox", func(client iris.Party) {
				client.Post("/listTmpl", r.VirtualBoxCtrl.ListTmpl).Name = "摧毁虚机"
				client.Post("/create", r.VirtualBoxCtrl.Create).Name = "创建虚机"
				client.Post("/destroy", r.VirtualBoxCtrl.Destroy).Name = "摧毁虚机"
			})

			v1.PartyFunc("/vmware", func(client iris.Party) {
				client.Post("/create", r.VmWareCtrl.Create).Name = "创建虚机"
				client.Post("/destroy", r.VmWareCtrl.Destroy).Name = "摧毁虚机"
			})

			v1.PartyFunc("/multipass", func(client iris.Party) {
				client.Post("/create", r.MultiPassCtrl.Create).Name = "创建虚机"
				client.Post("/{name:string}/reboot", r.MultiPassCtrl.Reboot).Name = "重启虚机"
				client.Post("/{name:string}/destroy", r.MultiPassCtrl.Destroy).Name = "摧毁虚机"
				client.Post("/{name:string}/suspend", r.MultiPassCtrl.Suspend).Name = "暂停虚机"
				client.Post("/{name:string}/resume", r.MultiPassCtrl.Resume).Name = "恢复虚机"
				client.Get("/getToken", r.MultiPassCtrl.GetToken).Name = "获取VNC的Token"
			})
		}
	}
}
