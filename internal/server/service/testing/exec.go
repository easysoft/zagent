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

	DeviceService   *serverService.DeviceService  `inject:""`
	TaskService     *serverService.TaskService    `inject:""`
	QueueService    *serverService.QueueService   `inject:""`
	SeleniumService *SeleniumService              `inject:""`
	AppiumService   *AppiumService                `inject:""`
	HostService     *kvmService.HostService       `inject:""`
	HistoryService  *serverService.HistoryService `inject:""`

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

	if queue.Progress == consts.ProgressCreated || queue.Progress == consts.ProgressPendingRes { // new queue
		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueue(queue)
		if found {
			// create kvm
			result := s.VmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
			if result.IsSuccess() { // success to create
				newTaskProgress = consts.ProgressLaunchVm
			} else {
				newTaskProgress = consts.ProgressCreateVmFail
			}
		} else {
			s.QueueRepo.Pending(queue.ID) // pending
			newTaskProgress = consts.ProgressPendingRes
		}

	} else if queue.Progress == consts.ProgressTimeout || queue.Status == consts.StatusFail { // retry queue
		// looking for valid host
		hostId, backingId, tmplId, found := s.HostService.GetValidForQueue(queue)
		if found {
			// create kvm
			result := s.VmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
			if result.IsSuccess() { // success to create
				newTaskProgress = consts.ProgressRunning
			}
		} // different from new queue, no need to update progress to 'ProgressPendingRes' when retry

	} else if queue.Progress == consts.ProgressLaunchVm { // vm launching
		vmId := queue.VmId
		vm := s.VmRepo.GetById(vmId)

		if vm.Status == consts.VmReady { // find ready vm, begin to run test
			result := s.SeleniumService.Run(queue)

			if result.IsSuccess() {
				s.QueueRepo.Run(queue)
				s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")

				newTaskProgress = consts.ProgressRunning
			} else {
				s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
			}
		}
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // queue's progress changed, update parent task
		s.TaskRepo.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}

func (s ExecService) CheckAndCallAppiumTest(queue model.Queue) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	originalProgress := queue.Progress
	var newTaskProgress consts.BuildProgress

	if s.DeviceService.IsDeviceReady(device) {
		rpcResult := s.AppiumService.Run(queue)

		if rpcResult.IsSuccess() {
			s.QueueRepo.Run(queue) // start
			s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressRunning, "")

			newTaskProgress = consts.ProgressRunning
		} else {
			s.QueueService.SaveResult(queue.ID, consts.ProgressRunFail, consts.StatusFail)
		}
	} else {
		s.QueueRepo.Pending(queue.ID) // pending
		newTaskProgress = consts.ProgressPendingRes
	}

	if newTaskProgress != "" && originalProgress != newTaskProgress { // progress changed
		s.TaskService.SetProgress(queue.TaskId, newTaskProgress)
		s.HistoryService.Create(consts.Task, queue.TaskId, 0, newTaskProgress, "")
	}
}

func (s ExecService) CheckTimeout() {
	queues := s.QueueRepo.QueryTimeout()

	for _, queue := range queues {
		s.QueueRepo.Timeout(queue.ID)
		s.HistoryService.Create(consts.Queue, queue.ID, queue.ID, consts.ProgressTimeout, "")
	}
}

func (s ExecService) CheckRetry() {
	queues := s.QueueRepo.QueryForRetry()

	for _, queue := range queues {
		s.CheckAndCall(queue)
	}
}
