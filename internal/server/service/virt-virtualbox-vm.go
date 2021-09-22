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
)

type VirtualboxCloudVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	CommonService   *CommonService   `inject:""`
	VmCommonService *VmCommonService `inject:""`
	HistoryService  *HistoryService  `inject:""`
}

func (s VirtualboxCloudVmService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RpcResp) {
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

	var err error
	var client *virtualboxapi.VirtualBox
	var osTypeId string
	var newMachineId string
	var tmpl *virtualboxapi.Machine
	var machine *virtualboxapi.Machine
	var newMachine *virtualboxapi.Machine
	var snapshot *virtualboxapi.Machine
	var snapshotMachine *virtualboxapi.Machine
	var adpt *virtualboxapi.NetworkAdapter
	var session *virtualboxapi.Session
	var progress *virtualboxapi.Progress

	client, err = s.CreateClient(host.Ip, host.Port, host.CloudIamUser, host.CloudIamPassword)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	// get backing tmpl
	tmpl, err = client.FindMachine(backing.Name)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	osTypeId, err = tmpl.GetOsTypeId()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	snapshot, err = tmpl.FindSnapshot()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	snapshotMachine, err = snapshot.FindSnapshotMachine()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	// create machine
	newMachineId, err = client.CreateMachine(vm.Name, osTypeId)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	progress, newMachine, err = snapshotMachine.CloneTo(newMachineId)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = newMachine.SetCPUCount(uint32(backing.SuggestCpuCount))
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	err = newMachine.SetMemorySize(uint32(backing.SuggestMemorySize))
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	adpt, err = newMachine.GetNetworkAdapter(0)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	err = adpt.SetBridge(host.Bridge)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	vm.MacAddress, err = adpt.GetMACAddress()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	_logUtils.Infof("machine mac address %s", vm.MacAddress)

	err = newMachine.SaveSettings()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	err = newMachine.Register()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	// launch machine
	machine, err = client.FindMachine(vm.Name)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	session, err = client.GetSession()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	progress, err = machine.Launch(session.ManagedObjectId)
	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

final:
	result.Pass("")
	s.VmRepo.UpdateVmCloudInst(vm)

	//vm.VncAddress, _ = huaweiCloudService.QueryVnc(vm.CloudInstId, ecsClient)
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncAddress, "", "")

	return
}

func (s VirtualboxCloudVmService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	var err error
	var virtualBox *virtualboxapi.VirtualBox
	var machine *virtualboxapi.Machine
	var machineState *virtualboxsrv.MachineState
	var session *virtualboxapi.Session
	var console *virtualboxapi.Console
	var progress *virtualboxapi.Progress
	var media []string

	virtualBox, err = s.CreateClient(host.Ip, host.Port, host.CloudIamUser, host.CloudIamPassword)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	machine, err = virtualBox.FindMachine(vm.Name)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	machineState, err = machine.GetMachineState()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}
	_logUtils.Infof("machine state %s", *machineState)

	session, err = virtualBox.GetSession()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = machine.Lock(session, virtualboxsrv.LockTypeShared)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	console, err = session.GetConsole()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	progress, err = console.PowerDown()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	media, err = machine.Unregister()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = machine.DiscardSettings()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = machine.DeleteConfig(media)
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

	err = session.Release()
	if err != nil {
		result.Fail(err.Error())
		goto final
	}

final:
	s.VmCommonService.SaveVmDestroyResult(result.IsSuccess(), result.Msg, queueId, vmId)

	return
}

func (s VirtualboxCloudVmService) CreateClient(ip string, port int, account, password string) (
	client *virtualboxapi.VirtualBox, err error) {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client = virtualboxapi.NewVirtualBox(account, password, url, false, "")

	err = client.Logon()

	return
}
