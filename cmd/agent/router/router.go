package router

import (
	"github.com/easysoft/zagent/cmd/agent/router/handler"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/smallnest/rpcx/server"
	"strconv"
)

type Router struct {
	ArithCtrl *handler.ArithCtrl `inject:""`
	TaskCtrl  *handler.TaskCtrl  `inject:""`
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) App() {
	addr := agentConf.Inst.NodeIp + ":" + strconv.Itoa(agentConf.Inst.NodePort)
	srv := server.NewServer()

	if agentConf.Inst.RunMode == agentConst.Host {
		srv.RegisterName("arith", r.ArithCtrl, "")
		srv.RegisterName("task", r.TaskCtrl, "")

		_logUtils.Info(_i118Utils.Sprintf("start_server", addr))
		err := srv.Serve("tcp", addr)
		if err != nil {
			_logUtils.Infof(_i118Utils.Sprintf("fail_to_start_server", addr, err.Error()))
		}

	} else if agentConf.Inst.RunMode == agentConst.Vm {
		srv.RegisterName("selenium", r.ArithCtrl, "")

	} else if agentConf.Inst.RunMode == agentConst.Android || agentConf.Inst.RunMode == agentConst.Ios {
		srv.RegisterName("appium", r.ArithCtrl, "")

	}
}
