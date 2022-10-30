package hostInit

import (
	"fmt"
	hostCron "github.com/easysoft/zagent/cmd/host/cron"
	hostRouter "github.com/easysoft/zagent/cmd/host/router"
	hostKvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	virtualService "github.com/easysoft/zagent/internal/host/service/virtual"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_db "github.com/easysoft/zagent/pkg/db"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Init() {
	agentConf.Init(consts.AppNameAgentHost)
	_db.InitDB("agent")

	irisServer := NewServer(nil)
	irisServer.App.Logger().SetLevel("info")

	router := hostRouter.NewRouter(irisServer.App)
	injectObj(router)

	router.InitService.InitModels()

	router.App()

	iris.RegisterOnInterrupt(func() {
		defer _db.GetInst().Close()
	})

	err := irisServer.Serve()
	if err != nil {
		panic(err)
	}
}

func injectObj(router *hostRouter.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},

		// setup
		&inject.Object{Value: virtualService.NewNovncService()},

		// cron
		&inject.Object{Value: hostCron.NewAgentCron()},

		// service
		&inject.Object{Value: hostKvmService.NewLibvirtService()},
		&inject.Object{Value: hostKvmService.NewVmService()},

		// controller
		&inject.Object{Value: router},
	)

	if err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err = g.Populate()
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
	if _commonUtils.IsPortInUse(agentConf.Inst.NodePort) {
		panic(fmt.Sprintf("端口 %d 已被使用", agentConf.Inst.NodePort))
	}

	err := s.App.Run(
		iris.Addr(fmt.Sprintf("%s:%d", "", agentConf.Inst.NodePort)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(time.RFC3339),
	)

	if err != nil {
		return err
	}

	return nil
}
