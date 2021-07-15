package hostRouter

import (
	"github.com/easysoft/zagent/cmd/agent/router/handler"
	hostHandler "github.com/easysoft/zagent/cmd/host/router/handler"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/smallnest/rpcx/server"
	"strconv"
)

type Router struct {
	ArithCtrl *handler.ArithCtrl  `inject:""`
	VmCtrl    *hostHandler.VmCtrl `inject:""`
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) App() {
	addr := agentConf.Inst.NodeIp + ":" + strconv.Itoa(agentConf.Inst.NodePort)
	srv := server.NewServer()

	srv.RegisterName("arith", r.ArithCtrl, "")
	srv.RegisterName("vm", r.VmCtrl, "")

	_logUtils.Info(_i118Utils.Sprintf("start_server", addr))
	err := srv.Serve("tcp", addr)
	if err != nil {
		_logUtils.Infof(_i118Utils.Sprintf("fail_to_start_server", addr, err.Error()))
	}
}