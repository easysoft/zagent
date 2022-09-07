package virtualboxService

import "C"
import (
	"fmt"
	"github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/pkg/vendors/virtualbox/api"
	"github.com/easysoft/zv/internal/pkg/vendors/virtualbox/srv"
	_domain "github.com/easysoft/zv/pkg/domain"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"regexp"
	"strings"
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

func (s VirtualBoxService) Create(req v1.VirtualBoxReq) (result _domain.RemoteResp, err error) {
	vmName := req.VmUniqueName
	vncPort := _commonUtils.GetVncPort()
	vncPassword := _stringUtils.Uuid()

	cmd := fmt.Sprintf("VBoxManage snapshot %s delete %s-snap", vmName, vmName)
	out, err := _shellUtils.ExeShell(cmd)

	cmd = fmt.Sprintf("VBoxManage snapshot %s take %s-snap --description 'for zv linked clones'", vmName, vmName)
	out, err = _shellUtils.ExeShell(cmd)

	cmd = fmt.Sprintf(
		"VBoxManage clonevm %s"+
			" --name=\"%s\""+
			"--snapshot=%s-snap"+
			" --register"+
			" --mode=all"+
			" --options=link",
		req.BackingName, vmName, vmName)
	out, err = _shellUtils.ExeShell(cmd)

	bridge, _, _ := s.getBridgeAndMacAddress(req.BackingName)

	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s --nic1 bridged", req.VmUniqueName))
	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s  --bridgeadapter1 %s",
		req.VmUniqueName, bridge))

	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s --vrde on", req.VmUniqueName))
	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s --vrdeproperty VNCPassword=%s",
		req.VmUniqueName, vncPassword))
	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s --vrdemulticon on --vrdeport %d",
		req.VmUniqueName, vncPort))
	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage modifyvm %s --vram 128", req.VmUniqueName))

	out, err = _shellUtils.ExeShell(fmt.Sprintf("VBoxManage startvm %s", req.VmUniqueName))

	_logUtils.Infof(out)

	_, macAddress, _ := s.getBridgeAndMacAddress(req.BackingName)

	result.Pass("")
	result.Payload = v1.VirtualBoxResp{
		Name:       req.VmUniqueName,
		MacAddress: macAddress,
	}

	return
}

func (s VirtualBoxService) ListTmpl(req v1.VirtualBoxReq) (result _domain.RemoteResp, err error) {
	virtualBox, err := s.CreateClient(ip, port, req.CloudIamUser, req.CloudIamPassword)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	machines, err := virtualBox.GetMachines()
	if err != nil {
		result.Fail(err.Error())
		return
	}

	list := make([]virtualboxapi.Machine, 0)
	for _, item := range machines {
		id, err1 := item.GetID()
		name, err2 := item.GetName()
		if err1 == nil && err2 == nil && (req.Prefix == "" || strings.Index(name, req.Prefix) > -1) {
			item.ID = id
			item.Name = name

			list = append(list, *item)
		}
	}

	result.Pass("")
	result.Payload = list

	return
}

func (s VirtualBoxService) Destroy(req v1.VirtualBoxReq) (result _domain.RemoteResp, err error) {
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

	_commonUtils.RemoveVncPort(req.VncPort)

	return
}

func (s VirtualBoxService) CreateClient(ip string, port int, account, password string) (
	client *virtualboxapi.VirtualBox, err error) {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client = virtualboxapi.NewVirtualBox(account, password, url, false, "")

	err = client.Logon()

	return
}

func (s VirtualBoxService) getBridgeAndMacAddress(vmName string) (bridge, macAddress string, err error) {
	out, err := _shellUtils.ExeShell(fmt.Sprintf("VBoxManage showvminfo %s", vmName))

	regx1, _ := regexp.Compile(`MAC: ([A-Z0-9]+),`)
	arr1 := regx1.FindStringSubmatch(out)
	macAddress = arr1[1]

	// MAC: 080027998EA2, Attachment: Bridged Interface 'br0',
	regx2, _ := regexp.Compile(`Bridged Interface '([a-zA-Z0-9]+)',`)
	arr2 := regx2.FindStringSubmatch(out)
	bridge = arr2[1]

	return
}
