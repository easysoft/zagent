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

	url := fmt.Sprintf(testconst.ALIYUN_ECS_URL, host.CloudRegion)
	ecsClient, err := s.AliyunEcsService.CreateEcsClient(url, host.CloudKey, host.CloudSecret)
	vpcClient, err := s.AliyunEcsService.CreateVpcClient(url, host.CloudKey, host.CloudSecret)

	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "CreateEcsClient fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	switchId, _, err := s.AliyunEcsService.GetSwitch(host.VpcId, host.CloudRegion, vpcClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "GetSwitch fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	securityGroupId, err := s.AliyunEcsService.QuerySecurityGroupByVpc(host.VpcId, host.CloudRegion, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "QuerySecurityGroupByVpc fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	vm.CouldInstId, _, err = s.AliyunEcsService.CreateInst(vm.Name, backing.Name, switchId, securityGroupId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "CreateInst fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	err = s.AliyunEcsService.StartInst(vm.CouldInstId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "StartInst fail %s"+err.Error(), queueId, vm.ID, "", "", "")
		return
	}

	for i := 0; i < 2*60; i++ {
		<-time.After(1 * time.Second)

		status := ""
		status, vm.MacAddress, err = s.AliyunEcsService.QueryInst(vm.CouldInstId, host.CloudRegion, ecsClient)
		if err != nil {
			result.Fail(err.Error())
			s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "QueryInst fail %s"+err.Error(), queueId, vm.ID, vm.VncAddress, "", "")
			return
		}

		if status == "Running" {
			break
		}
	}

	vm.NodeIp, err = s.AliyunEcsService.AllocateIp(vm.CouldInstId, ecsClient)
	if err != nil {
		result.Fail(err.Error())
		s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), "AllocateIp fail %s"+err.Error(), queueId, vm.ID, vm.VncAddress, "", "")
		return
	}

	result.Pass("")
	s.VmRepo.UpdateVmCloudInst(vm)

	vncPassword, _ := s.AliyunEcsService.QueryVncPassword(vm.CouldInstId, testconst.ALIYUN_REGION, ecsClient)
	vm.VncAddress, _ = s.AliyunEcsService.QueryVncUrl(
		vm.CouldInstId, vncPassword, host.CloudRegion, vm.OsCategory == consts.Windows, ecsClient)

	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")

	return
}

func (s AliyunVmService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	url := fmt.Sprintf(testconst.ALIYUN_ECS_URL, host.CloudRegion)
	ecsClient, err := s.AliyunEcsService.CreateEcsClient(url, host.CloudKey, host.CloudSecret)
	if err != nil {
		status = consts.VmFailDestroy
	} else {
		err = s.AliyunEcsService.RemoveInst(vm.CouldInstId, ecsClient)
		if err != nil {
			status = consts.VmFailDestroy
		}
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CouldInstId}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
