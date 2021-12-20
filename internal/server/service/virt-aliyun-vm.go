package serverService

import (
	"fmt"
	"github.com/easysoft/zv/internal/comm/const"
	_domain "github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	"github.com/easysoft/zv/internal/server/service/vendors/aliyun"
	serverConst "github.com/easysoft/zv/internal/server/utils/const"
	"time"
)

type AliyunVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService   *VmCommonService          `inject:""`
	HistoryService    *HistoryService           `inject:""`
	AliyunEcsService  *aliyun.AliyunEcsService  `inject:""`
	AliyunCommService *aliyun.AliyunCommService `inject:""`
}

func NewAliyunVmService() *AliyunVmService {
	return &AliyunVmService{}
}

func (s AliyunVmService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RpcResp) {
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

	url := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	ecsClient, err := s.AliyunCommService.CreateEcsClient(url, host.CloudKey, host.CloudSecret)
	vpcClient, err := s.AliyunCommService.CreateVpcClient(url, host.CloudKey, host.CloudSecret)

	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "CreateEcsClient fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	switchId, _, err := s.AliyunCommService.GetSwitch(host.VpcId, host.CloudRegion, vpcClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "GetSwitch fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	securityGroupId, err := s.AliyunCommService.QuerySecurityGroupByVpc(host.VpcId, host.CloudRegion, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "QuerySecurityGroupByVpc fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	vm.CloudInstId, _, err = s.AliyunEcsService.CreateInst(vm.Name, backing.Name, switchId, securityGroupId,
		host.CloudRegion, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "CreateInst fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	err = s.AliyunEcsService.StartInst(vm.CloudInstId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "StartInst fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	for i := 0; i < 2*60; i++ {
		<-time.After(1 * time.Second)

		status := ""
		status, vm.MacAddress, err = s.AliyunEcsService.QueryInst(vm.CloudInstId, host.CloudRegion, ecsClient)
		if err != nil {
			result.Fail(err.Error())
			s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "QueryInst fail %s"+err.Error(), queueId, vm.ID, vm.VncAddress, "", "")
			return
		}

		if status == "Running" {
			break
		}
	}

	vm.NodeIp, err = s.AliyunEcsService.AllocateIp(vm.CloudInstId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "AllocateIp fail %s"+err.Error(), queueId, vm.ID, vm.VncAddress, "", "")
		return
	}

	result.Pass("")
	s.VmRepo.UpdateVmCloudInst(vm)

	vncPassword, _ := s.AliyunEcsService.QueryVncPassword(vm.CloudInstId, host.CloudRegion, ecsClient)
	vm.VncAddress, _ = s.AliyunEcsService.QueryVncUrl(
		vm.CloudInstId, vncPassword, host.CloudRegion, vm.OsCategory == consts.Windows, ecsClient)

	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")

	return
}

func (s AliyunVmService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	url := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	ecsClient, err := s.AliyunCommService.CreateEcsClient(url, host.CloudKey, host.CloudSecret)
	if err != nil {
		status = consts.VmDestroyFail
	} else {
		err = s.AliyunEcsService.RemoveInst(vm.CloudInstId, ecsClient)
		if err != nil {
			status = consts.VmDestroyFail
		}
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CloudInstId}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
