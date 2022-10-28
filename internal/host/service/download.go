package hostAgentService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentModel "github.com/easysoft/zv/internal/host/model"
	hostRepo "github.com/easysoft/zv/internal/host/repo"
	consts "github.com/easysoft/zv/internal/pkg/const"
	downloadUtils "github.com/easysoft/zv/pkg/lib/download"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"sync"
)

const (
	key = "urls"
)

const (
	keyNotStart   = "not_start"
	keyInProgress = "in_progress"
	keyCompleted  = "completed"
)

var (
	syncMap sync.Map
)

type DownloadService struct {
	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

func (s *DownloadService) AddTasks(req v1.DownloadReq) (err error) {
	for _, item := range req.Urls {
		po := agentModel.Task{
			Url:      item,
			TaskType: consts.DownloadImage,
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *DownloadService) StartTask(po agentModel.Task) {
	ch := make(chan int, 1)
	syncMap.Store(int(po.ID), ch)

	go func() {
		filePath, finalStatus := downloadUtils.Start(po, ch)
		s.TaskRepo.UpdateStatus(po.ID, filePath, "", finalStatus)

		if ch != nil {
			close(ch)
		}
	}()
}

func (s *DownloadService) CancelTask(url string) {
	task, err := s.TaskRepo.GetByUrl(url)
	if err != nil {
		_logUtils.Infof("can not find task %s to cancel.")
		return
	}

	chVal, ok := syncMap.Load(task.ID)

	if !ok || chVal == nil {
		return
	}

	ch := chVal.(chan int)
	if ch != nil {
		ch <- 1
		ch = nil
	}

	return
}

func (s *DownloadService) RestartTask(po agentModel.Task) (ret bool) {
	s.CancelTask(po.Url)
	s.StartTask(po)

	return
}

func (s *DownloadService) RemoveTask(req v1.DownloadReq) {
	s.TaskRepo.Delete(uint(req.TaskId))

	return
}
