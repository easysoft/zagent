package hostAgentService

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	downloadUtils "github.com/easysoft/zagent/pkg/lib/download"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"sync"
)

var (
	channelMap sync.Map
)

type DownloadService struct {
	TaskService *TaskService `inject:""`

	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

func (s *DownloadService) AddTasks(req v1.DownloadReq) (err error) {
	for _, item := range req.Urls {
		po := agentModel.Task{
			Url:        item,
			ZentaoTask: req.ZentaoTask,
			TaskType:   consts.DownloadImage,
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *DownloadService) StartTask(po agentModel.Task) {
	ch := make(chan int, 1)
	channelMap.Store(int(po.ID), ch)

	go func() {
		filePath, finalStatus := downloadUtils.Start(po, ch)

		s.TaskRepo.UpdateStatus(po.ID, filePath, "", finalStatus)

		po = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		downloadUtils.TaskMap.Delete(po.ID)

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

	chVal, ok := channelMap.Load(task.ID)

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
	s.TaskRepo.Delete(uint(req.ZentaoTask))

	return
}
