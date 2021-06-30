package service

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type ExecService struct {
	ExecRepo   *repo.ExecRepo   `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`
	TaskRepo   *repo.TaskRepo   `inject:""`
	DeviceRepo *repo.DeviceRepo `inject:""`
	VmRepo     *repo.VmRepo     `inject:""`

	DeviceService   *DeviceService   `inject:""`
	SeleniumService *SeleniumService `inject:""`
	AppiumService   *AppiumService   `inject:""`
	TaskService     *TaskService     `inject:""`
	HostService     *HostService     `inject:""`

	VmService VmService `inject:""`
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
	if queue.BuildType == commConst.AutoSelenium {
		s.CheckAndCallSeleniumTest(queue)
	} else if queue.BuildType == commConst.AutoAppium {
		s.CheckAndCallAppiumTest(queue)
	}
}

func (s ExecService) CheckAndCallSeleniumTest(queue model.Queue) {
	originalProgress := queue.Progress
	var newTaskProgress commConst.BuildProgress

	if queue.Progress == commConst.ProgressCreated ||
		queue.Progress == commConst.ProgressPending {

		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueue(queue)
		if found {
			// create kvm
			result := s.VmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
			if result.IsSuccess() { // success to create
				newTaskProgress = commConst.ProgressInProgress
			} else {
				newTaskProgress = commConst.ProgressPending
			}
		}

	} else if queue.Progress == commConst.ProgressLaunchVm {
		vmId := queue.VmId
		vm := s.VmRepo.GetById(vmId)

		if vm.Status == commConst.VmActive { // find ready vm, begin to run test
			result := s.SeleniumService.Start(queue)

			if result.IsSuccess() {
				s.QueueRepo.Start(queue)
				newTaskProgress = commConst.ProgressInProgress
			} else { // busy, pending
				s.QueueRepo.Pending(queue.ID)
				newTaskProgress = commConst.ProgressPending
			}
		}
	}

	if originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
	}
}

func (s ExecService) CheckAndCallAppiumTest(queue model.Queue) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	originalProgress := queue.Progress
	var newProgress commConst.BuildProgress

	if s.DeviceService.IsDeviceReady(device) {
		rpcResult := s.AppiumService.Start(queue)

		if rpcResult.IsSuccess() {
			s.QueueRepo.Start(queue) // start
			newProgress = commConst.ProgressInProgress
		} else {
			s.QueueRepo.Pending(queue.ID) // pending
			newProgress = commConst.ProgressPending
		}
	} else {
		s.QueueRepo.Pending(queue.ID) // pending
		newProgress = commConst.ProgressPending
	}

	if originalProgress != newProgress { // progress changed
		s.TaskService.SetProgress(queue.TaskId, newProgress)
	}
}

func (s ExecService) SetTimeout() {
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
