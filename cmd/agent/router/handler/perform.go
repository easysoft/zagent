package handler

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	interfaceService "github.com/easysoft/zagent/internal/agent/service/interface"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type PerformCtrl struct {
	PerformService       *agentService.JobService               `inject:""`
	InterfaceTestService *interfaceService.InterfaceTestService `inject:""`
}

func NewPerformCtrl() *PerformCtrl {
	return &PerformCtrl{}
}

func (c *PerformCtrl) Perform(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	result := commDomain.TestResult{}

	reply.Pass("Pass to exec processor.")
	reply.Payload = result

	return nil
}
