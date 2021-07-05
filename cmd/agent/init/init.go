package init

import (
	"github.com/easysoft/zagent/cmd/agent/router"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentCron "github.com/easysoft/zagent/internal/agent/cron"
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

	if err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},
		&inject.Object{Value: agentCron.NewAgentCron()},

		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
