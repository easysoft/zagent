package init

import (
	"github.com/easysoft/zagent/cmd/agent/router"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentCron "github.com/easysoft/zagent/internal/agent/cron"
	kvmService "github.com/easysoft/zagent/internal/agent/service/kvm"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	"github.com/easysoft/zagent/internal/pkg/db"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func Init(router *router.Router) {
	agentConf.Init()
	_db.InitDB("agent")
	injectObj(router)
	router.App()
}

func injectObj(router *router.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	var err error
	if agentConf.Inst.RunMode == agentConst.Host {
		err = g.Provide(
			// db
			&inject.Object{Value: _db.GetInst().DB()},
			// cron
			&inject.Object{Value: agentCron.NewAgentCron()},

			// service
			&inject.Object{Value: kvmService.NewLibvirtService()},
			&inject.Object{Value: kvmService.NewVmService()},

			// controller
			&inject.Object{Value: router},
		)
	} else {
		err = g.Provide(
			// db
			&inject.Object{Value: _db.GetInst().DB()},
			// cron
			&inject.Object{Value: agentCron.NewAgentCron()},

			// controller
			&inject.Object{Value: router},
		)
	}

	if err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err = g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
