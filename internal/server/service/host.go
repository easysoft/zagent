package service

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_intUtils "github.com/easysoft/zagent/internal/pkg/lib/int"
	bizCasbin "github.com/easysoft/zagent/internal/server/biz/casbin"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type HostService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	CasbinService *bizCasbin.CasbinService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s HostService) Register(host commDomain.Host) (result _domain.RpcResp) {
	hostPo, err := s.HostRepo.Register(host)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", host.Ip))
	}

	s.updateVmsStatus(host, hostPo.ID)

	return
}

func (s HostService) GetValidForQueue(queue model.Queue) (hostId, backingId uint) {
	imageIds1 := s.BackingRepo.QueryByOs(queue.OsPlatform, queue.OsType, queue.OsLang)
	imageIds2 := s.BackingRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVersion)

	images := make([]int, 0)
	for id := range imageIds1 {
		if _intUtils.FindInArr(id, imageIds2) {
			images = append(images, id)
		}
	}

	if len(images) == 0 {
		return
	}

	hostIds := s.getIdleHost()
	if len(hostIds) == 0 {
		return
	}

	hostId, backingId = s.HostRepo.QueryByImages(images, hostIds)

	return
}

func (s HostService) getIdleHost() (ids []int) {
	// keys: hostId, vmCount
	hostToVmCountList := s.HostRepo.QueryIdle(commConst.MaxVmOnHost)

	hostIds := make([]int, 0)
	for _, mp := range hostToVmCountList {
		hostId := mp["hostId"]
		hostIds = append(hostIds, hostId)
	}

	return hostIds
}

func (s HostService) updateVmsStatus(host commDomain.Host, hostId uint) {
	vmNames := make([]string, 0)
	runningVms, destroyVms, unknownVms := s.getVmsByStatus(host, vmNames)

	if len(runningVms) > 0 {
		s.VmRepo.UpdateStatusByNames(runningVms, commConst.VmRunning)
	}
	if len(destroyVms) > 0 {
		s.VmRepo.UpdateStatusByNames(destroyVms, commConst.VmDestroy)
	}
	if len(unknownVms) > 0 {
		s.VmRepo.UpdateStatusByNames(unknownVms, commConst.VmUnknown)
	}

	// destroy vms already removed on agent side
	s.VmRepo.DestroyMissedVmsStatus(vmNames, hostId)

	return
}

func (s HostService) getVmsByStatus(host commDomain.Host, vmNames []string) (runningVms, destroyVms, unknownVms []string) {
	vms := host.Vms

	for _, vm := range vms {
		name := vm.Name
		status := vm.Status
		vmNames = append(vmNames, name)

		if status == commConst.VmRunning {
			runningVms = append(runningVms, name)
		} else if status == commConst.VmDestroy {
			destroyVms = append(destroyVms, name)
		} else if status == commConst.VmUnknown {
			unknownVms = append(unknownVms, name)
		}
	}

	return
}