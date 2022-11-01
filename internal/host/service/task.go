package hostAgentService

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	requestUtils "github.com/easysoft/zagent/internal/pkg/utils/request"
	downloadUtils "github.com/easysoft/zagent/pkg/lib/download"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
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
	if len(taskMap.InProgress) > 0 {
		runningTask := taskMap.InProgress[0]

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

	if toStartNewTask && len(taskMap.Created) > 0 {
		newTask := taskMap.Created[0]

		if newTask.TaskType == consts.DownloadImage {
			s.DownloadService.StartTask(newTask)
		} else {
			s.ExportService.StartTask(newTask)
		}

	}

	return
}

func (s *TaskService) ListTask() (ret v1.ListTaskResp, err error) {
	ret = v1.ListTaskResp{
		Created:    make([]agentModel.Task, 0),
		InProgress: make([]agentModel.Task, 0),
		Canceled:   make([]agentModel.Task, 0),
		Completed:  make([]agentModel.Task, 0),
		Failed:     make([]agentModel.Task, 0),
	}

	pos, _ := s.TaskRepo.Query()

	for _, po := range pos {
		status := po.Status
		if status == consts.Timeout || status == consts.Error {
			status = consts.InProgress
		}

		rate, ok := downloadUtils.TaskMap.Load(po.ID)
		if ok {
			po.CompletionRate = rate.(float64)
		}

		if status == "created" {
			ret.Created = append(ret.Created, po)
		} else if status == "inProgress" {
			ret.InProgress = append(ret.InProgress, po)
		} else if status == "canceled" {
			ret.Canceled = append(ret.Canceled, po)
		} else if status == "completed" {
			ret.Completed = append(ret.Completed, po)
		} else if status == "failed" {
			ret.Failed = append(ret.Failed, po)
		}
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
