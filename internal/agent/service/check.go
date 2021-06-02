package agentService

type CheckService struct {
	RegisterService *RegisterService `inject:""`
	TaskService     *TaskService     `inject:""`
	TestService     *TestService     `inject:""`
}

func NewCheckService() *CheckService {
	return &CheckService{}
}

func (s *CheckService) Check() {
	// is running，register busy
	if s.TaskService.IsRunning() {
		s.RegisterService.Register(true)
		return
	}

	// no task to run, submit free
	if s.TaskService.GetTaskSize() == 0 {
		s.RegisterService.Register(false)
		return
	}

	// has task to run，register busy, then run
	build := s.TaskService.PeekTask()
	s.RegisterService.Register(true)
	s.TestService.Exec(build)
}
