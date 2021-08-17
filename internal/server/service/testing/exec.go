package testing

import (
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	"strings"
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
	WebSocketService *commonService.WebSocketService `inject:""`

	KvmNativeService         *serverService.KvmNativeService         `inject:""`
	HuaweiCloudVmService     *serverService.HuaweiCloudVmService     `inject:""`
	HuaweiCloudDockerService *serverService.HuaweiCloudDockerService `inject:""`
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

	if queue.Progress == consts.ProgressResReady { // run if vm ready
		result := s.SeleniumService.RemoteRun(queue)

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
			host := s.HostRepo.Get(hostId)

			result := _domain.RpcResp{}
			if strings.Index(host.Platform.ToString(), consts.PlatformVm.ToString()) > -1 {
				if strings.Index(host.Platform.ToString(), consts.PlatformNative.ToString()) > -1 {
					result = s.KvmNativeService.CreateRemote(hostId, backingId, tmplId, queue.ID)
				} else if strings.Index(host.Platform.ToString(), consts.PlatformHuawei.ToString()) > -1 {
					result = s.HuaweiCloudVmService.CreateRemote(hostId, backingId, tmplId, queue.ID)
				}
			} else if strings.Index(host.Platform.ToString(), consts.PlatformDocker.ToString()) > -1 {
				if strings.Index(host.Platform.ToString(), consts.PlatformHuawei.ToString()) > -1 {
					s.HuaweiCloudDockerService.DestroyRemote(queue.VmId, queue.ID)
				}
			}

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
		rpcResult := s.AppiumService.RemoteRun(queue)

		if rpcResult.IsSuccess() {
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

	hostId, found := s.HostService.GetValidForQueueByContainer(queue)
	if found {
		host := s.HostRepo.Get(hostId)

		result := s.UnitService.RemoteRun(queue, host)

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
