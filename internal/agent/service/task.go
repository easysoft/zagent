package agentService

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"time"
)

type TaskService struct {
	TM  time.Time
	Tasks []_domain.BuildTo
	IsRunning bool
}

func NewTaskService() *TaskService {
	service := &TaskService{}
	service.TM        = time.Now()
	service.Tasks     = make([]_domain.BuildTo, 0)

	return service
}

func (s *TaskService) AddTask(task _domain.BuildTo) {
	s.Tasks = append(s.Tasks, task)
}

func (s *TaskService) PeekTask() _domain.BuildTo {
	return s.Tasks[0]
}

func (s *TaskService) RemoveTask() (task _domain.BuildTo) {
	if len(s.Tasks) == 0 {
		return task
	}

	task = s.Tasks[0]
	s.Tasks = s.Tasks[1:]

	return task
}

func (s *TaskService) StartTask() {
	s.TM = time.Now()
	s.IsRunning = true
}
func (s *TaskService) EndTask() {
	s.IsRunning = false
}

func (s *TaskService) GetTaskSize() int {
	return len(s.Tasks)
}

func (s *TaskService) CheckTaskRunning() bool {
	if time.Now().Unix()-s.TM.Unix() > _const.AgentRunTime*60*1000 {
		s.IsRunning = false
	}
	return s.IsRunning
}
