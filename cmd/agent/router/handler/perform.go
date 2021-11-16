package handler

import (
	"github.com/easysoft/zagent/internal/agent/service"
	"github.com/easysoft/zagent/internal/agent/service/interface"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type PerformCtrl struct {
	PerformService       *agentService.JobService                    `inject:""`
	InterfaceTestService *agentInterfaceService.InterfaceTestService `inject:""`
}

func NewPerformCtrl() *PerformCtrl {
	return &PerformCtrl{}
}

func (c *PerformCtrl) Perform(ctx context.Context, build domain.Build, reply *_domain.RpcResp) error {
	result := domain.TestResult{}

	reply.Pass("Pass to exec processor.")
	reply.Payload = result

	return nil
}
