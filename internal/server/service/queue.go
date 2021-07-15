package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"strings"
)

type QueueService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`

	TaskService    *TaskService    `inject:""`
	HistoryService *HistoryService `inject:""`
}

func NewQueueService() *QueueService {
	return &QueueService{}
}

func (s QueueService) GenerateFromTask(task *model.Task) (count int) {
	s.removeOldQueuesByTask(task.ID)

	if task.BuildType == consts.AutoSelenium {
		count = s.GenerateSeleniumQueuesFromTask(task)
	} else if task.BuildType == consts.AutoAppium {
		count = s.GenerateAppiumQueuesFromTask(task)
	}

	return
}

func (s QueueService) GenerateSeleniumQueuesFromTask(task *model.Task) (count int) {
	envs := task.Environments
	if len(envs) == 0 {
		return
	}

	var groupId uint
	if task.GroupId != 0 {
		groupId = task.GroupId
	} else {
		groupId = task.ID
	}

	for _, env := range envs {
		osCategory := env.OsCategory
		osType := env.OsType
		osLang := env.OsLang

		queue := model.NewQueue(
			task.BuildType, groupId, task.ID, task.Priority,
			osCategory, osType, osLang,
			task.ScriptUrl, task.ScmAddress, task.ScmAccount, task.ScmPassword,
			task.ResultFiles, task.KeepResultFiles, task.Name, task.UserName,
			"", "", task.BuildCommands, task.EnvVars,
			task.BrowserType, task.BrowserVersion,
		)

		s.QueueRepo.Save(&queue)
		s.HistoryService.Create(consts.Queue, queue.ID, consts.ProgressCreated, "")
		count++
	}

	return
}

func (s QueueService) GenerateAppiumQueuesFromTask(task *model.Task) (count int) {
	if len(task.Serials) == 0 {
		return
	}

	var groupId uint
	if task.GroupId != 0 {
		groupId = task.GroupId
	} else {
		groupId = task.ID
	}

	serials := strings.Split(task.Serials, ",")
	for _, serial := range serials {
		serial = strings.TrimSpace(serial)
		if serial == "" {
			continue
		}

		device := s.DeviceRepo.GetBySerial(serial)
		if device.ID != 0 {
			queue := model.NewQueue(task.BuildType, groupId, task.ID, task.Priority,
				"", "", "",
				task.ScriptUrl, task.ScmAddress, task.ScmAccount, task.ScmPassword,
				task.ResultFiles, task.KeepResultFiles, task.Name, task.UserName,
				serial, task.AppUrl, task.BuildCommands, task.EnvVars,
				task.BrowserType, task.BrowserVersion)

			s.QueueRepo.Save(&queue)
			s.HistoryService.Create(consts.Queue, queue.ID, consts.ProgressCreated, "")
			count++
		}
	}

	return
}

// SaveResult not just update queue status, but also update parent task
func (s QueueService) SaveResult(queueId uint, progress consts.BuildProgress, status consts.BuildStatus) {
	queue := s.QueueRepo.GetQueue(queueId)

	s.QueueRepo.SetQueueStatus(queueId, progress, status)
	s.TaskService.SetTaskStatus(queue.TaskId)

	s.HistoryService.Create(consts.Queue, queueId, progress, status.ToString())
}

func (s QueueService) removeOldQueuesByTask(taskId uint) {
	s.QueueRepo.RemoveOldQueuesByTask(taskId)
}
