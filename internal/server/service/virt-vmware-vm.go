package serverService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/api"
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/srv"
	vmwareService "github.com/easysoft/zagent/internal/server/service/vendors/vmware"
)

type VmWareCloudVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	CommonService   *CommonService   `inject:""`
	VmCommonService *VmCommonService `inject:""`
	HistoryService  *HistoryService  `inject:""`
}

func (s VmWareCloudVmService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RpcResp) {
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

	client := vmwareService.NewVMWareService() // 8697
	err := client.Connect(fmt.Sprintf("https://%s:%d", host.Ip, host.Port), host.CloudIamUser, host.CloudIamPassword)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	// create machine
	vmInst, err := client.CreateVm(backing.Name, vm.Name)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	vm.CloudInstId = vmInst.IdVM
	vm.MacAddress, _ = client.GetVmNic(vm.CloudInstId)

	// save to db
	result.Pass("")
	s.VmRepo.UpdateVmCloudInst(vm)

	//vm.VncAddress, _ = huaweiCloudService.QueryVnc(vm.CloudInstId, ecsClient)
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")

	return
}

func (s VmWareCloudVmService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	virtualBox, err := s.CreateClient(host.Ip, host.Port, host.CloudIamUser, host.CloudIamPassword)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	machine, err := virtualBox.FindMachine(vm.Name)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	machineState, err := machine.GetMachineState()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}
	_logUtils.Infof("machine state %s", *machineState)

	session, err := virtualBox.GetSession()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	err = machine.Lock(session, virtualboxsrv.LockTypeShared)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	console, err := session.GetConsole()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	progress, err := console.PowerDown()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}
	err = progress.WaitForCompletion(10000)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	media, err := machine.Unregister()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	err = machine.DiscardSettings()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	err = machine.DeleteConfig(media)
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	err = session.Release()
	if err != nil {
		s.CommonService.ReturnErr(&result, err, queueId, vm.ID)
		return
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CloudInstId}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}

func (s VmWareCloudVmService) CreateClient(ip string, port int, account, password string) (
	client *virtualboxapi.VirtualBox, err error) {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client = virtualboxapi.NewVirtualBox(account, password, url, false, "")

	err = client.Logon()

	return
}
