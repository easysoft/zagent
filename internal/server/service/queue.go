package serverService

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"strings"
)

type QueueService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`

	TaskService *TaskService `inject:""`
}

func NewQueueService() *QueueService {
	return &QueueService{}
}

func (s QueueService) GenerateFromTask(task *model.Task) (count int) {
	s.CancelQueuesNotExec(task)

	if task.BuildType == commConst.AutoSelenium {
		count = s.GenerateSeleniumQueuesFromTask(task)
	} else if task.BuildType == commConst.AutoAppium {
		count = s.GenerateAppiumQueuesFromTask(task)
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
				serial, task.AppUrl, task.BuildCommands, task.EnvVars)

			s.QueueRepo.Save(&queue)
			count++
		}
	}

	s.QueueRepo.DeleteInSameGroup(task.GroupId, serials) // disable same serial queues in same group

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
			"", "", task.BuildCommands, task.EnvVars)

		s.QueueRepo.Save(&queue)
		count++
	}

	return
}

func (s QueueService) SetQueueResult(queueId uint, progress commConst.BuildProgress, status commConst.BuildStatus) {
	queue := s.QueueRepo.GetQueue(queueId)

	s.QueueRepo.SetQueueStatus(queueId, progress, status)
	s.TaskService.CheckCompleted(queue.TaskId)
}

func (s QueueService) CancelQueuesNotExec(task *model.Task) (count int) {
	s.QueueRepo.CancelQueuesNotExec(task.ID)

	return
}
