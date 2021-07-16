package hostKvmService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/libvirt/libvirt-go"
	"strings"
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
	s.VmService.UpdateVmMapAndDestroyTimeout(host.Vms)

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "client/host/register")
	resp, ok := _httpUtils.Post(url, host)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}

func (s *HostService) getVms() (vms []domain.Vm) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name, _ = dom.GetName()
		if strings.Index(vm.Name, "test-") != 0 {
			continue
		}

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
