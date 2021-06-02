package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type BuildService struct {
	TaskService          *TaskService          `inject:""`
	InterfaceTestService *InterfaceTestService `inject:""`
}

func NewTestService() *BuildService {
	return &BuildService{}
}

func (s *BuildService) Exec(build commDomain.Build) {
	s.TaskService.StartTask()

	if build.BuildType == _const.InterfaceScenario || build.BuildType == _const.InterfaceSet {
		s.InterfaceTestService.ExecTest(&build)

	} else if build.BuildType == _const.AutomatedTest {

	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
