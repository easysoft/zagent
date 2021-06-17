package agentService

import "C"
import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/libs/shell"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/libs/string"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"path/filepath"
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

func (s *LibvirtService) GetVm(name string) (dom *libvirt.Domain) {
	s.Connect(ConnStrLocal)
	defer func() {
		if res, _ := Conn.Close(); res != 0 {
			_logUtils.Errorf("close() == %d, expected 0", res)
		}
	}()
	names, err := Conn.ListDefinedDomains()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	if len(names) == 0 {
		_logUtils.Errorf("length of domains shouldn't be 0.")
		return
	}

	dom, err = Conn.LookupDomainByName(name)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	return
}

func (s *LibvirtService) CloneVm(src, vmName string) (dom *libvirt.Domain, macAddress string, err error) {
	templ := s.GetDomainDef(src)
	xml, macAddress, _ := s.GenVmDef(templ, vmName, 0, "", "")

	dom, err = Conn.DomainCreateXML(xml, 0)
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

func (s *LibvirtService) GenVmDef(templ, vmName string, vmMemory uint, vmCdrom string, vmCdrom2 string) (
	xml, macAddress string, err error) {

	domCfg := &libvirtxml.Domain{}
	err = domCfg.Unmarshal(templ)
	if err != nil {
		return
	}

	domCfg.Name = vmName
	rawPath := filepath.Join(agentConf.Inst.DirImage, vmName+".qcow2")
	domCfg.Devices.Disks[0].Source.File = &libvirtxml.DomainDiskSourceFile{
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

	if vmCdrom != "" {
		cdromPath := filepath.Join(agentConf.Inst.DirIso, vmCdrom)
		domCfg.Devices.Disks[2].Source.File = &libvirtxml.DomainDiskSourceFile{
			File: cdromPath,
		}
	}
	if vmCdrom2 != "" {
		cdromPath2 := filepath.Join(agentConf.Inst.DirIso, vmCdrom2)
		domCfg.Devices.Disks[3].Source.File = &libvirtxml.DomainDiskSourceFile{
			File: cdromPath2,
		}
	}

	macAddress = s.GenMacAddress()
	domCfg.Devices.Interfaces[0].MAC.Address = macAddress

	machine := s.GenMachine()
	domCfg.OS.Type.Machine = machine

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
	dom.Create()
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
