package kvmService

import (
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"strings"
	"time"
)

type KvmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	LibvirtService *LibvirtService `inject:""`
}

func NewVmService() *KvmService {
	s := KvmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *KvmService) GetVms() (vms []domain.Vm) {
	domains := s.LibvirtService.ListVm()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name, _ = dom.GetName()

		vm.Name = strings.Replace(vm.Name, "tmpl-", "test-", -1)

		if strings.Index(vm.Name, "test-") != 0 {
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
		vm.Ip, _ = s.GetVmIpByMac(vm.MacAddress)

		vms = append(vms, vm)
	}

	return vms
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
		if strings.TrimSpace(cols[0]) == macAddress {
			ip = strings.Split(strings.TrimSpace(cols[1]), "/")[0]

			break
		}
	}

	return
}
