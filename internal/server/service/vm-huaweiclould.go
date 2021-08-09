package serverService

import (
	"fmt"
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
		return
	}

	huaweiCloudService := vendors.NewHuaweiCloudService()
	id, _, err := huaweiCloudService.CreateInst(vm.Name, backing.Name, ecsClient, imgClient, vpcClient)
	vm.CouldInstId = id
	s.VmRepo.UpdateVmCloudInstId(vm)

	url, _ := huaweiCloudService.QueryVnc(id, ecsClient)
	s.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, url, "", "")

	return
}

func (s HuaweiCloudService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)

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

func (s HuaweiCloudService) genVmName(backing model.VmBacking, vmId uint) (name string) {
	name = fmt.Sprintf("test-%s-%s-%s-%d", backing.OsType, backing.OsVersion, backing.OsLang, vmId)

	return
}
