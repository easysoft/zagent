package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"strings"
)

type TaskService struct {
	TaskRepo  *repo.TaskRepo  `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`

	QueueService   *QueueService   `inject:""`
	HistoryService *HistoryService `inject:""`
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
func (s *TaskService) GetDetail(id uint) (po model.Task) {
	po = s.Get(id)

	po.Queues = s.QueueRepo.QueryByTask(po.ID)

	return
}

func (s *TaskService) Save(po *model.Task, userId uint) (err error) {
	if strings.Index(po.ScriptUrl, ".zip") < 0 {
		po.ScmAddress = po.ScriptUrl
		po.ScriptUrl = ""
	}

	po.UserId = userId
	err = s.TaskRepo.Save(po)

	s.QueueService.GenerateFromTask(po)
	s.HistoryService.Create(consts.Task, po.ID, consts.ProgressCreated, "")

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

func (s *TaskService) SetProgress(id uint, progress consts.BuildProgress) {
	s.TaskRepo.SetProgress(id, progress)

	s.HistoryService.Create(consts.Task, id, progress, "")
}

func (s *TaskService) SetTaskStatus(taskId uint) {
	queues := s.QueueRepo.QueryByTask(taskId)

	progress := consts.ProgressCompleted
	status := consts.StatusPass
	isAllQueuesCompleted := true

	for _, queue := range queues {
		if queue.Progress != consts.ProgressCompleted &&
			queue.Progress != consts.ProgressTimeout { // 有queue在进行中
			isAllQueuesCompleted = false
			break
		}

		if queue.Progress == consts.ProgressTimeout { // 有一个超时，就超时
			progress = consts.ProgressTimeout
		}

		if queue.Status == consts.StatusFail { // 有一个失败，就失败
			status = consts.StatusFail
		}
	}

	if isAllQueuesCompleted {
		s.TaskRepo.SetResult(taskId, progress, status)
		s.HistoryService.Create(consts.Task, taskId, progress, status)
	}
}
