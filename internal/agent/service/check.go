package agentService

import (
	agentUtils "github.com/easysoft/zagent/internal/agent/utils/common"
)

type CheckService struct {
	HostService   *HostService   `inject:""`
	VmService     *VmService     `inject:""`
	DeviceService *DeviceService `inject:""`

	TaskService  *TaskService  `inject:""`
	BuildService *BuildService `inject:""`
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
	if s.TaskService.IsRunning() {
		s.VmService.Register(true)
		return
	}

	// no task to run, submit free
	if s.TaskService.GetTaskSize() == 0 {
		s.VmService.Register(false)
		return
	}

	// has task to run，register busy, then run
	task := s.TaskService.PeekTask()
	s.VmService.Register(true)
	s.BuildService.Exec(task)
}

func (s *CheckService) CheckDevice() {
	// TODO:
}
