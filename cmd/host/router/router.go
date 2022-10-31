package hostRouter

import (
	hostHandler "github.com/easysoft/zagent/cmd/host/router/handler"
	hostService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api *iris.Application

	TestCtrl     *hostHandler.TestCtrl     `inject:""`
	TaskCtrl     *hostHandler.TaskCtrl     `inject:""`
	DownloadCtrl *hostHandler.DownloadCtrl `inject:""`
	CheckCtrl    *hostHandler.ServiceCtrl  `inject:""`
	JobCtrl      *hostHandler.JobCtrl      `inject:""`

	KvmCtrl        *hostHandler.KvmCtrl        `inject:""`
	VirtualBoxCtrl *hostHandler.VirtualBoxCtrl `inject:""`
	VmWareCtrl     *hostHandler.VmWareCtrl     `inject:""`
	VirtualCtrl    *hostHandler.VirtualCtrl    `inject:""`
	MultiPassCtrl  *hostHandler.MultiPassCtrl  `inject:""`

	InitService *hostService.InitService `inject:""`
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

			v1.PartyFunc("/service", func(client iris.Party) {
				client.Post("/check", r.CheckCtrl.CheckService).Name = "检测宿主机服务状态"
			})

			v1.PartyFunc("/task", func(client iris.Party) {
				client.Post("/getStatus", r.TaskCtrl.GetStatus).Name = "获取任务状态"
			})

			v1.PartyFunc("/download", func(client iris.Party) {
				client.Post("/add", r.DownloadCtrl.Add).Name = "添加下载任务"
				client.Post("/cancel", r.DownloadCtrl.Cancel).Name = "强制终止下载任务"
			})

			//v1.PartyFunc("/job", func(client iris.Party) {
			//	client.Post("/add", r.JobCtrl.Add).Name = "创建任务"
			//})

			v1.PartyFunc("/virtual", func(client iris.Party) {
				client.Post("/notifyHost", r.VirtualCtrl.NotifyHost).Name = "虚拟机请求安全码"
				client.Get("/getVncToken", r.VirtualCtrl.GetToken).Name = "根据VNC Port获取Token"

				client.Post("/addVmPortMap", r.VirtualCtrl.AddVmPortMap).Name = "新增虚拟机到宿主机端口的映射"
				client.Post("/removeVmPortMap", r.VirtualCtrl.RemoveVmPortMap).Name = "移除虚拟机到宿主机的端口映射"
			})

			v1.PartyFunc("/kvm", func(client iris.Party) {
				//client.Get("/listTmpl", r.KvmCtrl.ListTmpl).Name = "列出KVM虚拟机镜像"
				client.Post("/create", r.KvmCtrl.Create).Name = "创建KVM虚拟机"
				client.Post("/clone", r.KvmCtrl.Clone).Name = "克隆KVM虚拟机"

				client.Post("/{name:string}/destroy", r.KvmCtrl.Destroy).Name = "摧毁KVM虚拟机"
				client.Post("/{name:string}/reboot", r.KvmCtrl.Reboot).Name = "重启KVM虚拟机"
				client.Post("/{name:string}/suspend", r.KvmCtrl.Suspend).Name = "暂停KVM虚拟机"
				client.Post("/{name:string}/resume", r.KvmCtrl.Resume).Name = "恢复KVM虚拟机"

				client.Post("/exportVm", r.KvmCtrl.ExportVm).Name = "导出KVM虚拟机为模板镜像"
			})

			v1.PartyFunc("/virtualbox", func(client iris.Party) {
				client.Post("/listTmpl", r.VirtualBoxCtrl.ListTmpl).Name = "列出VirtualBox镜像"
				client.Post("/create", r.VirtualBoxCtrl.Create).Name = "创建VirtualBox虚拟机"
				client.Post("/destroy", r.VirtualBoxCtrl.Destroy).Name = "摧毁VirtualBox虚拟机"
			})

			v1.PartyFunc("/vmware", func(client iris.Party) {
				client.Post("/create", r.VmWareCtrl.Create).Name = "创建VMWare虚拟机"
				client.Post("/destroy", r.VmWareCtrl.Destroy).Name = "摧毁VMWare虚拟机"
			})

			v1.PartyFunc("/multipass", func(client iris.Party) {
				client.Post("/create", r.MultiPassCtrl.Create).Name = "创建Multipass虚拟机"
				client.Post("/{name:string}/reboot", r.MultiPassCtrl.Reboot).Name = "重启Multipass虚拟机"
				client.Post("/{name:string}/destroy", r.MultiPassCtrl.Destroy).Name = "摧毁Multipass虚拟机"
				client.Post("/{name:string}/suspend", r.MultiPassCtrl.Suspend).Name = "暂停Multipass虚拟机"
				client.Post("/{name:string}/resume", r.MultiPassCtrl.Resume).Name = "恢复Multipass虚拟机"
				client.Get("/getToken", r.MultiPassCtrl.GetToken).Name = "获取VNC的Token"
			})
		}
	}
}
