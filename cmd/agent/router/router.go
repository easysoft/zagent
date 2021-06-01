package router

import (
	"fmt"
	"github.com/easysoft/zagent/cmd/agent/router/handler"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"github.com/smallnest/rpcx/server"
	"strconv"
)

type Router struct {
	ArithCtrl   *handler.ArithCtrl      `inject:""`
	InterfaceCtrl   *handler.InterfaceCtrl      `inject:""`
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) App() {
	addr := agentConf.Inst.NodeIp + ":" + strconv.Itoa(agentConf.Inst.NodePort)

	srv := server.NewServer()
	srv.RegisterName("Arith", r.ArithCtrl, "")
	srv.RegisterName("Interface", r.InterfaceCtrl, "")

	_logUtils.Info(fmt.Sprintf("start server on %s ...", addr))
	err := srv.Serve("tcp", addr)
	if err != nil {
		_logUtils.Infof("fail to start server on %s, err is %s", err.Error())
	}
}
