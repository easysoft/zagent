package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	"github.com/mitchellh/mapstructure"
)

type KvmNativeService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	IsoRepo     *repo.IsoRepo     `inject:""`
	TmplRepo    *repo.TmplRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService *VmCommonService          `inject:""`
	RpcService      *commonService.RpcService `inject:""`
	HistoryService  *HistoryService           `inject:""`
}

func (s KvmNativeService) CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	host := s.HostRepo.Get(hostId)
	backing := s.BackingRepo.Get(backingId)
	sysIso := s.IsoRepo.Get(backing.SysIsoId)
	sysIsoPath := sysIso.Path

	driverIsoPath := ""
	if backing.OsCategory == consts.Windows {
		driverIso := s.IsoRepo.Get(backing.DriverIsoId)
		driverIsoPath = driverIso.Path
	}

	tmpl := s.TmplRepo.Get(tmplId)

	macAddress := s.VmCommonService.genValidMacAddress() // get a unique mac address

	vm := model.Vm{
		HostId: host.ID, HostName: host.Name,
		Status:     consts.VmCreated,
		OsCategory: backing.OsCategory,
		OsType:     backing.OsType,
		OsVersion:  backing.OsVersion,
		OsLang:     backing.OsLang,
		BackingId:  backing.ID,

		BackingPath: backing.Path,
		MacAddress:  macAddress,
		TmplId:      tmpl.ID, TmplName: tmpl.Name,
		DiskSize: backing.SuggestDiskSize, MemorySize: backing.SuggestMemorySize,
		CdromSys: sysIsoPath, CdromDriver: driverIsoPath,
	}

	if backing.SuggestDiskSize == 0 {
		if vm.OsCategory == consts.Windows {
			vm.DiskSize = consts.DiskSizeWindows
		} else if vm.OsCategory == consts.Linux {
			vm.DiskSize = consts.DiskSizeLinux
		} else {
			vm.DiskSize = consts.DiskSizeDefault
		}
	}

	s.VmRepo.Save(&vm) // save vm to db, then update name with id
	vm.Name = s.VmCommonService.genVmName(backing, vm.ID)
	s.VmRepo.UpdateVmName(vm)

	kvmReq := model.GenKvmReq(vm)
	result = s.RpcService.CreateVm(host.Ip, host.Port, kvmReq)

	vmInResp := domain.Vm{}
	if result.IsSuccess() { // success to create vm
		mp := result.Payload.(map[string]interface{})
		mapstructure.Decode(mp, &vmInResp)
	}
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID,
		vmInResp.VncAddress, vmInResp.ImagePath, vmInResp.BackingPath)

	return
}

func (s KvmNativeService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	req := domain.KvmReq{VmUniqueName: vm.Name}

	result = s.RpcService.DestroyVm(host.Ip, host.Port, req)
	var status consts.VmStatus
	if result.IsSuccess() {
		status = consts.VmDestroy
	} else {
		status = consts.VmFailDestroy
	}
	s.VmRepo.UpdateStatusByNames([]string{vm.Name}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
