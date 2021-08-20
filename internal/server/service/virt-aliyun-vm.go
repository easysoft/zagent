package serverService

import (
	"fmt"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"time"
)

type AliyunVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService  *VmCommonService          `inject:""`
	HistoryService   *HistoryService           `inject:""`
	AliyunEcsService *vendors.AliyunEcsService `inject:""`
}

func (s AliyunVmService) CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
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

	url := fmt.Sprintf(testconst.ALIYUN_URL, host.CloudRegion)
	ecsClient, err := s.AliyunEcsService.CreateClient(url, host.CloudKey, host.CloudSecret)

	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, "", "", "")
		return
	}

	vm.CouldInstId, _, err = s.AliyunEcsService.CreateInst(vm.Name, backing.Name, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, "", "", "")
		return
	}

	result.Pass("")

	for i := 0; i < 60; i++ {
		<-time.After(1 * time.Second)

		vm.NodeIp, vm.MacAddress, err = s.AliyunEcsService.QueryInst(vm.CouldInstId, ecsClient)
		if err != nil {
			result.Fail(err.Error())
			s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")
			return
		}

		if vm.MacAddress != "" {
			break
		}
	}

	s.VmRepo.UpdateVmCloudInst(vm)

	vm.VncAddress, _ = s.AliyunEcsService.QueryVnc(vm.CouldInstId, ecsClient)
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")

	return
}

func (s AliyunVmService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {

	return
}
