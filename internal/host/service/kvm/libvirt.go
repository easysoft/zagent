package kvmService

import "C"
import (
	"fmt"
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	"time"

	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
	"github.com/libvirt/libvirt-go"
)

const (
	LibvirtConnStrLocal  = "qemu:///system"
	LibvirtConnStrRemote = "qemu+ssh://%s:22/system?socket=/var/run/libvirt/libvirt-sock"
)

type LibvirtService struct {
	LibvirtConn *libvirt.Connect
	QemuService *QemuService `inject:""`
	VmService   *KvmService  `inject:""`
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

//func (s *LibvirtService) CreateVmFromTmpl(req *v1.KvmReq, removeSameName bool) (
//	dom *libvirt.Domain, vmVncPortMapped, vmAgentPortMapped int, vmRawPath, vmBackingPath string, err error) {
//
//	reqMsg, err := json.Marshal(req)
//	_logUtils.Infof("%s", reqMsg)
//
//	vmMacAddress := req.VmMacAddress
//	vmUniqueName := req.VmUniqueName
//	vmTemplateName := req.VmTemplate
//	vmCpu := req.VmCpu
//	vmMemorySize := req.VmMemorySize
//	vmDiskSize := req.VmDiskSize
//	vmBackingPath = req.VmBacking
//
//	if removeSameName {
//		s.SafeDestroyVmByName(vmUniqueName, true)
//	}
//
//	// gen tmpl xml definition
//	tmplXml := s.GetVmDef(vmTemplateName)
//	tmplDomCfg := &libvirtxml.Domain{}
//	err = tmplDomCfg.Unmarshal(tmplXml)
//	if tmplXml == "" || err != nil {
//		err = errors.New(fmt.Sprintf("get tmpl %s err", vmTemplateName))
//		return
//	}
//
//	// get Backing
//	if vmBackingPath == "" {
//		// use tmpl vm's image as backing path
//		mainDiskIndex := s.QemuService.GetMainDiskIndex(tmplDomCfg)
//		vmBackingPath = tmplDomCfg.Devices.Disks[mainDiskIndex].Source.File.File
//	} else {
//		vmBackingPath = filepath.Join(agentConf.Inst.DirKvm, vmBackingPath)
//	}
//
//	vmXml := ""
//	vmXml, vmRawPath, _ = s.QemuService.GenVmDef(tmplXml, vmMacAddress, vmUniqueName, vmBackingPath, vmCpu, vmMemorySize)
//	if err != nil {
//		_logUtils.Errorf("err gen vm xml, err %s", err.Error())
//		return
//	}
//
//	err = s.QemuService.createDiskFile(vmBackingPath, vmRawPath, vmDiskSize)
//	if err != nil {
//		_logUtils.Errorf(err.Error())
//		return
//	}
//
//	dom, err = s.LibvirtConn.DomainCreateXML(vmXml, libvirt.DOMAIN_NONE)
//	if err != nil {
//		_logUtils.Errorf(err.Error())
//		return
//	}
//
//	// get new vm info
//	newXml, _ := dom.GetXMLDesc(0)
//	newDomCfg := &libvirtxml.Domain{}
//	newDomCfg.Unmarshal(newXml)
//
//	vmMacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
//	vmVncPortMapped = newDomCfg.Devices.Graphics[0].VNC.Port
//
//	return
//}

//func (s *LibvirtService) CloneVm(req *v1.CloneVmReq, removeSameName bool) (dom *libvirt.Domain,
//	vmIp string, vmVncPort, vmAgentPortMapped int, vmRawPath, vmBackingPath string, err error) {
//
//	reqMsg, err := json.Marshal(req)
//	_logUtils.Infof("%s", reqMsg)
//
//	// get src vm config
//	tmplXml := s.GetVmDef(req.VmSrc)
//	srcDomCfg := &libvirtxml.Domain{}
//	err = srcDomCfg.Unmarshal(tmplXml)
//	if err != nil {
//		return
//	}
//
//	vmUniqueName := req.VmUniqueName
//	vmMacAddress := req.VmMacAddress
//	if removeSameName {
//		s.RemoveVmByName(vmUniqueName, true)
//	}
//
//	vmCpu := req.VmCpu
//	vmMemorySize := req.VmMemorySize
//	vmDiskSize := req.VmDiskSize
//
//	// update empty values from tmpl
//	if vmCpu == 0 {
//		vmCpu = srcDomCfg.VCPU.Value
//	}
//	if vmMemorySize == 0 {
//		vmCpu = srcDomCfg.Memory.Value
//	}
//
//	// get Backing
//	mainDiskIndex := s.QemuService.GetMainDiskIndex(srcDomCfg)
//	if vmBackingPath == "" {
//		// try to use the src vm's backing file
//		backingStore := srcDomCfg.Devices.Disks[mainDiskIndex].BackingStore
//		if backingStore != nil && backingStore.Source.File != nil {
//			vmBackingPath = backingStore.Source.File.File
//		}
//
//		// use src vm image as if src is a template without backing path
//		if vmBackingPath == "" {
//			vmBackingPath = srcDomCfg.Devices.Disks[mainDiskIndex].Source.File.File
//		}
//	}
//
//	// gen xml definition
//	vmXml := ""
//	vmXml, vmRawPath, _ = s.QemuService.GenVmDef(tmplXml, vmMacAddress, vmUniqueName, vmBackingPath, vmCpu, vmMemorySize)
//	if err != nil {
//		_logUtils.Errorf("err gen vm xml, err %s", err.Error())
//		return
//	}
//
//	err = s.QemuService.createDiskFile(vmBackingPath, vmRawPath, vmDiskSize)
//	if err != nil {
//		_logUtils.Errorf(err.Error())
//		return
//	}
//
//	dom, err = s.LibvirtConn.DomainCreateXML(vmXml, libvirt.DOMAIN_NONE)
//	if err != nil {
//		_logUtils.Errorf(err.Error())
//		return
//	}
//
//	// get new vm info
//	newXml, _ := dom.GetXMLDesc(0)
//	newDomCfg := &libvirtxml.Domain{}
//	newDomCfg.Unmarshal(newXml)
//
//	vmMacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
//	vmVncPort = newDomCfg.Devices.Graphics[0].VNC.Port
//	vmIp, _ = s.VmService.GetVmIpByMac(vmMacAddress)
//
//	return
//}

func (s *LibvirtService) ListVm() (domains []libvirt.Domain) {
	if s.LibvirtConn == nil {
		return
	}

	domains, err := s.LibvirtConn.ListAllDomains(0)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	return
}

func (s *LibvirtService) GetVm(name string) (dom *libvirt.Domain, err error) {
	dom, err = s.LibvirtConn.LookupDomainByName(name)
	if err != nil {
		return
	}

	return
}

func (s *LibvirtService) StartVmByName(name string) (err error) {
	dom, err := s.GetVm(name)

	err = dom.Create()

	return
}

func (s *LibvirtService) DestroyVm(dom *libvirt.Domain) (err error) {
	err = dom.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)

	return
}
func (s *LibvirtService) SafeDestroyVmByName(name string) (bizErr *domain.BizErr) {
	dom, err := s.GetVm(name)
	if err != nil {
		bizErr = &domain.ResultVmNotFound
		return
	}

	err = dom.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)
	if err != nil {
		libvirtErr := err.(libvirt.Error)

		if libvirtErr.Code != libvirt.ERR_OPERATION_INVALID {
			tmp := domain.NewBizErr(err.Error())
			bizErr = &tmp
		} else {
			bizErr = nil
		}
	}

	return
}
func (s *LibvirtService) ForceDestroyVmByName(name string) (bizErr *domain.BizErr) {
	dom, err := s.GetVm(name)
	if err != nil {
		bizErr = &domain.ResultVmNotFound
		return
	}

	err = dom.DestroyFlags(libvirt.DOMAIN_DESTROY_DEFAULT)
	if err != nil {
		tmp := domain.NewBizErr(err.Error())
		bizErr = &tmp
		return
	}

	return
}
func (s *LibvirtService) TryThenForceDestroyVmByName(name string) (bizErr *domain.BizErr) {
	bizErr = s.SafeDestroyVmByName(name)

	if bizErr != nil {
		bizErr = s.ForceDestroyVmByName(name)
	}

	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			bizErr = s.ForceDestroyVmByName(name)
			break
		default:
			dom, err := s.GetVm(name)
			if err == nil {
				domainState, _, _ := dom.GetState()
				if domainState == libvirt.DOMAIN_SHUTOFF || domainState == libvirt.DOMAIN_SHUTDOWN {
					break
				}
			}

			time.Sleep(1 * time.Second)
		}
	}

	return
}

