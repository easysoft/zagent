package hostAgentService

import (
	"errors"
	"fmt"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	kvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/job"
	_channelUtils "github.com/easysoft/zagent/pkg/lib/channel"
	"github.com/gofrs/uuid"
	"github.com/libvirt/libvirt-go"
	"sync"
	"time"

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

	TaskRepo    *hostRepo.TaskRepo `inject:""`
	TaskService *TaskService       `inject:""`
}

func NewSnapService() *SnapService {
	s := SnapService{}
	s.TimeStamp = time.Now().Unix()

	return &s
}

// ListSnap 列出快照
func (s *SnapService) ListSnap(vm string) (ret []*libvirt.DomainSnapshot, err error) {
	ret = s.LibvirtService.GetVmSnaps(vm)

	return
}

func (s *SnapService) RemoveSnap(req *v1.SnapReq) (err error) {
	cmd := fmt.Sprintf("virsh snapshot-delete %s %s", req.Vm, req.Name)

	out, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Infof("remove snap '%s' err, output %s, error %s", cmd, out, err.Error())
		return
	}

	return
}

// RevertSnap 回滚到快照
func (s *SnapService) RevertSnap(req *v1.SnapReq) (err error) {
	cmd := fmt.Sprintf("virsh snapshot-revert %s %s", req.Vm, req.Name)

	out, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Infof("revert snap '%s' err, output %s, error %s", cmd, out, err.Error())
		return
	}

	return
}

func (s *SnapService) AddTasks(req []v1.SnapReq) (err error) {
	for _, item := range req {
		if item.Vm == "" || item.Name == "" {
			err = errors.New("vm OR item is empty")
			return
		}
		po := agentModel.Task{
			Vm:     item.Vm,
			Name:   item.Name,
			Task:   item.Task,
			Type:   consts.CreateSnap,
			Retry:  1,
			Status: consts.Created,
		}

		existTask, _ := s.TaskRepo.GetActiveTaskByVm(item.Vm)

		if existTask.ID != 0 {
			err = errors.New("the same vm task is downloading")
			return
		}

		s.TaskRepo.Save(&po)
	}

	return
}

func (s *SnapService) createSnap(po *agentModel.Task) (status consts.TaskStatus) {
	uuidStr := uuid.Must(uuid.NewV4()).String()

	ch := make(chan string)
	go func() {
		cmd := fmt.Sprintf("virsh snapshot-create-as %s %s -atomic -uuid-%s", po.Vm, po.Name, uuidStr)
		out, err := _shellUtils.ExeShell(cmd)

		statusMsg := consts.Completed.ToString()
		if err != nil {
			_logUtils.Infof("create snap '%s' err, output %s, error %s", cmd, out, err.Error())
			statusMsg = consts.Error.ToString()
		}

		ch <- statusMsg
	}()

	select {
	case val := <-ch:
		status = consts.TaskStatus(val)

	case <-time.After(consts.CreateSnapTimeout * time.Second):
		status = consts.Timeout
		_shellUtils.KillProcessByUUID(uuidStr)
	}

	return
}

func (s *SnapService) StartTask(po agentModel.Task) {
	ch := make(chan int, 1)
	channelMap.Store(po.ID, ch)

	go func() {
		s.TaskRepo.UpdateStatus(po.ID, "", 0.01, "", consts.Inprogress, true, false)

		// create ...
		finalStatus := s.createSnap(&po)

		s.TaskRepo.UpdateStatus(po.ID, "", 1, "", finalStatus, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		job.TaskStatus.Delete(po.ID)

		if ch != nil {
			channelMap.Delete(po.ID)
			close(ch)
		}
	}()
}

func (s *SnapService) CancelTask(taskId uint) {
	taskInfo, _ := s.TaskRepo.GetDetail(taskId)

	if taskInfo.ID > 0 {
		s.TaskRepo.SetCanceled(taskInfo)
	}

	s.stopTask(taskId)
}

func (s *SnapService) stopTask(taskId uint) {
	chVal, ok := channelMap.Load(taskId)

	if !ok || chVal == nil {
		return
	}

	channelMap.Delete(taskId)

	ch := chVal.(chan int)
	if ch != nil {
		if !_channelUtils.IsChanClose(ch) {
			ch <- 1
		}

		ch = nil
	}
}

func (s *SnapService) RestartTask(po agentModel.Task) (ret bool) {
	s.CancelTask(po.ID)

	s.StartTask(po)

	s.TaskRepo.AddRetry(po)

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
