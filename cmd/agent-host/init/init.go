package hostInit

import (
	hostCron "github.com/easysoft/zagent/cmd/host/cron"
	hostRouter "github.com/easysoft/zagent/cmd/host/router"
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/pkg/db"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func Init(router *hostRouter.Router) {
	agentConf.Init()
	_db.InitDB("agent")
	injectObj(router)
	router.App()
}

func injectObj(router *hostRouter.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},

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
