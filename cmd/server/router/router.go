package router

import (
	"fmt"
	"github.com/easysoft/zv/cmd/server/router/handler"
	_httpUtils "github.com/easysoft/zv/internal/pkg/lib/http"
	bizCasbin "github.com/easysoft/zv/internal/server/biz/casbin"
	"github.com/easysoft/zv/internal/server/biz/jwt"
	serverConf "github.com/easysoft/zv/internal/server/conf"
	"github.com/easysoft/zv/internal/server/repo"
	initService "github.com/easysoft/zv/internal/server/service/init"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"net/http"
)

type Router struct {
	api *iris.Application

	InitService   *initService.InitService `inject:""`
	JwtService    *jwt.JwtService          `inject:""`
	TokenService  *jwt.TokenService        `inject:""`
	CasbinService *bizCasbin.CasbinService `inject:""`

	AccountCtrl *handler.AccountCtrl `inject:""`
	TaskCtrl    *handler.TaskCtrl    `inject:""`
	HostCtrl    *handler.HostCtrl    `inject:""`
	VmCtrl      *handler.VmCtrl      `inject:""`
	BuildCtrl   *handler.BuildCtrl   `inject:""`

	PermCtrl *handler.PermCtrl `inject:""`
	RoleCtrl *handler.RoleCtrl `inject:""`
	UserCtrl *handler.UserCtrl `inject:""`

	Environment *handler.EnvironmentCtrl `inject:""`
	ValidCtrl   *handler.ValidCtrl       `inject:""`
	WsCtrl      *handler.WsCtrl          `inject:""`

	TokenRepo *repo.TokenRepo `inject:""`
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
		// 二进制模式
		//if serverConf.Config.BinData {
		//	app.GetDetail("/", func(ctx iris.Context) { // 首页模块
		//		_ = ctx.View("index.html")
		//	})
		//}

		v1 := app.Party("/v1")
		{
			v1.PartyFunc("/client", func(client iris.Party) {
				client.PartyFunc("/task", func(party iris.Party) {
					party.Post("/list", r.TaskCtrl.List).Name = "列出任务"
					party.Post("/create", r.TaskCtrl.Create).Name = "创建任务"
					party.Put("/{id:uint}", r.TaskCtrl.Update).Name = "更新任务"
					party.Delete("/{id:uint}", r.TaskCtrl.Delete).Name = "删除任务"
				})

				client.PartyFunc("/host", func(party iris.Party) {
					party.Post("/register", r.HostCtrl.Register).Name = "注册主机"
				})
				client.PartyFunc("/vm", func(party iris.Party) {
					party.Post("/register", r.VmCtrl.Register).Name = "注册虚机"
				})
				client.PartyFunc("/build", func(party iris.Party) {
					party.Post("/uploadResult", r.BuildCtrl.UploadResult).Name = "文件上传"
				})

				client.Get("/testWs", r.TaskCtrl.TestWs).Name = "测试"
			})

			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Post("/login", r.AccountCtrl.UserLogin)

				admin.Use(r.JwtService.Serve, r.TokenService.Serve, r.CasbinService.Serve)

				admin.Post("/logout", r.AccountCtrl.UserLogout).Name = "退出"
				admin.Get("/expire", r.AccountCtrl.UserExpire).Name = "刷新Token"
				admin.Get("/profile", r.UserCtrl.GetProfile).Name = "个人信息"

				admin.PartyFunc("/users", func(party iris.Party) {
					party.Get("/", r.UserCtrl.GetAllUsers).Name = "用户列表"
					party.Get("/{id:uint}", r.UserCtrl.GetUser).Name = "用户详情"
					party.Post("/", r.UserCtrl.CreateUser).Name = "创建用户"
					party.Put("/{id:uint}", r.UserCtrl.UpdateUser).Name = "编辑用户"
					party.Delete("/{id:uint}", r.UserCtrl.DeleteUser).Name = "删除用户"
				})
				admin.PartyFunc("/roles", func(party iris.Party) {
					party.Get("/", r.RoleCtrl.GetAllRoles).Name = "角色列表"
					party.Get("/{id:uint}", r.RoleCtrl.GetRole).Name = "角色详情"
					party.Post("/", r.RoleCtrl.CreateRole).Name = "创建角色"
					party.Put("/{id:uint}", r.RoleCtrl.UpdateRole).Name = "编辑角色"
					party.Delete("/{id:uint}", r.RoleCtrl.DeleteRole).Name = "删除角色"
				})
				admin.PartyFunc("/permissions", func(party iris.Party) {
					party.Get("/", r.PermCtrl.GetAllPermissions).Name = "权限列表"
					party.Get("/{id:uint}", r.PermCtrl.GetPermission).Name = "权限详情"
					party.Post("/", r.PermCtrl.CreatePermission).Name = "创建权限"
					party.Put("/{id:uint}", r.PermCtrl.UpdatePermission).Name = "编辑权限"
					party.Delete("/{id:uint}", r.PermCtrl.DeletePermission).Name = "删除权限"
				})
			})

			v1.PartyFunc("/test", func(admin iris.Party) {
				admin.Use(r.JwtService.Serve, r.TokenService.Serve, r.CasbinService.Serve)

				admin.PartyFunc("/tasks", func(party iris.Party) {
					party.Get("/", r.TaskCtrl.List).Name = "任务列表"
					party.Get("/{id:uint}", r.TaskCtrl.Get).Name = "任务详情"
					party.Post("/", r.TaskCtrl.Create).Name = "创建任务"
					party.Put("/{id:uint}", r.TaskCtrl.Update).Name = "更新任务"
					party.Delete("/{id:uint}", r.TaskCtrl.Delete).Name = "删除任务"
				})

				admin.PartyFunc("/envs", func(party iris.Party) {
					party.Post("/", r.Environment.GetData).Name = "获取测试环境数据"
				})

				admin.PartyFunc("/valid", func(party iris.Party) {
					party.Post("/", r.ValidCtrl.Valid).Name = "表单验证"
				})
			})
		}

		// enable websocket
		websocketAPI := r.api.Party("/api/v1/ws")
		m := mvc.New(websocketAPI)
		m.Register(
			&prefixedLogger{prefix: "DEV"},
		)
		m.HandleWebsocket(handler.NewWsCtrl())

		websocketServer := websocket.New(
			gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), m)
		websocketAPI.Get("/", websocket.Handler(websocketServer))
	}
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}
