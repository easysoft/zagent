package hostAgentService

import (
	"errors"
	"fmt"
	"sync"
	"time"

	agentModel "github.com/easysoft/zagent/internal/host/model"
	kvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/job"
	"github.com/gofrs/uuid"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

type SnapService struct {
	SyncHeartbeatMap sync.Map
	TimeStamp        int64

	LibvirtService *kvmService.LibvirtService `inject:""`
	QemuService    *kvmService.QemuService    `inject:""`
	KvmService     *kvmService.KvmService     `inject:""`

	TaskRepo             *hostRepo.TaskRepo    `inject:""`
	TaskService          *TaskService          `inject:""`
	AsyncExecutorService *AsyncExecutorService `inject:""`
}

func NewSnapService() *SnapService {
	s := SnapService{}
	s.TimeStamp = time.Now().Unix()

	return &s
}

// ListSnap 列出快照
func (s *SnapService) ListSnap(vm string) (ret []v1.SnapItemResp, err error) {
	ret = s.LibvirtService.GetVmSnaps(vm)

	return
}

func (s *SnapService) AddTasks(req []v1.SnapTaskReq) (err error) {
	for _, item := range req {
		if item.Vm == "" || item.Name == "" {
			err = errors.New("vm OR item is empty")
			return
		}
		po := agentModel.Task{
			Vm:     item.Vm,
			Name:   item.Name,
			Task:   item.Task,
			Type:   item.Type,
			Retry:  1,
			Status: consts.Created,
		}

		existTask, _ := s.TaskRepo.GetActiveTaskByVm(item.Vm)

		if existTask.ID != 0 {
			err = errors.New("the same create snapshot task is running")
			return
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *SnapService) StartCreateSnapTask(po agentModel.Task) {
	ch := make(chan int, 1)
	channelMap.Store(po.ID, ch)

	go func() {
		s.TaskRepo.UpdateStatus(po.ID, "", 0.01, "", consts.Inprogress, "", true, false)

		// create ...
		finalStatus, result := s.createSnap(&po)

		s.TaskRepo.UpdateStatus(po.ID, "", 1, "", finalStatus, result, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		job.TaskStatus.Delete(po.ID)

		if ch != nil {
			channelMap.Delete(po.ID)
			close(ch)
		}
	}()
}

func (s *SnapService) StartRevertSnapTask(po agentModel.Task) {
	ch := make(chan int, 1)
	channelMap.Store(po.ID, ch)

	go func() {
		s.TaskRepo.UpdateStatus(po.ID, "", 0.01, "", consts.Inprogress, "", true, false)

		// create ...
		finalStatus, result := s.revertSnap(&po)

		s.TaskRepo.UpdateStatus(po.ID, "", 1, "", finalStatus, result, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		job.TaskStatus.Delete(po.ID)

		if ch != nil {
			channelMap.Delete(po.ID)
			close(ch)
		}
	}()
}

// createSnap 创建快照
func (s *SnapService) createSnap(po *agentModel.Task) (status consts.TaskStatus, result string) {
	uuidStr := uuid.Must(uuid.NewV4()).String()

	status, result = s.AsyncExecutorService.Exec(po, uuidStr, func(po *agentModel.Task) (status consts.TaskStatus, result string) {
		cmd := fmt.Sprintf("virsh snapshot-create-as %s %s --atomic -uuid-%s", po.Vm, po.Name, uuidStr)
		result, err := _shellUtils.ExeShell(cmd)
		if err != nil {
			result = err.Error()
		}

		status = consts.Completed
		if err != nil {
			_logUtils.Infof("create snap '%s' err, output %s, error %s", cmd, result, err.Error())
			status = consts.Failed
		}

		return
	})

	return
}

// revertSnap 回滚到快照
func (s *SnapService) revertSnap(po *agentModel.Task) (status consts.TaskStatus, result string) {
	status, result = s.AsyncExecutorService.Exec(po, "", func(po *agentModel.Task) (status consts.TaskStatus, result string) {
		cmd := fmt.Sprintf("virsh snapshot-revert %s %s --running", po.Vm, po.Name)
		result, err := _shellUtils.ExeShell(cmd)
		if err != nil {
			result = err.Error()
		}

		status = consts.Completed
		if err != nil {
			_logUtils.Infof("revert snap '%s' err, output %s, error %s", cmd, result, err.Error())
			status = consts.Failed
		}

		return
	})

	return
}

func (s *SnapService) RemoveSnap(req *v1.SnapTaskReq) (err error) {
	cmd := fmt.Sprintf("virsh snapshot-delete %s %s", req.Vm, req.Name)

	out, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Infof("remove snap '%s' err, output %s, error %s", cmd, out, err.Error())
		return
	}

	return
}

func (s *SnapService) RemoveTask(req v1.DownloadReq) {
	s.TaskRepo.Delete(uint(req.Task))

	return
}

func (s *SnapService) isEmpty() bool {
	length := 0

	channelMap.Range(func(key, value interface{}) bool {
		length++
		return true
	})

	return length == 0
}