func (s *LibvirtService) BootVmByName(name string) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT)
	return
}
func (s *LibvirtService) RebootVmByName(name string) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT)
	return
}
func (s *LibvirtService) ShutdownVmByName(name string) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Shutdown()
	return
}

func (s *LibvirtService) SuspendVmByName(name string) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Suspend()
	return
}
func (s *LibvirtService) ResumeVmByName(name string) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	err = dom.Resume()
	return
}

func (s *LibvirtService) RemoveVmByName(name string, removeDiskImage bool) (err error) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	dickPath, err := s.QemuService.GetDisk(dom)
	if err != nil {
		return
	}

	dom.DestroyFlags(libvirt.DOMAIN_DESTROY_DEFAULT)

	err = dom.Undefine()

	if removeDiskImage && dickPath != "" {
		_fileUtils.RemoveFile(dickPath)
	}

	return
}

//func (s *LibvirtService) ListTmpl() (ret []v1.KvmRespTempl, err error) {
//	if s.LibvirtConn == nil {
//		return
//	}
//
//	domains, err := s.LibvirtConn.ListAllDomains(0)
//	if err != nil {
//		_logUtils.Errorf(err.Error())
//		return
//	}
//
//	for _, domain := range domains {
//		name, _ := domain.GetName()
//		if strings.Index(name, "tmpl-") > -1 {
//			newXml := ""
//			newXml, err = domain.GetXMLDesc(0)
//			if err != nil {
//				continue
//			}
//
//			domainCfg := &libvirtxml.Domain{}
//			err = domainCfg.Unmarshal(newXml)
//
//			tmpl := v1.KvmRespTempl{
//				Name: domainCfg.Name,
//				Type: domainCfg.Type,
//				UUID: domainCfg.UUID,
//
//				CpuCoreNum:  domainCfg.VCPU.Value,
//				MemoryValue: domainCfg.Memory.Value,
//				MemoryUnit:  domainCfg.Memory.Unit,
//
//				OsArch:     domainCfg.OS.Type.Arch,
//				MacAddress: domainCfg.Devices.Interfaces[0].MAC.Address,
//			}
//
//			if domainCfg.Devices != nil && len(domainCfg.Devices.Graphics) > 0 {
//				for _, item := range domainCfg.Devices.Graphics {
//					if item.VNC != nil {
//						tmpl.VncPost = item.VNC.Port
//						break
//					}
//				}
//			}
//
//			mainDiskIndex := s.QemuService.GetMainDiskIndex(domainCfg)
//			tmpl.DiskFile = domainCfg.Devices.Disks[mainDiskIndex].Source.File.File
//
//			backingStore := domainCfg.Devices.Disks[mainDiskIndex].BackingStore
//			if backingStore != nil && backingStore.Source.File != nil {
//				tmpl.BackingFile = backingStore.Source.File.File
//				tmpl.BackingFormat = backingStore.Format.Type
//			}
//
//			ret = append(ret, tmpl)
//		}
//	}
//
//	return
//}

