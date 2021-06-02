package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type TestService struct {
	TaskService      *TaskService      `inject:""`
	InterfaceService *InterfaceService `inject:""`
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) Exec(build commDomain.BuildTo) {
	s.TaskService.StartTask()

	if build.BuildType == _const.InterfaceTest {
		s.InterfaceService.ExecTest(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
