package vmwareService

import _logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"

type VMWareService struct {
	client *Client
}

func NewVMWareService() *VMWareService {
	return &VMWareService{}
}

func (s *VMWareService) CreateVm(id string) (err error) {
	//err = s.client.CreateVM(id)
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

func (s *VMWareService) Connect(url, account, password string) (err error) {
	s.client, err = NewClient(url, account, password, true, false)
	if err != nil {
		_logUtils.Errorf("NewClient error %s", err.Error())
	}

	return
}
