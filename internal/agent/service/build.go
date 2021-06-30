package agentService

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type BuildService struct {
	TaskService          *TaskService          `inject:""`
	InterfaceTestService *InterfaceTestService `inject:""`
	AutomatedTestService *AutomatedTestService `inject:""`
}

func NewBuildService() *BuildService {
	return &BuildService{}
}

func (s *BuildService) Exec(build commDomain.Build) {
	s.TaskService.StartTask()

	if build.BuildType == commConst.AutoSelenium || build.BuildType == commConst.AutoAppium {
		s.AutomatedTestService.Exec(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
