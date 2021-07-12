package handler

import (
	vmAgentService "github.com/easysoft/zagent/internal/agent-vm/service"
	interfaceService "github.com/easysoft/zagent/internal/agent/service/interface"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type PerformCtrl struct {
	PerformService       *vmAgentService.JobService             `inject:""`
	InterfaceTestService *interfaceService.InterfaceTestService `inject:""`
}

func NewPerformCtrl() *PerformCtrl {
	return &PerformCtrl{}
}

func (c *PerformCtrl) Perform(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	result := commDomain.TestResult{}

	if build.BuildType == commConst.InterfaceScenario {
		result = c.InterfaceTestService.ExecScenario(&build)
	} else if build.BuildType == commConst.InterfaceSet {
		c.InterfaceTestService.ExecSet(&build, &result)
	}

	reply.Pass("Pass to exec processor.")
	reply.Payload = result

	return nil
}
