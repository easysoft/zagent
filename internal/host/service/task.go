package hostAgentService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentModel "github.com/easysoft/zv/internal/host/model"
	hostRepo "github.com/easysoft/zv/internal/host/repo"
	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	consts "github.com/easysoft/zv/internal/pkg/const"
	requestUtils "github.com/easysoft/zv/internal/pkg/utils/request"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"time"
)

type TaskService struct {
	DownloadService *DownloadService `inject:""`
	ExportService   *ExportService   `inject:""`

	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CheckTask() (err error) {
	taskMap, _ := s.ListTask()

	toStartNewTask := false
	if len(taskMap[consts.InProgress]) > 0 {
		runningTask := taskMap[consts.InProgress][0]

		if runningTask.TaskType == consts.DownloadImage &&
			(s.IsError(runningTask) || s.IsTimeout(runningTask) && s.NeedRetry(runningTask)) {
			s.DownloadService.RestartTask(runningTask)
		} else {
			s.TaskRepo.SetFailed(runningTask)
			toStartNewTask = true
		}

	} else {
		toStartNewTask = true
	}

	if toStartNewTask && len(taskMap[consts.Created]) > 0 {
		newTask := taskMap[consts.Created][0]

		if newTask.TaskType == consts.DownloadImage {
			s.DownloadService.StartTask(newTask)
		} else {
			s.ExportService.StartTask(newTask)
		}

	}

	return
}

func (s *TaskService) ListTask() (ret map[consts.TaskStatus][]agentModel.Task, err error) {
	ret = map[consts.TaskStatus][]agentModel.Task{}

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

func (s *TaskService) SubmitResult(task agentModel.Task) (err error) {
	// only submit vm task
	if task.TaskType == consts.DownloadImage {
		url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/submitResult")

		data := v1.ExportVmResp{
			Backing:    task.Backing,
			Xml:        task.Xml,
			Status:     task.Status,
			ZentaoTask: task.ZentaoTask,
		}

		_, err = _httpUtils.Post(url, data)

	} else if task.TaskType == consts.ExportVm {
		url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/submitResult")

		data := v1.ExportVmResp{
			Backing:    task.Backing,
			Xml:        task.Xml,
			Status:     task.Status,
			ZentaoTask: task.ZentaoTask,
		}

		_, err = _httpUtils.Post(url, data)

	}

	return
}

func (s *TaskService) IsError(po agentModel.Task) bool {
	return po.Status == consts.Error
}

func (s *TaskService) IsTimeout(po agentModel.Task) bool {
	return time.Now().Unix()-po.StartTime.Unix() > consts.DownloadTimeout
}

func (s *TaskService) NeedRetry(po agentModel.Task) bool {
	return po.Retry <= consts.DownloadRetry
}
