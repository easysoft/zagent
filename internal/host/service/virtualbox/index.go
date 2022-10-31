package virtualboxService

import "C"
import (
	"errors"
	"fmt"
	"github.com/easysoft/zagent/cmd/host/router/v1"
	"github.com/easysoft/zagent/internal/pkg/vendors/virtualbox/api"
	_domain "github.com/easysoft/zagent/pkg/domain"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
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
	backingName := req.BackingName
	vncPort := _commonUtils.GetVncPort()
	vncPassword := _stringUtils.Uuid()

	if s.isVmExist(vmName) {
		err = errors.New("vm with same name exist")
		return
	}

	cmd := fmt.Sprintf("VBoxManage snapshot %s delete %s-snap", backingName, backingName)
	out, err := _shellUtils.ExeShell(cmd)

	cmd = fmt.Sprintf("VBoxManage snapshot %s take %s-snap --description 'for zv linked clones'",
		backingName, backingName)
	out, err = _shellUtils.ExeShell(cmd)

	cmd = fmt.Sprintf(
		"VBoxManage clonevm %s"+
			" --name=%s"+
			" --snapshot=%s-snap"+
			" --options=link"+
			" --register",
		backingName, vmName, backingName)
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

func (s VirtualBoxService) Destroy(req v1.VirtualBoxReq) (result _domain.RemoteResp, err error) {
	cmd := fmt.Sprintf("VBoxManage controlvm %s poweroff", req.VmUniqueName)
	_, err = _shellUtils.ExeShell(cmd)

	cmd = fmt.Sprintf("VBoxManage unregistervm --delete %s", req.VmUniqueName)
	_, err = _shellUtils.ExeShell(cmd)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	_commonUtils.RemoveVncPort(req.VncPort)

	return
}

func (s VirtualBoxService) ListTmpl(req v1.VirtualBoxReq) (result _domain.RemoteResp, err error) {
	out, _ := _shellUtils.ExeShell(fmt.Sprintf("VBoxManage list vms"))
	arr := strings.Split(out, "\n")

	list := make([]virtualboxapi.Machine, 0)
	regx1, _ := regexp.Compile(`"(.+)"`)
	regx2, _ := regexp.Compile(`\{(.+)\}`)

	for _, item := range arr {
		arr1 := regx1.FindStringSubmatch(item)
		name := arr1[1]

		arr2 := regx2.FindStringSubmatch(item)
		id := arr2[1]

		if req.Prefix == "" || strings.Index(name, req.Prefix) > -1 {
			item := virtualboxapi.Machine{
				ID:   id,
				Name: name,
			}
			list = append(list, item)
		}
	}

	result.Pass("")
	result.Payload = list

	return
}

func (s VirtualBoxService) isVmExist(vmName string) (ret bool) {
	out, _ := _shellUtils.ExeShell(fmt.Sprintf("VBoxManage list vms"))

	if strings.Index(out, fmt.Sprintf("\"%s\"", vmName)) > -1 {
		return true
	}

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
