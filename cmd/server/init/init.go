package init

import (
	"fmt"
	"github.com/easysoft/zagent/cmd/server/router"
	"github.com/easysoft/zagent/cmd/server/router/handler"
	"github.com/easysoft/zagent/internal/pkg/db"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	bizCasbin "github.com/easysoft/zagent/internal/server/biz/casbin"
	"github.com/easysoft/zagent/internal/server/biz/jwt"
	"github.com/easysoft/zagent/internal/server/biz/redis"
	"github.com/easysoft/zagent/internal/server/conf"
	serverCron "github.com/easysoft/zagent/internal/server/cron"
	"github.com/easysoft/zagent/internal/server/model"
	initService "github.com/easysoft/zagent/internal/server/service/init"
	serverRes "github.com/easysoft/zagent/res/server"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12/context"

	_ "github.com/easysoft/zagent/res/server/docs"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
)

func Init(version string, printVersion, printRouter *bool) {
	serverConf.Init()
	_db.InitDB("server")

	irisServer := NewServer(nil)
	irisServer.App.Logger().SetLevel(serverConf.Inst.LogLevel)
	mapStaticRes(irisServer)

	router := router.NewRouter(irisServer.App)
	injectObj(router)

	router.InitService.InitDataIfNeeded()

	router.App()

	config := swagger.Config{
		URL:          "http://localhost:8085/swagger/doc.json",
		DeepLinking:  true,
		DocExpansion: "list",
		DomID:        "#swagger-ui",
		Prefix:       "/swagger",
	}
	swaggerUI := swagger.Handler(swaggerFiles.Handler, config)

	irisServer.App.Get("/swagger", swaggerUI)
	irisServer.App.Get("/swagger/{any:path}", swaggerUI)

	if serverConf.Inst.Redis.Enable {
		redisUtils.InitRedisCluster(serverConf.GetRedisUris(), serverConf.Inst.Redis.Pwd)
	}

	iris.RegisterOnInterrupt(func() {
		defer _db.GetInst().Close()
	})

	// deal with the command
	if *printVersion {
		fmt.Println(fmt.Sprintf("版本号：%s", version))
	}

	if *printRouter {
		fmt.Println("系统权限：")
		fmt.Println()
		routes := irisServer.GetRoutes()
		for _, route := range routes {
			fmt.Println("+++++++++++++++")
			fmt.Println(fmt.Sprintf("名称 ：%s ", route.DisplayName))
			fmt.Println(fmt.Sprintf("路由地址 ：%s ", route.Name))
			fmt.Println(fmt.Sprintf("请求方式 ：%s", route.Act))
			fmt.Println()
		}
	}

	// start the service
	err := irisServer.Serve()
	if err != nil {
		panic(err)
	}
}

func injectObj(router *router.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},

		// repo

		// middleware
		&inject.Object{Value: bizCasbin.NewEnforcer()},
		&inject.Object{Value: jwt.NewJwtService()},

		// service
		&inject.Object{Value: initService.NewSeeder()},
		&inject.Object{Value: serverCron.NewServerCron()},

		// controller
		&inject.Object{Value: handler.NewPermCtrl()},
		&inject.Object{Value: handler.NewRoleCtrl()},
		&inject.Object{Value: handler.NewUserCtrl()},

		&inject.Object{Value: handler.NewTaskCtrl()},

		// router
		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}

type Server struct {
	App       *iris.Application
	AssetFile http.FileSystem
}

func NewServer(assetFile http.FileSystem) *Server {
	app := iris.Default()
	return &Server{
		App:       app,
		AssetFile: assetFile,
	}
}

func (s *Server) Serve() error {
	if _commonUtils.IsPortInUse(serverConf.Inst.Port) {
		panic(fmt.Sprintf("端口 %d 已被使用", serverConf.Inst.Port))
	}

	if serverConf.Inst.Https {
		host := fmt.Sprintf("%s:%d", serverConf.Inst.Host, serverConf.Inst.Port)
		err := s.App.Run(iris.TLS(host, serverConf.Inst.CertPath, serverConf.Inst.CertKey))

		if err != nil {
			return err
		}

	} else {
		err := s.App.Run(
			iris.Addr(fmt.Sprintf("%s:%d", serverConf.Inst.Host, serverConf.Inst.Port)),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			iris.WithTimeFormat(time.RFC3339),
		)

		if err != nil {
			return err
		}
	}

	return nil
}

type PathName struct {
	Name   string
	Path   string
	Method string
}

// 获取路由信息
func (s *Server) GetRoutes() []*model.Permission {
	var rrs []*model.Permission
	names := getPathNames(s.App.GetRoutesReadOnly())
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("路由权限集合：%v", names))
		fmt.Println(fmt.Sprintf("Iris App ：%v", s.App))
	}
	for _, pathName := range names {
		if !isPermRoute(pathName.Name) {
			rr := &model.Permission{Name: pathName.Path, DisplayName: pathName.Name, Description: pathName.Name, Act: pathName.Method}
			rrs = append(rrs, rr)
		}
	}
	return rrs
}

func getPathNames(routeReadOnly []context.RouteReadOnly) []*PathName {
	var pns []*PathName
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("routeReadOnly：%v", routeReadOnly))
	}
	for _, s := range routeReadOnly {
		pn := &PathName{
			Name:   s.Name(),
			Path:   s.Path(),
			Method: s.Method(),
		}
		pns = append(pns, pn)
	}

	return pns
}

func mapStaticRes(irisServer *Server) {
	// ui
	if _commonUtils.IsRelease() {
		irisServer.App.HandleDir("/",
			&assetfs.AssetFS{Asset: serverRes.Asset, AssetDir: serverRes.AssetDir, AssetInfo: serverRes.AssetInfo,
				Prefix: "ui/dist"}, iris.DirOptions{
				IndexName: "index.html",
				Compress:  false,
			})
	}

	// testing res
	irisServer.App.HandleDir("/down", "down")
}

// 过滤非必要权限
func isPermRoute(name string) bool {
	exceptRouteName := []string{"OPTIONS", "GET", "POST", "HEAD", "PUT", "PATCH", "payload"}
	for _, er := range exceptRouteName {
		if strings.Contains(name, er) {
			return true
		}
	}
	return false
}
