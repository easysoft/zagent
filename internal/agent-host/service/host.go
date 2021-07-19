package hostAgentService

import (
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	testingService "github.com/easysoft/zagent/internal/agent/service/testing"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
)

type HostService struct {
	VmService *hostKvmService.VmService `inject:""`

	JobService  *agentService.JobService   `inject:""`
	TestService *testingService.RunService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Check() {
	// is running，register busy
	if s.JobService.IsRunning() {
		s.Register(true)
		return
	}

	// no task to run, submit free
	if s.JobService.GetTaskSize() == 0 {
		s.Register(false)
		return
	}

	// has task to run，register busy, then run
	job := s.JobService.PeekJob()
	s.Register(true)

	s.JobService.StartTask()
	s.TestService.Run(&job)
	s.JobService.RemoveTask()
	s.JobService.EndTask()
}

func (s *HostService) Register(isBusy bool) {
	host := domain.HostNode{
		Node: domain.Node{Ip: agentConf.Inst.NodeIp, Port: agentConf.Inst.NodePort},
	}
	if isBusy {
		host.Status = consts.HostBusy
	} else {
		host.Status = consts.HostReady
	}

	host.Vms = s.VmService.GetVms()
	s.VmService.UpdateVmMapAndDestroyTimeout(host.Vms)

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "client/host/register")
	resp, ok := _httpUtils.Post(url, host)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
