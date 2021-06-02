package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/libs/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/libs/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
)

type RegisterService struct {
}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (s *RegisterService) Register(isBusy bool) {
	node := commDomain.Node{Name: agentConf.Inst.NodeName, WorkDir: agentConf.Inst.WorkDir,
		Ip: agentConf.Inst.NodeIp, Port: agentConf.Inst.NodePort}

	if isBusy {
		node.Status = _const.ServiceBusy
	} else {
		node.Status = _const.ServiceActive
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "vms/register")
	resp, ok := _httpUtils.Post(url, node)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
