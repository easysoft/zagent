package kvmService

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	natHelper "github.com/easysoft/zagent/internal/pkg/utils/net"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type KvmService struct {
	VmMapVar         map[string]domain.Vm
	SyncHeartbeatMap sync.Map
	TimeStamp        int64

	LibvirtService *LibvirtService `inject:""`
	QemuService    *QemuService    `inject:""`

	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewVmService() *KvmService {
	s := KvmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *KvmService) CreateVmFromImage(req *v1.CreateVmReq, removeSameName bool) (
	dom *libvirt.Domain, vmMacAddress string, vmVncPort int, err error) {

	vmName := req.Name
	srcFile := filepath.Join(agentConf.Inst.DirImage, vmName+".qcow2")
	targetBacking := req.Path
	cpuCores := req.Cpu
	ramSize := req.Memory
	diskSize := req.Disk

	if removeSameName {
		s.LibvirtService.SafeDestroyVmByName(vmName)
	}

	virtualSize := s.QemuService.GetBackingFileVirtualSize(targetBacking)
	if diskSize > virtualSize {
		virtualSize = diskSize
	}

	// create image
	cmdCreateImage := fmt.Sprintf(consts.CmdCreateImage, targetBacking, srcFile, virtualSize)
	out, err := _shellUtils.ExeShell(cmdCreateImage)
	if err != nil {
		_logUtils.Infof("exec cmd '%s' err, output %s, error %s", cmdCreateImage, out, err.Error())
		return
	}

	// create vm
	cmdCreateVm := fmt.Sprintf(consts.CmdCreateVm, vmName, cpuCores, ramSize*1000, srcFile, diskSize*1000)
	out, err = _shellUtils.ExeShell(cmdCreateVm)
	if err != nil {
		_logUtils.Infof("exec cmd '%s' err, output %s, error %s", cmdCreateVm, out, err.Error())
		return
	}

	// start vm
	err = s.StartVm(vmName)
	if err != nil {
		return
	}

	// get new vm info
	dom, err = s.LibvirtService.GetVm(vmName)
	if err != nil {
		return
	}

	newXml, _ := dom.GetXMLDesc(0)
	newDomCfg := &libvirtxml.Domain{}
	newDomCfg.Unmarshal(newXml)

	vmMacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
	vmVncPort = newDomCfg.Devices.Graphics[0].VNC.Port

	return
}

func (s *KvmService) StartVm(vmName string) error {
	cmdStartVm := fmt.Sprintf(consts.CmdStartVm, vmName)
	out, err := _shellUtils.ExeShell(cmdStartVm)

	if err != nil {
		_logUtils.Infof("exec cmd '%s' err, output %s, error %s", cmdStartVm, out, err.Error())
		return err
	}

	return nil
}

func (s *KvmService) GetVms() (vms []domain.Vm) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name, _ = dom.GetName()

		// if strings.Index(vm.Name, "test-") < 0 {
		// 	continue
		// }

		vm.Status = consts.VmUnknown
		domainState, _, _ := dom.GetState()
		if domainState == libvirt.DOMAIN_RUNNING {
			vm.Status = consts.VmRunning
		} else if domainState == libvirt.DOMAIN_SHUTOFF || domainState == libvirt.DOMAIN_SHUTDOWN {
			vm.Status = consts.VmShutOff
		} else if domainState == libvirt.DOMAIN_PAUSED {
			vm.Status = consts.VmPaused
		}

		newXml, _ := dom.GetXMLDesc(0)
		newDomCfg := &libvirtxml.Domain{}
		newDomCfg.Unmarshal(newXml)

		vm.MacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
		if newDomCfg.Devices.Graphics[0].VNC != nil {
			vm.VncPortOnHost = newDomCfg.Devices.Graphics[0].VNC.Port
		}
		vm.Ip, _ = s.GetVmIpByMac(vm.MacAddress)

		if vm.Status == consts.VmRunning && vm.Ip != "" {
			vm.AgentPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.AgentVmServicePort, consts.Http)
			vm.ZtfPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.ZtfServicePort, consts.Http)
			vm.ZdPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.ZdServicePort, consts.Http)

			vm.SshPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.SshServicePost, consts.Stream)
		}

		vms = append(vms, vm)
	}

	return vms
}

func (s *KvmService) GetIpByName(vmName string) (ip string) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		name, _ := dom.GetName()

		if name != vmName {
			continue
		}

		newXml, _ := dom.GetXMLDesc(0)
		newDomCfg := &libvirtxml.Domain{}
		newDomCfg.Unmarshal(newXml)

		macAddress := newDomCfg.Devices.Interfaces[0].MAC.Address
		ip, _ = s.GetVmIpByMac(macAddress)
		break
	}
	return
}

func (s *KvmService) AddExportVmTask(req v1.ExportVmReq) (err error) {
	if len(req.Backing) < 6 || req.Backing[len(req.Backing)-6:] != ".qcow2" {
		req.Backing += ".qcow2"
	}

	po := agentModel.Task{
		Vm:      req.Vm,
		Backing: req.Backing,

		Task:   req.Task,
		Type:   consts.ExportVm,
		Status: consts.Created,
	}

	s.TaskRepo.Save(&po)

	return
}

func (s *KvmService) UpdateVmMapAndDestroyTimeout(vms []domain.Vm) {
	names := map[string]bool{}

	for key, vm := range vms {
		name := vm.Name
		names[name] = true

		if _, ok := s.VmMapVar[name]; ok { // update status in map
			v := s.VmMapVar[name]
			v.Status = vm.Status
			s.VmMapVar[name] = v
			vms[key].Heartbeat = s.GetHeartbeat(v.MacAddress)
		} else { // update time then add
			if vm.FirstDetectedTime.IsZero() {
				vm.FirstDetectedTime = time.Now()
			}
			s.VmMapVar[name] = vm
		}
	}

	keys := s.getKeys(s.VmMapVar)
	for _, key := range keys {
		if !names[key] { // remove vm in map but not found this time
			delete(s.VmMapVar, key)
			continue
		}

		// destroy and remove timeout vm
		v := s.VmMapVar[key]
		if time.Now().Unix()-v.FirstDetectedTime.Unix() > consts.WaitVmLifecycleTimeout { // timeout
			s.LibvirtService.SafeDestroyVmByName(v.Name)
			delete(s.VmMapVar, key)
		}
	}
}

func (s *KvmService) UpdateHeartbeat(mac string) {
	s.SyncHeartbeatMap.Store(mac, time.Now())
}

func (s *KvmService) GetHeartbeat(mac string) (val time.Time) {
	inf, ok := s.SyncHeartbeatMap.Load(mac)
	if ok {
		val = inf.(time.Time)
	}
	return
}

func (s *KvmService) getKeys(m map[string]domain.Vm) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (s *KvmService) GetVmIpByMac(macAddress string) (ip string, err error) {
	cmd := `virsh net-dhcp-leases default | grep ipv4 | awk '{print $3,$5 }'`

	out, err := _shellUtils.ExeSysCmd(cmd)
	arr := strings.Split(out, "\n")

	for _, line := range arr {
		cols := strings.Split(line, " ")
		if len(cols) == 0 {
			continue
		}

		if strings.TrimSpace(cols[0]) == macAddress {
			ip = strings.Split(strings.TrimSpace(cols[1]), "/")[0]
			ip = strings.TrimSpace(ip)

			break
		}
	}

	return
}
