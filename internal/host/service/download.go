package hostAgentService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentModel "github.com/easysoft/zv/internal/host/model"
	hostRepo "github.com/easysoft/zv/internal/host/repo"
	consts "github.com/easysoft/zv/internal/pkg/const"
	downloadUtils "github.com/easysoft/zv/pkg/lib/download"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"sync"
	"time"
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

func (s *DownloadService) CheckTask() (err error) {
	taskMap, _ := s.ListTask()

	toStartNewTask := false
	if len(taskMap[consts.InProgress]) > 0 {
		runningTask := taskMap[consts.InProgress][0]

		if s.IsError(runningTask) || s.IsTimeout(runningTask) {
			if s.NeedRetry(runningTask) {
				s.RestartTask(runningTask)
			} else {
				s.TaskRepo.SetFailed(runningTask)
				toStartNewTask = true
			}
		}
	} else {
		toStartNewTask = true
	}

	if toStartNewTask && len(taskMap[consts.Created]) > 0 {
		s.StartTask(taskMap[consts.Created][0])
	}

	return
}

func (s *DownloadService) ListTask() (ret map[consts.DownloadStatus][]agentModel.Download, err error) {
	ret = map[consts.DownloadStatus][]agentModel.Download{}

	pos, _ := s.TaskRepo.Query()

	for _, po := range pos {
		status := po.Status
		if status == consts.Timeout || status == consts.Error {
			status = consts.InProgress
		}

		ret[status] = append(ret[status], po)
	}

	return
}

func (s *DownloadService) AddTasks(req v1.DownloadReq) (err error) {
	for _, item := range req.Urls {
		po := agentModel.Download{
			Url: item,
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *DownloadService) StartTask(po agentModel.Download) {
	ch := make(chan int, 1)
	syncMap.Store(int(po.ID), ch)

	go func() {
		filePath, finalStatus := downloadUtils.Start(po, ch)
		s.TaskRepo.UpdateStatus(po.ID, finalStatus, filePath)

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

func (s *DownloadService) RestartTask(po agentModel.Download) (ret bool) {
	s.CancelTask(po.Url)
	s.StartTask(po)

	return
}

func (s *DownloadService) RemoveTask(req v1.DownloadReq) {
	s.TaskRepo.Delete(uint(req.TaskId))

	return
}

func (s *DownloadService) IsError(po agentModel.Download) bool {
	return po.Status == consts.Error
}

func (s *DownloadService) IsTimeout(po agentModel.Download) bool {
	return time.Now().Unix()-po.StartTime.Unix() > consts.DownloadTimeout
}

func (s *DownloadService) NeedRetry(po agentModel.Download) bool {
	return po.Retry <= consts.DownloadRetry
}
