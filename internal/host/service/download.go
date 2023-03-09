package hostAgentService

import (
	"errors"
	"os"
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
		if item.Url == "" {
			err = errors.New("url is empty")
			return
		}
		po := agentModel.Task{
			Url:    item.Url,
			Md5:    item.Md5,
			Task:   item.Task,
			Type:   consts.DownloadImage,
			Retry:  1,
			Status: consts.Created,
		}

		existTask, _ := s.TaskRepo.GetActiveTaskByMd5(item.Md5)

		if existTask.ID != 0 {
			err = errors.New("the same md5 task is downloading")
			return
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

		if po.Md5 == "" {
			downloadUtils.GetMd5FromRemote(&po, consts.DownloadDir)
		}
		//query same md5 task
		sameMd5Task, _ := s.TaskRepo.GetCompletedTaskByMd5(po.Md5)
		if sameMd5Task.ID > 0 {
			_, err := os.Stat(sameMd5Task.Path)
			if err == nil && os.IsExist(err) && downloadUtils.CheckMd5(sameMd5Task) {
				po.Path = sameMd5Task.Path
				filePath = downloadUtils.GetPath(po)
			}
		}

		s.TaskRepo.UpdateStatus(po.ID, filePath, 0.01, "", consts.Inprogress, "", true, false)

		finalStatus, existFile := downloadUtils.Start(&po, filePath, ch)
		if existFile != "" {
			filePath = existFile
		}

		s.TaskRepo.UpdateSpeed(po.ID, po.Speed)
		s.TaskRepo.UpdateStatus(po.ID, filePath, 1, "", finalStatus, "", false, true)

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
	taskInfo, _ := s.TaskRepo.GetDetail(taskId)

	if taskInfo.ID > 0 {
		s.TaskRepo.SetCanceled(taskInfo)
	}

	s.StopTask(taskId)
}

func (s *DownloadService) StopTask(taskId uint) {
	chVal, ok := channelMap.Load(taskId)

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
}

func (s *DownloadService) RestartTask(po agentModel.Task) (ret bool) {
	s.CancelTask(po.ID)

	s.StartTask(po)

	s.TaskRepo.AddRetry(po)

	return
}

func (s *DownloadService) RemoveTask(req v1.DownloadReq) {
	s.TaskRepo.Delete(uint(req.Task))

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
