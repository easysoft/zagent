package hostHandler

import (
	"fmt"
	vmWareService "github.com/easysoft/zagent/internal/agent-host/service/vmware"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmWareCtrl struct {
	VmWareService *vmWareService.VmWareService `inject:""`
}

func NewVmWareCtrl() *VmWareCtrl {
	return &VmWareCtrl{}
}

func (c *VmWareCtrl) Create(ctx context.Context, req domain.VmWareReq, reply *_domain.RpcResp) error {

	id, macAddress, err := c.VmWareService.CreateVm(&req, true)
	if err == nil {
		reply.Pass("success to create VmWare.")

		VmWare := domain.VmWareResp{
			VmId: id,
			Name: req.VmUniqueName,
			Mac:  macAddress,
			//VncAddress:  strconv.Itoa(VmWareVncPort),
		}

		reply.Payload = VmWare

	} else {
		reply.Fail(fmt.Sprintf("fail to create VmWare, error: %s", err.Error()))
	}

	return nil
}

func (c *VmWareCtrl) Destroy(ctx context.Context, req domain.VmWareReq, reply *_domain.RpcResp) error {
	c.VmWareService.DestroyVm(&req, true)

	reply.Passf("success to destroy VmWare %s .", req.VmId)
	return nil
}
