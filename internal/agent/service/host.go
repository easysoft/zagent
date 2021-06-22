package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/libs/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/libs/shell"
	"strings"
)

type HostService struct {
	VmService *VmService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Register() {
	host := commDomain.HostNode{
		Node:       commDomain.Node{Ip: agentConf.Inst.NodeIp, Port: agentConf.Inst.NodePort},
		HostStatus: commConst.HostActive,
	}
	host.Vms = s.getVms()
	s.VmService.UpdateVms(host.Vms)

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "host/register")
	resp, ok := _httpUtils.Post(url, host)

	msg := ""
	str := "%s to register host, response is %#v"
	if ok {
		msg = "success"
		_logUtils.Infof(str, msg, resp)
	} else {
		msg = "fail"
		_logUtils.Errorf(str, msg, resp)
	}
}

func (s *HostService) getVms() (vms []commDomain.Vm) {
	cmd := "virsh list --all"
	out, _ := _shellUtils.ExeShell(cmd)

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Index(line, "Id") == 0 || strings.Index(line, "---") == 0 {
			continue
		}

		cols := strings.Split(line, " ")
		name := strings.TrimSpace(cols[1])
		status := strings.TrimSpace(cols[2])

		if len(name) < 32 { // not created by farm
			continue
		}

		vm := commDomain.Vm{}
		vm.Name = name

		vm.Status = commConst.VmUnknown
		if status == "running" {
			vm.Status = commConst.VmRunning
		} else if status == "shut off" {
			vm.Status = commConst.VmDestroy
		}

		vms = append(vms, vm)
	}

	return vms
}
