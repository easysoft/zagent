package service

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type TaskService struct {
	TaskRepo  *repo.TaskRepo  `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`

	QueueService *QueueService `inject:""`
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) List(keywords, status string, pageNo int, pageSize int) (pos []model.Task, total int64) {
	pos, total = s.TaskRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *TaskService) Get(id uint) (po model.Task) {
	po = s.TaskRepo.Get(id)
	return
}

func (s *TaskService) Save(po *model.Task, userId uint) (err error) {
	po.UserId = userId
	err = s.TaskRepo.Save(po)

	s.QueueService.GenerateFromTask(po)

	return
}

func (s *TaskService) Update(po *model.Task) (err error) {
	err = s.TaskRepo.Update(po)

	s.QueueService.GenerateFromTask(po)

	return
}

func (s *TaskService) Disable(id uint) (err error) {
	err = s.TaskRepo.Disable(id)

	return
}

func (s *TaskService) Delete(id uint) (err error) {
	err = s.TaskRepo.Delete(id)

	return
}

func (s *TaskService) SetProgress(id uint, progress commConst.BuildProgress) {
	s.TaskRepo.SetProgress(id, progress)
}

func (s *TaskService) CheckCompleted(taskId uint) {
	queues := s.QueueRepo.QueryByTask(taskId)

	progress := commConst.ProgressCompleted
	status := commConst.StatusPass
	isAllQueuesCompleted := true

	for _, queue := range queues {
		if queue.Progress != commConst.ProgressCompleted && queue.Progress != commConst.ProgressTimeout { // 有queue在进行中
			isAllQueuesCompleted = false
			break
		}

		if queue.Progress == commConst.ProgressTimeout { // 有一个超时，就超时
			progress = commConst.ProgressTimeout
		}

		if queue.Status == commConst.StatusFail { // 有一个失败，就失败
			status = commConst.StatusFail
		}
	}

	if isAllQueuesCompleted {
		s.TaskRepo.SetResult(taskId, progress, status)
	}
}
