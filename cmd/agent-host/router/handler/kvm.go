package hostHandler

import (
	"fmt"
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
	"strconv"
)

type KvmCtrl struct {
	VmService      *hostKvmService.VmService      `inject:""`
	LibvirtService *hostKvmService.LibvirtService `inject:""`
}

func NewKvmCtrl() *KvmCtrl {
	return &KvmCtrl{}
}

func (c *KvmCtrl) Create(ctx context.Context, req domain.KvmReq, reply *_domain.RpcResp) error {

	dom, vmVncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)
	if err == nil {
		reply.Pass("success to create vm.")

		vmName, _ := dom.GetName()
		vm := domain.Vm{
			Name:        vmName,
			VncAddress:  strconv.Itoa(vmVncPort),
			ImagePath:   vmRawPath,
			BackingPath: vmBackingPath,
		}

		reply.Payload = vm

	} else {
		reply.Fail(fmt.Sprintf("fail to create vm, error: %s", err.Error()))
	}

	return nil
}

func (c *KvmCtrl) Destroy(ctx context.Context, req domain.KvmReq, reply *_domain.RpcResp) error {
	c.LibvirtService.DestroyVmByName(req.VmUniqueName, true)

	reply.Passf("success to destroy vm %s .", req.VmUniqueName)
	return nil
}
