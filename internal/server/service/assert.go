package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type AssertService struct {
	HostRepo  *repo.HostRepo  `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`

	HistoryService   *HistoryService                 `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewAssertService() *AssertService {
	return &AssertService{}
}

func (s AssertService) RegisterHost(host domain.HostNode) (result bool) {
	po := model.HostFromDomain(host)
	hostPo, err := s.HostRepo.Register(po)
	if err == nil {
		result = true
	}

	s.updateVmsStatus(host, hostPo.ID)

	return
}

func (s AssertService) RegisterVm(vmObj domain.Vm) (result bool) {
	vm, statusChanged, err := s.VmRepo.Register(vmObj)
	if err == nil {
		result = true
	}

	if statusChanged {
		queue := s.QueueRepo.GetByVmId(vm.ID)

		if vm.Status == consts.VmReady {
			s.QueueRepo.ResReady(queue.ID)
		}

		s.HistoryService.Create(consts.Vm, vm.ID, queue.ID, "", vm.Status.ToString())
		s.WebSocketService.UpdateTask(queue.TaskId, "vm ready")
	}

	return
}

func (s AssertService) updateVmsStatus(host domain.HostNode, hostId uint) {
	runningVms, shutOffVms, unknownVms, vmNames := s.getVmsByStatus(host)

	// only 3 kind of status from host register
	if len(runningVms) > 0 {
		// should not update vm status that is active like launch, ready etc.
		//s.VmRepo.UpdateStatusByNames(runningVms, consts.VmRunning)
	}
	if len(shutOffVms) > 0 {
		s.VmRepo.UpdateStatusByNames(shutOffVms, consts.VmShutOff)
	}
	if len(unknownVms) > 0 {
		s.VmRepo.UpdateStatusByNames(unknownVms, consts.VmUnknown)
	}

	// destroy timeout vms
	s.VmRepo.DestroyTimeoutVms()

	// destroy vms already removed by host agent
	s.VmRepo.DestroyMissedVmsStatus(vmNames, hostId)

	return
}

func (s AssertService) getVmsByStatus(host domain.HostNode) (
	runningVms, shutOffVms, unknownVms, vmNames []string) {
	vms := host.Vms

	for _, vm := range vms {
		name := vm.Name
		status := vm.Status
		vmNames = append(vmNames, name)

		if status == consts.VmRunning {
			runningVms = append(runningVms, name)
		} else if status == consts.VmShutOff {
			shutOffVms = append(shutOffVms, name)
		} else if status == consts.VmUnknown {
			unknownVms = append(unknownVms, name)
		}
	}

	return
}
