package multiPassService

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	"time"
)

type VmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	MultiPassService *MultiPassService `inject:""`
}

func NewVmService() *VmService {
	s := VmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *VmService) GetVmsInfos() (vms []domain.Vm) {
	domains, _ := s.MultiPassService.GetVms()

	for _, dom := range domains {
		vm := domain.Vm{}
		vm.Name = dom.Name

		vm.Status = consts.VmStatus(dom.State)
		if dom.State == "Running" {
			vm.Status = consts.VmRunning
		} else if dom.State == "Deleted" || dom.State == "Suspended" {
			vm.Status = consts.VmShutOff
		}

		vms = append(vms, vm)
	}

	return vms
}
