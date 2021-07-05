package agentService

import (
	interfaceService "github.com/easysoft/zagent/internal/agent/service/interface"
	testingService "github.com/easysoft/zagent/internal/agent/service/testing"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
)

type BuildService struct {
	TaskService          *TaskService                           `inject:""`
	InterfaceTestService *interfaceService.InterfaceTestService `inject:""`
	AutomatedTestService *testingService.TestService            `inject:""`
}

func NewBuildService() *BuildService {
	return &BuildService{}
}

func (s *BuildService) Exec(build domain.Build) {
	s.TaskService.StartTask()

	if build.BuildType == consts.AutoSelenium || build.BuildType == consts.AutoAppium {
		s.AutomatedTestService.Run(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
