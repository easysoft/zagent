package handler

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmCtrl struct {
	VmService      *agentService.VmService      `inject:""`
	LibvirtService *agentService.LibvirtService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Create(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {
	dom, vncPort, err := c.LibvirtService.CreateVm(&req)
	if err == nil {
		reply.Success("success to create vm.")
	} else {
		reply.Fail(fmt.Sprintf("fail to create vm, error: %s", err.Error()))
	}

	vmName, _ := dom.GetName()
	reply.Payload = map[string]interface{}{"vmName": vmName, "vncPort": vncPort}

	return nil
}

func (c *VmCtrl) Destroy(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {

	return nil
}
