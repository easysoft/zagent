package hostAgentService

import (
	"fmt"
	"strconv"
	"time"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	downloadUtils "github.com/easysoft/zagent/internal/pkg/job"
	requestUtils "github.com/easysoft/zagent/internal/pkg/utils/request"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
)

type TaskService struct {
	DownloadService *DownloadService `inject:""`
	ExportService   *ExportService   `inject:""`
	SnapService     *SnapService     `inject:""`

	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CheckTask() (err error) {
	taskMap, _ := s.ListTask()

	toStartNewTask := false
	if len(taskMap.Inprogress) > 0 {
		runningTask := taskMap.Inprogress[0]

		if runningTask.Type == consts.DownloadImage {
			if s.IsError(runningTask) || s.IsTimeout(runningTask) || s.DownloadService.isEmpty() {
				if s.NeedRetry(runningTask) {
					s.DownloadService.RestartTask(runningTask)
				} else {
					s.TaskRepo.SetFailed(runningTask)
					toStartNewTask = true
				}
			}
		}

	} else {
		toStartNewTask = true
	}

	if toStartNewTask && len(taskMap.Created) > 0 {
		newTask := taskMap.Created[0]

		if newTask.Type == consts.DownloadImage {
			s.DownloadService.StartTask(newTask)
		} else if newTask.Type == consts.ExportVm {
			s.ExportService.StartTask(newTask)
		} else if newTask.Type == consts.CreateSnap {
			s.SnapService.StartCreateSnapTask(newTask)
		} else if newTask.Type == consts.RevertSnap {
			s.SnapService.StartRevertSnapTask(newTask)
		}
	}

	return
}

func (s *TaskService) ListTask() (ret v1.ListTaskResp, err error) {
	ret = v1.ListTaskResp{
		Created:    make([]agentModel.Task, 0),
		Inprogress: make([]agentModel.Task, 0),
		Canceled:   make([]agentModel.Task, 0),
		Completed:  make([]agentModel.Task, 0),
		Failed:     make([]agentModel.Task, 0),
	}

	pos, _ := s.TaskRepo.Query()

	for _, po := range pos {
		status := po.Status
		if po.Type == consts.DownloadImage && (status == consts.Timeout || status == consts.Error) { // only retry download task
			status = consts.Inprogress
		}

		completionRate, speed := downloadUtils.GetTaskStatus(downloadUtils.TaskStatus, po.ID)
		if completionRate > 0 {
			po.Rate = completionRate
		}
		if speed > 0 {
			po.Speed = speed
		}

		po.Rate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", po.Rate), 64)
		po.Speed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", po.Speed), 64)

		if status == consts.Created {
			ret.Created = append(ret.Created, po)
		} else if status == consts.Inprogress {
			ret.Inprogress = append(ret.Inprogress, po)
		} else if status == consts.Canceled {
			ret.Canceled = append(ret.Canceled, po)
		} else if status == consts.Completed {
			ret.Completed = append(ret.Completed, po)
		} else if status == consts.Failed {
			ret.Failed = append(ret.Failed, po)
		}
	}

	return
}

func (s *TaskService) SubmitResult(task agentModel.Task) (err error) {
	fmt.Println("=======submit task", task.Task, task.Status)
	// only submit vm task
	if task.Type == consts.DownloadImage {
		url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/submitResult")

		data := v1.ExportVmResp{
			Backing: task.Backing,
			Xml:     task.Xml,
			Status:  task.Status,
			Rate:    task.Rate,
			Speed:   task.Speed,
			Task:    task.Task,
		}

		_, err = _httpUtils.Post(url, data)

	} else if task.Type == consts.ExportVm {
		url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/submitResult")

		data := v1.ExportVmResp{
			Backing: task.Backing,
			Xml:     task.Xml,
			Status:  task.Status,
			Task:    task.Task,
		}

		_, err = _httpUtils.Post(url, data)

	} else if task.Type == consts.CreateSnap || task.Type == consts.RevertSnap {
		url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/submitResult")

		data := v1.ExportVmResp{
			Status: task.Status,
			Task:   task.Task,
		}

		_, err = _httpUtils.Post(url, data)
	}

	return
}

func (s *TaskService) IsError(po agentModel.Task) bool {
	return po.Status == consts.Error
}

func (s *TaskService) IsTimeout(po agentModel.Task) bool {
	dur := time.Now().Unix() - po.StartDate.Unix()
	//return dur > 3
	return po.Status == consts.Inprogress && dur > consts.DownloadImageTimeout
}

func (s *TaskService) NeedRetry(po agentModel.Task) bool {
	return po.Retry < consts.DownloadRetry
}
