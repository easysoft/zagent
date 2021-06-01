package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
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
	node := _domain.Node{HostName: agentConf.Inst.NodeName, WorkDir: agentConf.Inst.WorkDir,
		PublicIp: agentConf.Inst.NodeIp, PublicPort: agentConf.Inst.NodePort}

	if isBusy {
		node.Status = _const.NodeBusy
	} else {
		node.Status = _const.NodeActive
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "vms/register")
	resp, ok := _httpUtils.Post(url, node)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
