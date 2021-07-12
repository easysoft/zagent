package hostKvmService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type VmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	LibvirtService *LibvirtService `inject:""`
}

func NewVmService() *VmService {
	s := VmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *VmService) UpdateVmMapAndDestroyTimeout(vms []domain.Vm) {
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

func (s *VmService) getKeys(m map[string]domain.Vm) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
