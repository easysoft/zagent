package server

import (
	"github.com/easysoft/zagent/cmd/agent/router"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/agent/db"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func Init(router *router.Router) {
	agentConf.Init()
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

		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
