package testing

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	kvmService "github.com/easysoft/zagent/internal/server/service/kvm"
)

type ExecService struct {
	ExecRepo   *repo.ExecRepo   `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`
	TaskRepo   *repo.TaskRepo   `inject:""`
	DeviceRepo *repo.DeviceRepo `inject:""`
	VmRepo     *repo.VmRepo     `inject:""`

	DeviceService   *serverService.DeviceService `inject:""`
	TaskService     *serverService.TaskService   `inject:""`
	SeleniumService *SeleniumService             `inject:""`
	AppiumService   *AppiumService               `inject:""`
	HostService     *kvmService.HostService      `inject:""`

	VmService kvmService.VmService `inject:""`
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s ExecService) CheckExec() {
	queuesToBuild := s.QueueRepo.QueryForExec()
	for _, queue := range queuesToBuild {
		s.CheckAndCall(queue)
	}
}

func (s ExecService) CheckAndCall(queue model.Queue) {
	if queue.BuildType == consts.AutoSelenium {
		s.CheckAndCallSeleniumTest(queue)
	} else if queue.BuildType == consts.AutoAppium {
		s.CheckAndCallAppiumTest(queue)
	}
}

func (s ExecService) CheckAndCallSeleniumTest(queue model.Queue) {
	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if queue.Progress == consts.ProgressCreated ||
		queue.Progress == consts.ProgressPendingRes {

		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueue(queue)
		if found {
			// create kvm
			result := s.VmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
			if result.IsSuccess() { // success to create
				newTaskProgress = consts.ProgressRunning
			} else {
				newTaskProgress = consts.ProgressPendingRes
			}
		} else {
			s.QueueRepo.Pending(queue.ID) // pending
			newTaskProgress = consts.ProgressPendingRes
		}

	} else if queue.Progress == consts.ProgressLaunchVm {
		vmId := queue.VmId
		vm := s.VmRepo.GetById(vmId)

		if vm.Status == consts.VmReady { // find ready vm, begin to run test
			result := s.SeleniumService.Start(queue)

			if result.IsSuccess() {
				s.QueueRepo.Start(queue)
				newTaskProgress = consts.ProgressRunning
			}
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
	}
}

func (s ExecService) CheckAndCallAppiumTest(queue model.Queue) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if s.DeviceService.IsDeviceReady(device) {
		rpcResult := s.AppiumService.Start(queue)

		if rpcResult.IsSuccess() {
			s.QueueRepo.Start(queue) // start
			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueRepo.SetQueueStatus(queue.ID, consts.ProgressAppiumServiceFail, consts.StatusFail)
			newTaskProgress = consts.ProgressPendingRes
		}
	} else {
		s.QueueRepo.Pending(queue.ID) // pending
		newTaskProgress = consts.ProgressPendingRes
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // progress changed
		s.TaskService.SetProgress(queue.TaskId, newTaskProgress)
	}
}

func (s ExecService) CheckTimeout() {
	queues := s.QueueRepo.QueryTimeout()

	for _, queue := range queues {
		s.QueueRepo.SetTimeout(queue.ID)
	}
}

func (s ExecService) RetryTimeoutOrFailed() {
	queues := s.QueueRepo.QueryTimeoutOrFailedForRetry()

	for _, queue := range queues {
		s.CheckAndCall(queue)
	}
}
