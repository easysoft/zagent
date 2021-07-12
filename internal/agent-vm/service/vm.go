package vmAgentService

import (
	"github.com/easysoft/zagent/internal/agent/conf"
	testingService "github.com/easysoft/zagent/internal/agent/service/testing"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/pkg/lib/i118"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	"time"
)

type VmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	VmService   *VmService                 `inject:""`
	JobService  *JobService                `inject:""`
	TestService *testingService.RunService `inject:""`
}

func NewVmService() *VmService {
	s := VmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *VmService) Check() {
	// is running，register busy
	if s.JobService.IsRunning() {
		s.VmService.Register(true)
		return
	}

	// no task to run, submit free
	if s.JobService.GetTaskSize() == 0 {
		s.VmService.Register(false)
		return
	}

	// has task to run，register busy, then run
	job := s.JobService.PeekJob()
	s.VmService.Register(true)

	s.JobService.StartTask()
	s.TestService.Run(&job)
	s.JobService.RemoveTask()
	s.JobService.EndTask()
}

func (s *VmService) Register(isBusy bool) {
	vm := domain.Vm{
		MacAddress: agentConf.Inst.MacAddress,
		Name:       agentConf.Inst.NodeName, WorkDir: agentConf.Inst.WorkDir,
		PublicIp: agentConf.Inst.NodeIp, PublicPort: agentConf.Inst.NodePort,
	}

	if isBusy {
		vm.Status = consts.VmBusy
	} else {
		vm.Status = consts.VmReady
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "client/vm/register")
	resp, ok := _httpUtils.Post(url, vm)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
