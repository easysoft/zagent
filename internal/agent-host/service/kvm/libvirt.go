package hostAgentKvmService

import "C"
import (
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
	"path/filepath"
)

const (
	LibvirtConnStrLocal  = "qemu:///system"
	LibvirtConnStrRemote = "qemu+ssh://%s:22/system?socket=/var/run/libvirt/libvirt-sock"
)

var (
	LibvirtConn *libvirt.Connect
)

type LibvirtService struct {
	QemuService *QemuService `inject:""`
}

func NewLibvirtService() *LibvirtService {
	connStr := ""
	if agentConf.Inst.Host == "" {
		connStr = LibvirtConnStrLocal
	} else {
		connStr = fmt.Sprintf(LibvirtConnStrRemote, agentConf.Inst.Host)
	}

	s := &LibvirtService{}
	s.Connect(connStr)

	return s
}

func (s *LibvirtService) CreateVm(req *domain.KvmReq, removeSameName bool) (dom *libvirt.Domain,
	vmVncPort int, vmRawPath, vmBackingPath string, err error) {
	vmMacAddress := req.VmMacAddress
	vmUniqueName := req.VmUniqueName
	vmBackingPath = filepath.Join(agentConf.Inst.DirKvm, req.VmBackingPath)
	vmTemplateName := req.VmTemplateName
	vmMemorySize := req.VmMemorySize
	vmDiskSize := req.VmDiskSize

	if removeSameName {
		s.DestroyVmByName(vmUniqueName, true)
	}

	tmplXml := s.GetVmDef(vmTemplateName)
	vmXml := ""
	vmXml, vmRawPath, _ = s.QemuService.GenVmDef(tmplXml, vmMacAddress, vmUniqueName, vmBackingPath, vmMemorySize)

	if err != nil {
		_logUtils.Errorf("err gen vm xml, err %s", err.Error())
		return
	}

	s.QemuService.createDiskFile(vmBackingPath, vmUniqueName, vmDiskSize)

	dom, err = LibvirtConn.DomainCreateXML(vmXml, 0)

	if err == nil {
		newXml := ""
		newXml, err = dom.GetXMLDesc(0)
		if err != nil {
			return
		}

		newDomCfg := &libvirtxml.Domain{}
		err = newDomCfg.Unmarshal(newXml)

		vmMacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
		vmVncPort = newDomCfg.Devices.Graphics[0].VNC.Port
	}

	return
}

func (s *LibvirtService) CreateVmTest(vm *domain.Vm) (
	dom *libvirt.Domain, macAddress string, vncPort int, err error) {
	s.setVmProps(vm)

	srcXml := s.GetVmDef(vm.Tmpl)

	backingPath := ""
	if vm.Backing != "" {
		backingPath = filepath.Join(agentConf.Inst.DirBaking, vm.Backing)
	}
	backingPath += ".qcow2"

	vmXml := ""
	rawPath := filepath.Join(agentConf.Inst.DirImage, vm.Name+".qcow2")
	vmXml, vm.MacAddress, _ = s.QemuService.GenVmDefTest(srcXml, vm.Name, rawPath, backingPath, 0)

	if err != nil || vm.DiskSize == 0 {
		_logUtils.Errorf("wrong vm disk size %d, err %s", vm.DiskSize, err.Error())
		return
	}

	s.QemuService.createDiskFile(backingPath, vm.Name, vm.DiskSize)

	dom, err = LibvirtConn.DomainCreateXML(vmXml, 0)

	if err == nil {
		newXml := ""
		newXml, err = dom.GetXMLDesc(0)
		if err != nil {
			return
		}

		newDomCfg := &libvirtxml.Domain{}
		err = newDomCfg.Unmarshal(newXml)

		macAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
		vncPort = newDomCfg.Devices.Graphics[0].VNC.Port
	}

	return
}

func (s *LibvirtService) ListVm() (doms []libvirt.Domain) {
	doms, err := LibvirtConn.ListAllDomains(0)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	return
}

func (s *LibvirtService) GetVm(name string) (dom *libvirt.Domain, err error) {
	dom, err = LibvirtConn.LookupDomainByName(name)
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
func (s *LibvirtService) DestroyVmByName(name string, removeDiskImage bool) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Destroy()

	if removeDiskImage {
		s.QemuService.RemoveDisk(dom)
	}

	return
}
func (s *LibvirtService) UndefineVm(dom *libvirt.Domain) (err error) {
	err = dom.Undefine()

	return
}

func (s *LibvirtService) GetVmDef(name string) (xml string) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	xml, _ = dom.GetXMLDesc(0)

	return
}

func (s *LibvirtService) Connect(str string) {
	if LibvirtConn != nil {
		active, err := LibvirtConn.IsAlive()
		if err != nil {
			_logUtils.Errorf(err.Error())
		}
		if active {
			return
		}
	}

	var err error
	LibvirtConn, err = libvirt.NewConnect(str)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	active, err := LibvirtConn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	if !active {
		_logUtils.Errorf("not active")
	}
	return
}

func (s *LibvirtService) setVmProps(vm *domain.Vm) {
	osCategory := consts.Windows
	osType := consts.Win10
	osVersion := "x64-pro"
	osLang := consts.ZH_CN

	vm.Backing = fmt.Sprintf("%s/%s/%s-%s", osCategory.ToString(), osType.ToString(),
		osVersion, osLang.ToString())

	vm.Tmpl = fmt.Sprintf("tmpl-%s-%s-%s",
		osType.ToString(), osVersion, osLang.ToString())
	vm.Name = fmt.Sprintf("test-%s-%s-%s-%s",
		osType.ToString(), osVersion, osLang.ToString(), _stringUtils.NewUuid())
}
