package agentService

import (
	deviceService "github.com/easysoft/zagent/internal/agent/service/device"
	kvmService "github.com/easysoft/zagent/internal/agent/service/kvm"
	testingService "github.com/easysoft/zagent/internal/agent/service/testing"
	agentUtils "github.com/easysoft/zagent/internal/agent/utils/common"
)

type CheckService struct {
	HostService   *kvmService.HostService      `inject:""`
	VmService     *kvmService.VmService        `inject:""`
	DeviceService *deviceService.DeviceService `inject:""`

	JobService  *JobService                `inject:""`
	TestService *testingService.RunService `inject:""`
}

func NewCheckService() *CheckService {
	return &CheckService{}
}

func (s *CheckService) Check() {
	if agentUtils.IsHostAgent() { // host
		s.HostService.Register()

	} else if agentUtils.IsVmAgent() { // vm
		s.CheckVm()

	} else if agentUtils.IsDeviceAgent() { // device
		s.CheckDevice()

	}
}

func (s *CheckService) CheckVm() {
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

func (s *CheckService) CheckDevice() {
	// TODO:
}
