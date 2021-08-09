package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/service/vendors"
)

type HuaweiCloudService struct {
	VmCommonService
}

func (s HuaweiCloudService) CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	host := s.HostRepo.Get(hostId)
	backing := s.BackingRepo.Get(backingId)

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
	vm.Name = s.genVmName(backing, vm.ID)
	s.VmRepo.UpdateVmName(vm)

	srv := vendors.NewHuaweiCloudService()
	ecsClient, err := srv.CreateEcsClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	imgClient, err := srv.CreateImgClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	vpcClient, err := srv.CreateVpcClient(host.CloudKey, host.CloudSecret, host.CloudRegion)

	if err != nil {
		result.Fail(err.Error())
		s.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, "", "", "")
		return
	}

	huaweiCloudService := vendors.NewHuaweiCloudService()
	vm.CouldInstId, _, err = huaweiCloudService.CreateInst(vm.Name, backing.Name, ecsClient, imgClient, vpcClient)
	if err != nil {
		result.Fail(err.Error())
		s.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, "", "", "")
		return
	}

	_, result.Msg, vm.NodeIp, vm.MacAddress, err = huaweiCloudService.QueryVm(vm.CouldInstId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, "", "", "")
		return
	}

	s.VmRepo.UpdateVmCloudInst(vm)

	url, _ := huaweiCloudService.QueryVnc(vm.CouldInstId, ecsClient)
	s.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, url, "", "")

	return
}

func (s HuaweiCloudService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	srv := vendors.NewHuaweiCloudService()
	ecsClient, err := srv.CreateEcsClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	if err != nil {
		status = consts.VmFailDestroy
	} else {
		err = srv.RemoveInst(vm.CouldInstId, ecsClient)
		if err != nil {
			status = consts.VmFailDestroy
		}
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CouldInstId}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
