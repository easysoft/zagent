package agentService

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	"sync"
	"time"
)

var lock sync.Mutex

type JobService struct {
	TimeStamp time.Time
	running   bool
	jobs      []domain.Build
}

func NewTaskService() *JobService {
	service := &JobService{}

	service.TimeStamp = time.Now()
	service.jobs = make([]domain.Build, 0)

	return service
}

func (s *JobService) AddTask(build domain.Build) {
	lock.Lock()

	s.jobs = append(s.jobs, build)

	lock.Unlock()
}

func (s *JobService) PeekJob() domain.Build {
	lock.Lock()
	defer lock.Unlock()

	return s.jobs[0]
}

func (s *JobService) RemoveTask() {
	lock.Lock()

	if len(s.jobs) == 0 {
		return
	}
	s.jobs = s.jobs[1:]

	lock.Unlock()
}

func (s *JobService) StartTask() {
	lock.Lock()

	s.TimeStamp = time.Now()
	s.running = true

	lock.Unlock()
}
func (s *JobService) EndTask() {
	lock.Lock()

	s.running = false

	lock.Unlock()
}

func (s *JobService) GetTaskSize() int {
	lock.Lock()
	defer lock.Unlock()

	return len(s.jobs)
}

func (s *JobService) IsRunning() bool {
	lock.Lock()
	defer lock.Unlock()

	if time.Now().Unix()-s.TimeStamp.Unix() > consts.WaitAgentRunTaskTimeout {
		s.running = false
	}
	return s.running
}
