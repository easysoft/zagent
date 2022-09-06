package handler

import (
	"github.com/easysoft/zv/internal/agent/service"
	"github.com/easysoft/zv/internal/agent/service/interface"
	"github.com/easysoft/zv/internal/comm/domain"
	_domain "github.com/easysoft/zv/pkg/domain"
	"golang.org/x/net/context"
)

type PerformCtrl struct {
	PerformService       *agentService.JobService                    `inject:""`
	InterfaceTestService *agentInterfaceService.InterfaceTestService `inject:""`
}

func NewPerformCtrl() *PerformCtrl {
	return &PerformCtrl{}
}

func (c *PerformCtrl) Perform(ctx context.Context, build domain.Build, reply *_domain.RemoteResp) error {
	result := domain.TestResult{}

	reply.Pass("Pass to exec processor.")
	reply.Payload = result

	return nil
}
