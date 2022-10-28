package kvmService

import (
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentModel "github.com/easysoft/zv/internal/host/model"
	hostRepo "github.com/easysoft/zv/internal/host/repo"
	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	natHelper "github.com/easysoft/zv/internal/pkg/utils/net"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"path/filepath"
	"strings"
	"time"
)

const ()

type KvmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

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
		s.LibvirtService.DestroyVmByName(vmName, true)
	}

	cmdCreateImage := fmt.Sprintf(consts.CmdCreateImage, targetBacking, srcFile)
	_shellUtils.ExeShell(cmdCreateImage)

	cmdCreateVm := fmt.Sprintf(consts.CmdCreateVm, vmName, cpuCores, ramSize, diskSize)
	_shellUtils.ExeShell(cmdCreateVm)

	cmdStartVm := fmt.Sprintf(consts.CmdStartVm, vmName)
	_shellUtils.ExeShell(cmdStartVm)

	dom, err = s.LibvirtService.GetVm(vmName)
	if err != nil {
		return
	}

	// get new vm info
	newXml, _ := dom.GetXMLDesc(0)
	newDomCfg := &libvirtxml.Domain{}
	newDomCfg.Unmarshal(newXml)

	vmMacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
	vmVncPort = newDomCfg.Devices.Graphics[0].VNC.Port

	return
}

func (s *KvmService) GetVms() (vms []domain.Vm) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name, _ = dom.GetName()

		if strings.Index(vm.Name, "test-") < 0 {
			continue
		}

		vm.Status = consts.VmUnknown
		domainState, _, _ := dom.GetState()
		if domainState == libvirt.DOMAIN_RUNNING {
			vm.Status = consts.VmRunning
		} else if domainState == libvirt.DOMAIN_SHUTOFF || domainState == libvirt.DOMAIN_SHUTDOWN {
			vm.Status = consts.VmShutOff
		}

		newXml, _ := dom.GetXMLDesc(0)
		newDomCfg := &libvirtxml.Domain{}
		newDomCfg.Unmarshal(newXml)

		vm.MacAddress = newDomCfg.Devices.Interfaces[0].MAC.Address
		vm.VncPortOnHost = newDomCfg.Devices.Graphics[0].VNC.Port
		vm.Ip, _ = s.GetVmIpByMac(vm.MacAddress)

		if vm.Status == consts.VmRunning && vm.Ip != "" {
			vm.AgentPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.AgentServicePost, consts.Http)
			vm.SshPortOnHost, _, _ = natHelper.ForwardPortIfNeeded(vm.Ip, consts.SshServicePost, consts.Stream)
		}

		vms = append(vms, vm)
	}

	return vms
}

func (s *KvmService) AddExportVmTask(req v1.ExportVmReq) (resp v1.ExportVmReq, err error) {
	po := agentModel.Task{
		Vm:      req.Vm,
		Backing: req.Backing,

		ZentaoTask: req.ZentaoTask,
		TaskType:   consts.ExportVm,
	}

	s.TaskRepo.Save(&po)

	return
}

func (s *KvmService) UpdateVmMapAndDestroyTimeout(vms []domain.Vm) {
	names := map[string]bool{}

	for _, vm := range vms {
		name := vm.Name
		names[name] = true

		if _, ok := s.VmMapVar[name]; ok { // update status in map
			v := s.VmMapVar[name]
			v.Status = vm.Status
			s.VmMapVar[name] = v
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
			s.LibvirtService.DestroyVmByName(v.Name, true)
			delete(s.VmMapVar, key)
		}
	}
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

			break
		}
	}

	return
}
