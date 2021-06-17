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
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ConnStrLocal  = "qemu:///system"
	ConnStrRemote = "qemu+ssh://192.168.0.56:22/system?keyfile=~/.ssh/id_rsa"
)

var (
	Conn *libvirt.Connect
)

type LibvirtService struct {
}

func NewLibvirtService() *LibvirtService {
	return &LibvirtService{}
}

func (s *LibvirtService) CreateVm(vm *commDomain.Vm) (dom *libvirt.Domain, macAddress string, err error) {
	s.setVmProps(vm)

	srcXml := s.GetDomainDef(vm.Src)

	vmXml := ""
	baseDiskPath := ""
	rawPath := filepath.Join(agentConf.Inst.DirImage, vm.Name+".qcow2")
	vmXml, vm.MacAddress, baseDiskPath, _ = s.GenVmDef(srcXml, vm.Name, rawPath, 0)

	vm.DiskSize, err = s.getDiskSize(baseDiskPath)
	if err != nil || vm.DiskSize == 0 {
		_logUtils.Errorf("wrong vm disk size %d, err %s", vm.DiskSize, err.Error())
		return
	}

	s.createDiskFile(vm.Base, vm.Name, vm.DiskSize)

	dom, err = Conn.DomainCreateXML(vmXml, 0)
	return
}

func (s *LibvirtService) GetVm(name string) (dom *libvirt.Domain) {
	s.Connect(ConnStrLocal)
	defer func() {
		if res, _ := Conn.Close(); res != 0 {
			_logUtils.Errorf("close() == %d, expected 0", res)
		}
	}()

	dom, err := Conn.LookupDomainByName(name)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	return
}

func (s *LibvirtService) StartVm(dom *libvirt.Domain) (err error) {
	err = dom.Create()
	return
}
func (s *LibvirtService) DestroyVm(dom *libvirt.Domain) (err error) {
	err = dom.Destroy()
	return
}
func (s *LibvirtService) UndefineVm(dom *libvirt.Domain) (err error) {
	err = dom.Undefine()
	return
}

func (s *LibvirtService) GenVmDef(src, vmName, rawPath string, vmMemory uint) (
	xml, macAddress, diskPath string, err error) {

	domCfg := &libvirtxml.Domain{}
	err = domCfg.Unmarshal(src)
	if err != nil {
		return
	}

	mainDiskIndex := s.getMainDiskIndex(domCfg)
	diskPath = domCfg.Devices.Disks[mainDiskIndex].Source.File.File

	domCfg.Name = vmName
	domCfg.Devices.Disks[mainDiskIndex].Source.File = &libvirtxml.DomainDiskSourceFile{
		File: rawPath,
	}

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

	machine := s.GenMachine()
	domCfg.OS.Type.Machine = machine

	xml, _ = domCfg.Marshal()

	return
}

func (s *LibvirtService) GenMacAddress() string {
	cmd := "dd if=/dev/urandom count=1 2>/dev/null | md5sum | sed 's/^\\(..\\)\\(..\\)\\(..\\)\\(..\\).*$/\\1:\\2:\\3:\\4/'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	mac := "fa:92:" + _stringUtils.TrimAll(output)
	return mac
}

func (s *LibvirtService) GenMachine() string {
	cmd := "qemu-system-x86_64 -M ? | grep default | awk '{print $1}'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	machine := _stringUtils.TrimAll(output)
	return machine
}

func (s *LibvirtService) GetDomainDef(name string) (xml string) {
	dom := s.GetVm(name)
	if dom == nil {
		return
	}

	xml, _ = dom.GetXMLDesc(0)

	return
}

func (s *LibvirtService) Connect(str string) {
	if Conn != nil {
		active, err := Conn.IsAlive()
		if err != nil {
			_logUtils.Errorf(err.Error())
		}
		if active {
			return
		}
	}

	var err error
	Conn, err = libvirt.NewConnect(str)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	active, err := Conn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	if !active {
		_logUtils.Errorf("not active")
	}
	return
}

func (s *LibvirtService) GetBaseImagePath(vm commDomain.Vm) (path string) {
	dir := filepath.Join(agentConf.Inst.DirBase, vm.OsCategory.ToString(), vm.OsType.ToString())
	name := fmt.Sprintf("%s-%s", vm.OsVersion, vm.OsLang.ToString())

	path = filepath.Join(dir, name)

	return
}

func (s *LibvirtService) createDiskFile(base, vmName string, diskSize int) {
	basePath := ""
	if base != "" {
		basePath = filepath.Join(agentConf.Inst.DirBase, base)
	}
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

func (s *LibvirtService) getDiskSize(path string) (size int, err error) {
	cmd := fmt.Sprintf("qemu-img info %s | grep 'virtual size' | grep -v 'grep' | awk '{print $3}", path)

	output, err := _shellUtils.ExeShellInDir(cmd, agentConf.Inst.DirKvm)
	if err != nil {
		return
	}
	size, err = strconv.Atoi(strings.TrimSpace(output))

	return
}

func (s *LibvirtService) setVmProps(vm *commDomain.Vm) {
	osCategory := commConst.Windows
	osType := commConst.Win10
	osVersion := "x64-pro"
	osLang := commConst.ZH_CN

	vm.Base = fmt.Sprintf("%s/%s/%s-%s", osCategory.ToString(), osType.ToString(),
		osVersion, osLang.ToString())
}

func (s *LibvirtService) getMainDiskIndex(domCfg *libvirtxml.Domain) (ret int) {
	for index, item := range domCfg.Devices.Disks {
		if item.Device == "disk" && *item.Address.Drive.Unit == 0 {
			ret = index
			return
		}
	}
}
