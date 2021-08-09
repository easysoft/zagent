package hostHandler

import (
	"fmt"
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmCtrl struct {
	VmService      *hostKvmService.VmService      `inject:""`
	LibvirtService *hostKvmService.LibvirtService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Create(ctx context.Context, req domain.KvmReq, reply *_domain.RpcResp) error {
	dom, vncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)
	if err == nil {
		reply.Pass("success to create vm.")

		vmName, _ := dom.GetName()
		vm := domain.Vm{
			Name:        vmName,
			VncAddress:  vncPort,
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
	c.LibvirtService.DestroyVmByName(req.VmUniqueName, true)

	reply.Passf("success to destroy vm %s .", req.VmUniqueName)
	return nil
}
