package vmwareService

type VMWareService struct {
	client *Client
}

func NewVMWareService() *VMWareService {
	return &VMWareService{}
}

func (s *VMWareService) GetVms() (vms []Vm, err error) {
	vms, err = s.client.GetAllVMs()
	return
}

func (s *VMWareService) Connect(url, account, password string) (err error) {
	s.client, err = NewClient(url, account, password, true, true)

	return
}
