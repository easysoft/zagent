package agentService

import "C"
import (
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
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

func (s *LibvirtService) CreateVm(req *commDomain.KvmReq) (dom *libvirt.Domain, vncPort int, err error) {
	vmMacAddress := req.VmMacAddress
	vmUniqueName := req.VmUniqueName
	vmBackingPath := req.VmBackingPath
	vmTemplateName := req.VmTemplateName

	return
}

func (s *LibvirtService) CreateVmTest(vm *commDomain.Vm) (
	dom *libvirt.Domain, macAddress string, vncPort int, err error) {
	s.setVmProps(vm)

	srcXml := s.GetVmDef(vm.Tmpl)

	backingPath := ""
	if vm.Backing != "" {
		backingPath = filepath.Join(agentConf.Inst.DirBase, vm.Backing)
	}
	backingPath += ".qcow2"

	vmXml := ""
	rawPath := filepath.Join(agentConf.Inst.DirImage, vm.Name+".qcow2")
	vmXml, vm.MacAddress, _ = s.QemuService.GenVmDef(srcXml, vm.Name, rawPath, backingPath, 0)

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

func (s *LibvirtService) GetVm(name string) (dom *libvirt.Domain) {
	defer func() {
		if res, _ := LibvirtConn.Close(); res != 0 {
			_logUtils.Errorf("close() == %d, expected 0", res)
		}
	}()

	dom, err := LibvirtConn.LookupDomainByName(name)
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
func (s *LibvirtService) DestroyVmByName(name string) (err error) {
	dom := s.GetVm(name)
	err = dom.Destroy()
	return
}
func (s *LibvirtService) UndefineVm(dom *libvirt.Domain) (err error) {
	err = dom.Undefine()
	return
}

func (s *LibvirtService) GetVmDef(name string) (xml string) {
	dom := s.GetVm(name)
	if dom == nil {
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

func (s *LibvirtService) setVmProps(vm *commDomain.Vm) {
	osCategory := commConst.Windows
	osType := commConst.Win10
	osVersion := "x64-pro"
	osLang := commConst.ZH_CN

	vm.Backing = fmt.Sprintf("%s/%s/%s-%s", osCategory.ToString(), osType.ToString(),
		osVersion, osLang.ToString())

	vm.Tmpl = fmt.Sprintf("tmpl-%s-%s-%s",
		osType.ToString(), osVersion, osLang.ToString())
	vm.Name = fmt.Sprintf("test-%s-%s-%s-%s",
		osType.ToString(), osVersion, osLang.ToString(), _stringUtils.NewUuid())
}
