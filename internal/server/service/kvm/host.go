package kvmService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type HostService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	TmplRepo    *repo.TmplRepo    `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s HostService) Register(host domain.Host) (result _domain.RpcResp) {
	po := model.HostFromDomain(host)
	hostPo, err := s.HostRepo.Register(po)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", host.Ip))
	}

	s.updateVmsStatus(host, hostPo.ID)

	result.Success("")

	return
}

func (s HostService) GetValidForQueue(queue model.Queue) (hostId, backingId, tmplId uint, found bool) {
	backingIdsByBrowser := s.BackingRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVersion)
	backingIds, found := s.BackingRepo.QueryByOs(queue.OsPlatform, queue.OsType, queue.OsLang, backingIdsByBrowser)
	if !found {
		return
	}

	busyHostIds := s.getBusyHosts()
	hostId, backingId = s.HostRepo.QueryByBackings(backingIds, busyHostIds)

	tmplId, found = s.TmplRepo.QueryByOs(queue.OsPlatform, queue.OsType, queue.OsLang)

	return
}

func (s HostService) getBusyHosts() (ids []uint) {
	// keys: hostId, vmCount
	hostToVmCountList := s.HostRepo.QueryBusy()

	hostIds := make([]uint, 0)
	for _, mp := range hostToVmCountList {
		hostId := mp["hostId"]
		hostIds = append(hostIds, hostId)
	}

	return hostIds
}

func (s HostService) updateVmsStatus(host domain.Host, hostId uint) {
	vmNames := make([]string, 0)
	runningVms, shutOffVms, unknownVms := s.getVmsByStatus(host, vmNames)

	if len(runningVms) > 0 {
		s.VmRepo.UpdateStatusByNames(runningVms, consts.VmRunning)
	}
	if len(shutOffVms) > 0 {
		s.VmRepo.UpdateStatusByNames(shutOffVms, consts.VmShutOff)
	}
	if len(unknownVms) > 0 {
		s.VmRepo.UpdateStatusByNames(unknownVms, consts.VmUnknown)
	}

	// destroy vms already removed on agent side
	s.VmRepo.DestroyMissedVmsStatus(vmNames, hostId)

	return
}

func (s HostService) getVmsByStatus(host domain.Host, vmNames []string) (runningVms, shutOffVms, unknownVms []string) {
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
