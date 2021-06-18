package agentService

import "C"
import (
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/libs/shell"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/libs/string"
	"github.com/libvirt/libvirt-go-xml"
	"path/filepath"
	"strings"
)

type QemuService struct {
}

func NewQemuService() *QemuService {
	return &QemuService{}
}

func (s *QemuService) GenVmDef(src, vmName, rawPath, basePath string, vmMemory uint) (
	xml, macAddress string, err error) {

	domCfg := &libvirtxml.Domain{}
	err = domCfg.Unmarshal(src)
	if err != nil {
		return
	}

	mainDiskIndex := s.getMainDiskIndex(domCfg)

	domCfg.Name = vmName
	domCfg.UUID = _stringUtils.NewUuidWithSep()
	domCfg.Devices.Disks[mainDiskIndex].Source.File = &libvirtxml.DomainDiskSourceFile{
		File: rawPath,
	}
	domCfg.Devices.Disks[mainDiskIndex].BackingStore = &libvirtxml.DomainDiskBackingStore{
		Index: 0,
		Format: &libvirtxml.DomainDiskFormat{
			Type: "qcow2",
		},
		Source: &libvirtxml.DomainDiskSource{
			File: &libvirtxml.DomainDiskSourceFile{
				File: basePath,
			},
		},
	}

	//<graphics type="vnc" port="-1" autoport="yes" listen="0.0.0.0" passwd="P2ssw0rd">
	//<listen type="address" address="0.0.0.0"/>
	//</graphics>
	domCfg.Devices.Graphics = []libvirtxml.DomainGraphic{
		{
			VNC: &libvirtxml.DomainGraphicVNC{
				AutoPort: "yes",
				Port:     -1,
				Listen:   "0.0.0.0",
				Passwd:   "P2ssw0rd",
				Listeners: []libvirtxml.DomainGraphicListener{
					{
						Address: &libvirtxml.DomainGraphicListenerAddress{
							Address: "0.0.0.0",
						},
					},
				},
			},
		},
	}

	domCfg.Devices.Controllers = s.removeUnnecessaryPciCtrl(domCfg)

	if vmMemory != 0 {
		domCfg.Memory = &libvirtxml.DomainMemory{
			Unit:     "M",
			Value:    vmMemory,
			DumpCore: "yes",
		}
		domCfg.CurrentMemory = &libvirtxml.DomainCurrentMemory{
			Unit:  "M",
			Value: vmMemory,
		}
	}

	macAddress = s.GenMacAddress()
	domCfg.Devices.Interfaces[0].MAC.Address = macAddress

	//domCfg.OS.Type.Machine = s.GenMachine()

	xml, _ = domCfg.Marshal()

	return
}

func (s *QemuService) GenMacAddress() string {
	cmd := "dd if=/dev/urandom count=1 2>/dev/null | md5sum | sed 's/^\\(..\\)\\(..\\)\\(..\\)\\(..\\).*$/\\1:\\2:\\3:\\4/'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	mac := "fa:92:" + _stringUtils.TrimAll(output)
	return mac
}

func (s *QemuService) GenMachine() string {
	cmd := "qemu-system-x86_64 -M ? | grep default | awk '{print $1}'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	machine := _stringUtils.TrimAll(output)
	return machine
}

func (s *QemuService) GetBaseImagePath(vm commDomain.Vm) (path string) {
	dir := filepath.Join(agentConf.Inst.DirBase, vm.OsCategory.ToString(), vm.OsType.ToString())
	name := fmt.Sprintf("%s-%s", vm.OsVersion, vm.OsLang.ToString())

	path = filepath.Join(dir, name)

	return
}

func (s *QemuService) createDiskFile(basePath, vmName string, diskSize int) {
	vmRawPath := filepath.Join(agentConf.Inst.DirImage, vmName+".qcow2")

	var cmd string
	if basePath == "" {
		cmd = fmt.Sprintf("qemu-img create -f qcow2 %s %dG",
			vmRawPath, diskSize/1000)
	} else {
		cmd = fmt.Sprintf("qemu-img create -f qcow2 -o cluster_size=2M,backing_file=%s %s %dG",
			basePath, vmRawPath, diskSize/1000)
	}
	_, err := _shellUtils.ExeShellInDir(cmd, agentConf.Inst.DirKvm)
	if err != nil {
		_logUtils.Errorf("fail to generate vm, cmd %s, err %s.", cmd, err.Error())
		return
	}
}

func (s *QemuService) setVmProps(vm *commDomain.Vm) {
	osCategory := commConst.Windows
	osType := commConst.Win10
	osVersion := "x64-pro"
	osLang := commConst.ZH_CN

	vm.Base = fmt.Sprintf("%s/%s/%s-%s", osCategory.ToString(), osType.ToString(),
		osVersion, osLang.ToString())
}

func (s *QemuService) getMainDiskIndex(domCfg *libvirtxml.Domain) (ret int) {
	for index, item := range domCfg.Devices.Disks {
		if item.Device == "disk" && strings.Index(item.Source.File.File, "share") < 0 {
			ret = index
			return
		}
	}

	return
}

func (s *QemuService) getFirstPciCtrlIndex(domCfg *libvirtxml.Domain) (ret int) {
	for index, item := range domCfg.Devices.Controllers {
		if item.Type == "pci" && *item.Index == 0 {
			ret = index
			return
		}
	}

	return
}

func (s *QemuService) removeUnnecessaryPciCtrl(domCfg *libvirtxml.Domain) (ret []libvirtxml.DomainController) {
	for _, item := range domCfg.Devices.Controllers {
		if item.Type == "pci" && *item.Index != 0 {
			continue
		}
		//item.Model = "pci-root"
		ret = append(ret, item)
	}

	return
}
