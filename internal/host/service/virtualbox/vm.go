package virtualboxService

import "C"
import (
	"fmt"
	"github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/pkg/lib/log"
	"github.com/easysoft/zv/internal/server/service/vendors/virtualbox/api"
	"github.com/easysoft/zv/internal/server/service/vendors/virtualbox/srv"
)

const (
	ip   = "127.0.0.1"
	port = 18083
)

type VirtualBoxService struct {
}

func NewVirtualBoxService() *VirtualBoxService {
	return &VirtualBoxService{}
}

func (s VirtualBoxService) Create(req v1.VirtualBoxReq) (result _domain.RpcResp, err error) {
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

	client, err = s.CreateClient(ip, port, req.CloudIamUser, req.CloudIamPassword)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	// get backing tmpl
	tmpl, err = client.FindMachine(req.BackingName)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	osTypeId, err = tmpl.GetOsTypeId()
	if err != nil {
		result.Fail(err.Error())
		return
	}
	snapshot, err = tmpl.FindSnapshot()
	if err != nil {
		result.Fail(err.Error())
		return
	}
	snapshotMachine, err = snapshot.FindSnapshotMachine()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	// create machine
	newMachineId, err = client.CreateMachine(req.VmUniqueName, osTypeId)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	progress, newMachine, err = snapshotMachine.CloneTo(newMachineId)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = newMachine.SetCPUCount(uint32(req.VmCpu))
	if err != nil {
		result.Fail(err.Error())
		return
	}
	err = newMachine.SetMemorySize(uint32(req.VmMemorySize))
	if err != nil {
		result.Fail(err.Error())
		return
	}

	adpt, err = newMachine.GetNetworkAdapter(0)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	err = adpt.SetBridge(req.Bridge)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	macAddress, err := adpt.GetMACAddress()
	if err != nil {
		result.Fail(err.Error())
		return
	}
	_logUtils.Infof("machine mac address %s", macAddress)

	err = newMachine.SaveSettings()
	if err != nil {
		result.Fail(err.Error())
		return
	}
	err = newMachine.Register()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	// launch machine
	machine, err = client.FindMachine(req.VmUniqueName)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	session, err = client.GetSession()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	progress, err = machine.Launch(session.ManagedObjectId)
	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	result.Pass("")
	result.Payload = v1.VirtualBoxResp{
		Name:       req.VmUniqueName,
		MacAddress: macAddress,
	}

	return
}

func (s VirtualBoxService) Destroy(req v1.VirtualBoxReq) (result _domain.RpcResp) {
	var err error
	var virtualBox *virtualboxapi.VirtualBox
	var machine *virtualboxapi.Machine
	var machineState *virtualboxsrv.MachineState
	var session *virtualboxapi.Session
	var console *virtualboxapi.Console
	var progress *virtualboxapi.Progress
	var media []string

	virtualBox, err = s.CreateClient(ip, port, req.CloudIamUser, req.CloudIamPassword)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	machine, err = virtualBox.FindMachine(req.VmUniqueName)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	machineState, err = machine.GetMachineState()
	if err != nil {
		result.Fail(err.Error())
		return
	}
	_logUtils.Infof("machine state %s", *machineState)

	session, err = virtualBox.GetSession()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = machine.Lock(session, virtualboxsrv.LockTypeShared)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	console, err = session.GetConsole()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	progress, err = console.PowerDown()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = progress.WaitForCompletion(10000)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	media, err = machine.Unregister()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = machine.DiscardSettings()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = machine.DeleteConfig(media)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	err = session.Release()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	return
}

func (s VirtualBoxService) CreateClient(ip string, port int, account, password string) (
	client *virtualboxapi.VirtualBox, err error) {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client = virtualboxapi.NewVirtualBox(account, password, url, false, "")

	err = client.Logon()

	return
}
