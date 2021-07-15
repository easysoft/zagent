package vmInit

import (
	vmCron "github.com/easysoft/zagent/cmd/vm/cron"
	vmRouter "github.com/easysoft/zagent/cmd/vm/router"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/pkg/db"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func Init(router *vmRouter.Router) {
	agentConf.Init()
	_db.InitDB("agent")
	injectObj(router)
	router.App()
}

func injectObj(router *vmRouter.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},
		// cron
		&inject.Object{Value: vmCron.NewAgentCron()},

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