func (s *LibvirtService) GetVmDef(name string) (xml string) {
	dom, err := s.GetVm(name)
	if err != nil {
		return
	}

	xml, _ = dom.GetXMLDesc(0)

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
		osType.ToString(), osVersion, osLang.ToString(), _stringUtils.Uuid())
}

func (s *LibvirtService) Connect(str string) {
	if s.LibvirtConn != nil {
		active, err := s.LibvirtConn.IsAlive()
		if err != nil {
			_logUtils.Errorf(err.Error())
		}
		if active {
			return
		}
	}

	var err error
	s.LibvirtConn, err = libvirt.NewConnect(str)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	active, err := s.LibvirtConn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	if !active {
		_logUtils.Errorf("not active")
	}

	return
}

func (s *LibvirtService) IsAlive() (ret bool) {
	if s.LibvirtConn == nil {
		return false
	}

	ret, err := s.LibvirtConn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	if !ret {
		_logUtils.Errorf("connect is not active")
	}

	return
}

func (s *LibvirtService) GetVmSnaps(vm string) (ret []v1.SnapItemResp) {
	dom, err := s.GetVm(vm)
	if err != nil {
		return
	}

	snaps, _ := dom.ListAllSnapshots(0)

	for _, snap := range snaps {
		name, _ := snap.GetName()
		parent, _ := snap.GetParent(0)
		parentName, _ := parent.GetName()

		item := v1.SnapItemResp{
			Name:   name,
			Parent: parentName,
		}

		ret = append(ret, item)
	}

	return
}
