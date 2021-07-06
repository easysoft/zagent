package handler

import (
	"fmt"
	kvmService "github.com/easysoft/zagent/internal/agent/service/kvm"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmCtrl struct {
	VmService      *kvmService.VmService      `inject:""`
	LibvirtService *kvmService.LibvirtService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Create(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {
	dom, vncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req)
	if err == nil {
		reply.Success("success to create vm.")
	} else {
		reply.Fail(fmt.Sprintf("fail to create vm, error: %s", err.Error()))
	}

	vmName, _ := dom.GetName()
	vm := commDomain.Vm{
		Name:        vmName,
		VncPort:     vncPort,
		ImagePath:   vmRawPath,
		BackingPath: vmBackingPath,
	}

	reply.Payload = vm

	return nil
}

func (c *VmCtrl) Destroy(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {

	return nil
}
