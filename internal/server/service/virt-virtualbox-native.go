package serverService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	commonService "github.com/easysoft/zv/internal/server/service/common"
	_domain "github.com/easysoft/zv/pkg/domain"
	"github.com/mitchellh/mapstructure"
)

type VirtualboxCloudVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	CommonService   *CommonService               `inject:""`
	VmCommonService *VmCommonService             `inject:""`
	HistoryService  *HistoryService              `inject:""`
	RpcService      *commonService.RemoteService `inject:""`
}

func (s VirtualboxCloudVmService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	host := s.HostRepo.Get(hostId)
	backing := s.BackingRepo.Get(backingId)
	backing.Name = s.VmCommonService.genTmplName(backing)

	vm := model.Vm{
		HostId: host.ID, HostName: host.Name,
		Status:     consts.VmCreated,
		OsCategory: backing.OsCategory,
		OsType:     backing.OsType,
		OsVersion:  backing.OsVersion,
		OsLang:     backing.OsLang,
		BackingId:  backing.ID,
	}
	s.VmRepo.Save(&vm) // save vm to db, then update name with id
	vm.Name = s.VmCommonService.genVmName(backing, vm.ID)
	s.VmRepo.UpdateVmName(vm)

	req := model.GenVirtualBoxReq(vm, backing, host)
	result = s.RpcService.CreateVirtualBox(host.Ip, host.Port, req)

	vmInResp := v1.VirtualBoxResp{}
	if result.IsSuccess() { // success to create vm
		mp := result.Payload.(map[string]interface{})
		mapstructure.Decode(mp, &vmInResp)
	}
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID,
		vmInResp.VncPort, vm.VncUrl, vmInResp.ImagePath, vmInResp.BackingPath)
	return
}

func (s VirtualboxCloudVmService) DestroyRemote(vmId, queueId uint) (result _domain.RemoteResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	req := v1.VirtualBoxReq{VmUniqueName: vm.Name, VncPort: vm.VncPort}
	result = s.RpcService.DestroyVirtualBox(host.Ip, host.Port, req)

	if !result.IsSuccess() {
		status = consts.VmDestroyFail
	}
	s.VmRepo.UpdateStatusByNames([]string{vm.Name}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
