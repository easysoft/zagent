package serverService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type AssertService struct {
	HostRepo *repo.HostRepo `inject:""`
	VmRepo   *repo.VmRepo   `inject:""`
}

func NewAssertService() *AssertService {
	return &AssertService{}
}

func (s AssertService) RegisterHost(host domain.Host) (result _domain.RpcResp) {
	po := model.HostFromDomain(host)
	hostPo, err := s.HostRepo.Register(po)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", host.Ip))
	}

	s.updateVmsStatus(host, hostPo.ID)

	result.Pass("")

	return
}

func (s AssertService) RegisterVm(vm domain.Vm) (result _domain.RpcResp) {
	po := model.VmFromDomain(vm)

	err := s.VmRepo.Register(po)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", po.NodeIp))
	}

	result.Pass("")

	return
}

func (s AssertService) updateVmsStatus(host domain.Host, hostId uint) {
	runningVms, shutOffVms, unknownVms, vmNames := s.getVmsByStatus(host)

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

	// destroy vms already removed by host agent
	s.VmRepo.DestroyMissedVmsStatus(vmNames, hostId)

	return
}

func (s AssertService) getVmsByStatus(host domain.Host) (
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
