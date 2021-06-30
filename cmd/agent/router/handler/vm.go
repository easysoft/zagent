package handler

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmCtrl struct {
	VmService *agentService.VmService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Create(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {

	//reply.Success("Success to add job.")
	//reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))

	reply.Payload = "vm"

	return nil
}

func (c *VmCtrl) Destroy(ctx context.Context, req commDomain.KvmReq, reply *_domain.RpcResp) error {

	return nil
}
