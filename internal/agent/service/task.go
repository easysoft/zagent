package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type TaskService struct {
	TM        time.Time
	Tasks     []commDomain.BuildTo
	IsRunning bool
}

func NewTaskService() *TaskService {
	service := &TaskService{}
	service.TM = time.Now()
	service.Tasks = make([]commDomain.BuildTo, 0)

	return service
}

func (s *TaskService) AddTask(task commDomain.BuildTo) {
	s.Tasks = append(s.Tasks, task)
}

func (s *TaskService) PeekTask() commDomain.BuildTo {
	return s.Tasks[0]
}

func (s *TaskService) RemoveTask() (task commDomain.BuildTo) {
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
