package agentService

import (
	"github.com/easysoft/zagent/internal/agent/model"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type TestService struct {
	TaskService      *TaskService      `inject:""`
	InterfaceService *InterfaceService `inject:""`
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) Exec(build domain.BuildTo) {
	s.TaskService.StartTask()

	if build.BuildType == _const.InterfaceTest {
		s.InterfaceService.ExecTest(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
