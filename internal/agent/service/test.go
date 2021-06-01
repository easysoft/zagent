package agentService

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type TestService struct {
	TaskService   *TaskService      `inject:""`
	InterfaceService   *InterfaceService      `inject:""`
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) Exec(build _domain.BuildTo) {
	s.TaskService.StartTask()

	if build.BuildType == _const.InterfaceTest {
		s.InterfaceService.ExecTest(&build)
	}

	s.TaskService.RemoveTask()
	s.TaskService.EndTask()
}
