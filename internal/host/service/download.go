package hostAgentService

import (
	"errors"
	"sync"

	downloadUtils "github.com/easysoft/zagent/internal/pkg/job"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_channelUtils "github.com/easysoft/zagent/pkg/lib/channel"
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

func (s *DownloadService) AddTasks(req []v1.DownloadReq) (err error) {
	for _, item := range req {
		po := agentModel.Task{
			Url:        item.Url,
			Md5:        item.Md5,
			ZentaoTask: item.ZentaoTask,
			TaskType:   consts.DownloadImage,
			Retry:      1,
			Status:     consts.Created,
		}

		existInfo, _ := s.TaskRepo.GetByMd5(item.Md5)

		if existInfo.ID != 0 {
			if existInfo.Status == consts.InProgress {
				err = errors.New("the same md5 task exists and downloading")
				return
			} else {
				s.TaskRepo.SetFailed(existInfo)
			}
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *DownloadService) StartTask(po agentModel.Task) {
	ch := make(chan int, 1)
	channelMap.Store(po.ID, ch)

	go func() {
		filePath := downloadUtils.GetPath(po)

		s.TaskRepo.UpdateStatus(po.ID, filePath, 0.01, "", consts.InProgress, true, false)

		finalStatus, existFile := downloadUtils.Start(&po, filePath, ch)
		if existFile != "" {
			filePath = existFile
		}

		s.TaskRepo.UpdateSpeed(po.ID, po.Speed)
		s.TaskRepo.UpdateStatus(po.ID, filePath, 1, "", finalStatus, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		downloadUtils.TaskStatus.Delete(po.ID)

		if ch != nil {
			channelMap.Delete(po.ID)
			close(ch)
		}
	}()
}

func (s *DownloadService) CancelTask(taskId uint) {
	chVal, ok := channelMap.Load(taskId)

	taskInfo, _ := s.TaskRepo.GetDetail(taskId)
	if taskInfo.ID > 0 && taskInfo.Status == consts.Created {
		s.TaskRepo.SetFailed(taskInfo)
	}

	if !ok || chVal == nil {
		return
	}

	channelMap.Delete(taskId)

	ch := chVal.(chan int)
	if ch != nil {
		if !_channelUtils.IsChanClose(ch) {
			ch <- 1
		}

		ch = nil
	}

	return
}

func (s *DownloadService) RestartTask(po agentModel.Task) (ret bool) {
	s.CancelTask(po.ID)

	s.StartTask(po)

	s.TaskRepo.AddRetry(po)

	return
}

func (s *DownloadService) RemoveTask(req v1.DownloadReq) {
	s.TaskRepo.Delete(uint(req.ZentaoTask))

	return
}

func (s *DownloadService) isEmpty() bool {
	length := 0

	channelMap.Range(func(key, value interface{}) bool {
		length++
		return true
	})

	return length == 0
}
