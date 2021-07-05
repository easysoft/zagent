package agentService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
)

type BuildService struct {
	TaskService          *TaskService          `inject:""`
	InterfaceTestService *InterfaceTestService `inject:""`
	AutomatedTestService *AutomatedTestService `inject:""`
}

func NewBuildService() *BuildService {
	return &BuildService{}
}

func (s *BuildService) Exec(build domain.Build) {
	s.TaskService.StartTask()

	if build.BuildType == consts.AutoSelenium || build.BuildType == consts.AutoAppium {
		s.AutomatedTestService.Exec(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
