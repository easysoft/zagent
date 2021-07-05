package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/libvirt/libvirt-go"
)

type HostService struct {
	VmService      *VmService      `inject:""`
	LibvirtService *LibvirtService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Register() {
	host := domain.HostNode{
		Node:       domain.Node{Ip: agentConf.Inst.NodeIp, Port: agentConf.Inst.NodePort},
		HostStatus: consts.HostActive,
	}
	host.Vms = s.getVms()
	s.VmService.UpdateVmsStatus(host.Vms)

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

func (s *HostService) getVms() (vms []domain.Vm) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name, _ = dom.GetName()

		vm.Status = consts.VmUnknown
		domainState, _, _ := dom.GetState()
		if domainState == libvirt.DOMAIN_RUNNING {
			vm.Status = consts.VmRunning
		} else if domainState == libvirt.DOMAIN_SHUTOFF {
			vm.Status = consts.VmShutOff
		}

		vms = append(vms, vm)
	}

	return vms
}
