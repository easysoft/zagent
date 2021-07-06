package handler

import (
	"fmt"
	"github.com/easysoft/zagent/internal/agent/service/kvm"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmCtrl struct {
	VmService      *kvmService.VmService      `inject:""`
	LibvirtService *kvmService.LibvirtService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Create(ctx context.Context, req domain.KvmReq, reply *_domain.RpcResp) error {
	dom, vncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)
	if err == nil {
		reply.Success("success to create vm.")

		vmName, _ := dom.GetName()
		vm := domain.Vm{
			Name:        vmName,
			VncPort:     vncPort,
			ImagePath:   vmRawPath,
			BackingPath: vmBackingPath,
		}

		reply.Payload = vm

	} else {
		reply.Fail(fmt.Sprintf("fail to create vm, error: %s", err.Error()))
	}

	return nil
}

func (c *VmCtrl) Destroy(ctx context.Context, req domain.KvmReq, reply *_domain.RpcResp) error {

	return nil
}
