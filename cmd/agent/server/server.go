package server

import (
	"github.com/easysoft/zagent/cmd/agent/router"
	"github.com/easysoft/zagent/cmd/agent/router/handler"
	agentCron "github.com/easysoft/zagent/internal/agent/cron"
	"github.com/easysoft/zagent/internal/agent/db"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func Init(router *router.Router) {
	db.InitDB()
	injectObj(router)
	router.App()
}

func injectObj(router *router.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		// db
		//&inject.Object{Value: db.GetInst().DB()},

		&inject.Object{Value: handler.NewArithCtrl()},
		&inject.Object{Value: handler.NewTaskCtrl()},

		&inject.Object{Value: agentService.NewCheckService()},
		&inject.Object{Value: agentService.NewCommonService()},
		&inject.Object{Value: agentService.NewRegisterService()},
		&inject.Object{Value: agentService.NewScmService()},
		&inject.Object{Value: agentService.NewTaskService()},

		&inject.Object{Value: agentService.NewBuildService()},
		&inject.Object{Value: agentService.NewAutomatedExecService()},
		&inject.Object{Value: agentService.NewAutomatedTestService()},
		&inject.Object{Value: agentService.NewInterfaceExecService()},
		&inject.Object{Value: agentService.NewInterfaceRequestService()},
		&inject.Object{Value: agentService.NewInterfaceTestService()},

		&inject.Object{Value: agentCron.NewCronService()},
		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
