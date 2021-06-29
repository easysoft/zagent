package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"sync"
	"time"
)

var lock sync.Mutex

type TaskService struct {
	TimeStamp time.Time
	running   bool
	tasks     []commDomain.IntfTest
}

func NewTaskService() *TaskService {
	service := &TaskService{}

	service.TimeStamp = time.Now()
	service.tasks = make([]commDomain.IntfTest, 0)

	return service
}

func (s *TaskService) AddTask(task commDomain.IntfTest) {
	lock.Lock()

	s.tasks = append(s.tasks, task)

	lock.Unlock()
}

func (s *TaskService) PeekTask() commDomain.IntfTest {
	lock.Lock()
	defer lock.Unlock()

	return s.tasks[0]
}

func (s *TaskService) RemoveTask() {
	lock.Lock()

	if len(s.tasks) == 0 {
		return
	}
	s.tasks = s.tasks[1:]

	lock.Unlock()
}

func (s *TaskService) StartTask() {
	lock.Lock()

	s.TimeStamp = time.Now()
	s.running = true

	lock.Unlock()
}
func (s *TaskService) EndTask() {
	lock.Lock()

	s.running = false

	lock.Unlock()
}

func (s *TaskService) GetTaskSize() int {
	lock.Lock()
	defer lock.Unlock()

	return len(s.tasks)
}

func (s *TaskService) IsRunning() bool {
	lock.Lock()
	defer lock.Unlock()

	if time.Now().Unix()-s.TimeStamp.Unix() > _const.AgentRunTime*60*1000 {
		s.running = false
	}
	return s.running
}
