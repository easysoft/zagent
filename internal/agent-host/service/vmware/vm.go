package vmWareService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	vmwareService "github.com/easysoft/zagent/internal/server/service/vendors/vmware"
	"time"
)

type VmWareService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64
}

func NewVmWareService() *VmWareService {
	s := VmWareService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *VmWareService) CreateVm(req *domain.VmWareReq, removeSameName bool) (id, macAddress string, err error) {
	client := vmwareService.NewVMWareService() // 8697
	err = client.Connect(fmt.Sprintf("https://127.0.0.1:8697"), req.UserName, req.Password)
	if err != nil {
		_logUtils.Errorf("Connect to vmware err %s", err.Error())
		return
	}

	// create machine
	vmInst, err := client.CreateVm(req.BackingName, req.VmUniqueName, req.Processors, req.Memory)
	if err != nil {
		_logUtils.Errorf("Create vmware vm err %s", err.Error())
		return
	}

	id = vmInst.IdVM
	macAddress, _ = client.GetVmNic(id)

	return
}

func (s *VmWareService) DestroyVm(req *domain.VmWareReq, removeDiskImage bool) (err error) {
	client := vmwareService.NewVMWareService()
	err = client.Connect(fmt.Sprintf("https://127.0.0.1:8697"), req.UserName, req.Password)
	if err != nil {
		_logUtils.Errorf("Connect to vmware err %s", err.Error())
		return
	}

	client.DestroyVm(req.VmId)

	return
}
