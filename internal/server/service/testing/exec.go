package testing

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
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
	SeleniumService  *SeleniumService                `inject:""`
	AppiumService    *AppiumService                  `inject:""`
	UnitService      *UnitService                    `inject:""`
	HostService      *serverService.HostService      `inject:""`
	HistoryService   *serverService.HistoryService   `inject:""`
	VmCommonService  *serverService.VmCommonService  `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s ExecService) QueryForExec() {
	queuesToBuild := s.QueueRepo.QueryForExec()
	for _, queue := range queuesToBuild {
		s.CheckAndCall(queue)
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

func (s ExecService) QueryForRetry() {
	queues := s.QueueRepo.QueryForRetry()

	for _, queue := range queues {
		s.CheckAndCall(queue)
	}
}

func (s ExecService) CheckAndCall(queue model.Queue) {
	if queue.BuildType == consts.AutoSelenium {
		s.CheckAndCallSeleniumTest(queue)
	} else if queue.BuildType == consts.AutoAppium {
		s.CheckAndCallAppiumTest(queue)
	} else if queue.BuildType == consts.UnitJunit || queue.BuildType == consts.UnitTestNG {
		s.CheckAndCallUnitTest(queue)
	}
}

func (s ExecService) CheckAndCallSeleniumTest(queue model.Queue) {
	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if queue.Progress == consts.ProgressLaunchVm { // run if vm launched
		vmId := queue.VmId
		vm := s.VmRepo.GetById(vmId)

		if vm.Status == consts.VmReady { // begin to run if vm ready
			result := s.SeleniumService.Run(queue)

			if result.IsSuccess() {
				s.QueueRepo.Run(queue)
				s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")
				s.WebSocketService.UpdateTask(queue.TaskId, "success to run selenium queue")

				newTaskProgress = consts.ProgressRunning
			} else {
				s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
			}
		}
	} else {
		s.QueueRepo.Retry(queue)

		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueueByVm(queue)
		if found {
			// create kvm
			vmService := s.VmCommonService.GetVmService(hostId)
			result := vmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
			if result.IsSuccess() { // success to create
				newTaskProgress = consts.ProgressLaunchVm
			} else {
				newTaskProgress = consts.ProgressCreateVmFail
			}
		} else {
			// only pending new queue
			if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressPendingRes {
				s.QueueRepo.Pending(queue.ID) // pending
				newTaskProgress = consts.ProgressPendingRes
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
		rpcResult := s.AppiumService.Run(queue)

		if rpcResult.IsSuccess() {
			s.QueueRepo.Run(queue) // start
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")
			s.WebSocketService.UpdateTask(queue.TaskId, "success to run appium queue")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
		}
	} else {
		// only pending new queue
		if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressPendingRes {
			s.QueueRepo.Pending(queue.ID) // pending
			newTaskProgress = consts.ProgressPendingRes
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

	hostId, found := s.HostService.GetValidForQueueByContainer(queue)
	if found {
		host := s.HostRepo.Get(hostId)

		result := s.UnitService.Run(queue, host)

		if result.IsSuccess() {
			s.QueueRepo.Run(queue)
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")
			s.WebSocketService.UpdateTask(queue.TaskId, "success to run unit queue")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
		}
	} else {
		// only pending new queue
		if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressPendingRes {
			s.QueueRepo.Pending(queue.ID) // pending
			newTaskProgress = consts.ProgressPendingRes
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}
