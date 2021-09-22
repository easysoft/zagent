package vmwareService

import (
	"errors"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
)

type VMWareService struct {
	client *Client
}

func NewVMWareService() *VMWareService {
	return &VMWareService{}
}

func (s *VMWareService) CreateVm(tmpl, name string) (vm *Vm, err error) {
	vms, _ := s.GetVms()

	tmplId := ""
	for _, vm := range vms {
		if tmpl == vm.Denomination {
			tmplId = vm.IdVM
			break
		}
	}
	if tmplId == "" {
		msg := "vm %S not found"
		_logUtils.Errorf(msg, name)
		err = errors.New(msg)
		return
	}

	vm, err = s.client.CreateVM(tmplId, name, "")
	if err != nil {
		_logUtils.Errorf("DestroyVM error %s", err.Error())
	}
	return
}

func (s *VMWareService) DestroyVm(id string) (err error) {
	err = s.client.DestroyVM(id)
	if err != nil {
		_logUtils.Errorf("DestroyVM error %s", err.Error())
	}
	return
}

func (s *VMWareService) GetVms() (vms []Vm, err error) {
	vms, err = s.client.GetAllVMs()
	if err != nil {
		_logUtils.Errorf("GetAllVMs error %s", err.Error())
	}
	return
}

func (s *VMWareService) GetVmNic(id string) (macAddress string, err error) {
	nic, err := s.client.GetVmNic(id)
	if err != nil {
		_logUtils.Errorf("GetVmNic error %s", err.Error())
	}

	macAddress = nic.MacAddress

	return
}

func (s *VMWareService) Connect(url, account, password string) (err error) {
	s.client, err = NewClient(url, account, password, true, false)
	if err != nil {
		_logUtils.Errorf("NewClient error %s", err.Error())
	}

	return
}
