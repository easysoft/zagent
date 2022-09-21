package testing

import (
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	serverService "github.com/easysoft/zv/internal/server/service"
	commonService "github.com/easysoft/zv/internal/server/service/common"
)

type ExecService struct {
	ExecRepo   *repo.ExecRepo   `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`
	TaskRepo   *repo.TaskRepo   `inject:""`
	DeviceRepo *repo.DeviceRepo `inject:""`
	VmRepo     *repo.VmRepo     `inject:""`
	HostRepo   *repo.HostRepo   `inject:""`

	DeviceService    *serverService.DeviceService    `inject:""`
	TaskService      *serverService.TaskService      `inject:""`
	QueueService     *serverService.QueueService     `inject:""`
	SeleniumService  *serverService.SeleniumService  `inject:""`
	AppiumService    *serverService.AppiumService    `inject:""`
	UnitService      *serverService.UnitService      `inject:""`
	HostService      *serverService.HostService      `inject:""`
	HistoryService   *serverService.HistoryService   `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`

	KvmNativeService         *serverService.NativeKvmService         `inject:""`
	HuaweiCloudVmService     *serverService.HuaweiCloudVmService     `inject:""`
	HuaweiCloudDockerService *serverService.HuaweiCloudDockerService `inject:""`

	FacadeService *serverService.FacadeService `inject:""`
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s ExecService) QueryForExec() {
	queuesToBuild := s.QueueRepo.QueryForExec()
	for _, queue := range queuesToBuild {
		s.CheckAndCall(queue)
		s.WebSocketService.UpdateTask(queue.TaskId, "CheckAndCall Test")
	}
}

func (s ExecService) QueryForRetry() {
	queues := s.QueueRepo.QueryForRetry()

	for _, queue := range queues {
		s.CheckAndCall(queue)
		s.WebSocketService.UpdateTask(queue.TaskId, "CheckAndCall Test")
	}
}

func (s ExecService) QueryForTimeout() {
	queues := s.QueueRepo.QueryForTimeout()

	for _, queue := range queues {
		s.QueueRepo.Timeout(queue.ID)
		s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressTimeout, "")
		s.WebSocketService.UpdateTask(queue.TaskId, "set queue timeout")
	}
}

func (s ExecService) CheckAndCall(queue model.Queue) {
	if queue.BuildType == consts.SeleniumTest {
		s.CheckAndCallSeleniumTest(queue)
	} else if queue.BuildType == consts.AppiumTest {
		s.CheckAndCallAppiumTest(queue)
	} else if queue.BuildType == consts.UnitTest {
		s.CheckAndCallUnitTest(queue)
	}
}

func (s ExecService) CheckAndCallSeleniumTest(queue model.Queue) {
	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if queue.Progress == consts.ProgressResReady { // run if vm ready
		result := s.FacadeService.RunSeleniumTest(queue)

		if result.IsSuccess() {
			s.QueueRepo.Run(queue)
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunFail, consts.StatusFail.ToString())
		}
	} else {
		s.QueueRepo.Retry(queue)

		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueueByVm(queue)

		if found {
			// create kvm
			result := s.FacadeService.Create(hostId, backingId, tmplId, queue.ID)

			if result.IsSuccess() { // success to create
				newTaskProgress = consts.ProgressResLaunched
			} else {
				newTaskProgress = consts.ProgressResFailed
			}
		} else {
			// only pending new queue
			if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressResPending {
				s.QueueRepo.ResPending(queue.ID) // pending
				newTaskProgress = consts.ProgressResPending
			}
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}

func (s ExecService) CheckAndCallAppiumTest(queue model.Queue) {
	s.QueueRepo.Retry(queue)

	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if s.DeviceService.IsDeviceReady(device) {
		remoteResult := s.FacadeService.RunAppiumTest(queue)

		if remoteResult.IsSuccess() {
			s.QueueRepo.Run(queue) // start
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
		}
	} else {
		// only pending new queue
		if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressResPending {
			s.QueueRepo.ResPending(queue.ID) // pending
			newTaskProgress = consts.ProgressResPending
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // progress changed
		s.TaskService.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}

func (s ExecService) CheckAndCallUnitTest(queue model.Queue) {
	s.QueueRepo.Retry(queue)

	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	hostId, found := s.HostService.GetValidForQueueByDocker(queue)
	if found {
		host := s.HostRepo.Get(hostId)

		result := s.FacadeService.RunUnitTest(queue, host)

		if result.IsSuccess() {
			s.QueueRepo.Run(queue)
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
		}
	} else {
		// only pending new queue
		if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressResPending {
			s.QueueRepo.ResPending(queue.ID) // pending
			newTaskProgress = consts.ProgressResPending
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}
