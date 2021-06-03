package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type BuildService struct {
	TaskService          *TaskService          `inject:""`
	InterfaceTestService *InterfaceTestService `inject:""`
	AutomatedTestService *AutomatedTestService `inject:""`
}

func NewTestService() *BuildService {
	return &BuildService{}
}

func (s *BuildService) Exec(build commDomain.Build) {
	s.TaskService.StartTask()

	if build.BuildType == _const.InterfaceScenario {
		s.InterfaceTestService.ExecScenario(&build)

	} else if build.BuildType == _const.InterfaceSet {
		s.InterfaceTestService.ExecSet(&build)

	} else if build.BuildType == _const.AutomatedTest {
		s.AutomatedTestService.Exec(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
