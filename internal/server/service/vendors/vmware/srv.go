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

func (s *VMWareService) CreateVm(tmpl, name string, processors, memory uint) (vm *Vm, err error) {
	//get vm by name
	vms, _ := s.GetVms()

	tmplId := ""
	for _, vm := range vms {
		if tmpl == vm.Denomination {
			tmplId = vm.IdVM
			break
		}
	}
	if tmplId == "" {
		msg := "vm %s not found"
		_logUtils.Errorf(msg, name)
		err = errors.New(msg)
		return
	}

	vm, err = s.client.CreateVM(tmplId, name, "")
	if err != nil {
		_logUtils.Errorf("DestroyVM error %s", err.Error())
	}

	//// get vm path
	//vms, _ = s.GetVms()
	//path := ""
	//for _, vm := range vms {
	//	if name == vm.Denomination {
	//		path = vm.Path
	//		break
	//	}
	//}
	//if path == "" {
	//	msg := "vm %s path not found"
	//	_logUtils.Errorf(msg, name)
	//	err = errors.New(msg)
	//	return
	//}
	//
	//// register vm
	//vm, err = s.client.RegisterVM(name, path)
	//if err != nil {
	//	_logUtils.Errorf("RegisterVM error %s", err.Error())
	//}

	// set cpu and memory
	vm, err = s.client.UpdateVM(vm.IdVM, "", "", processors, memory)
	if err != nil {
		_logUtils.Errorf("UpdateVM error %s", err.Error())
	}

	err = s.client.PowerOn(vm.IdVM)
	if err != nil {
		_logUtils.Errorf("PowerOn error %s", err.Error())
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
