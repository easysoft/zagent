package handler

import (
	testingService "github.com/easysoft/zagent/internal/agent/service/testing"
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type PerformCtrl struct {
	TestService *testingService.TestService `inject:""`
}

func NewPerformCtrl() *PerformCtrl {
	return &PerformCtrl{}
}

func (c *PerformCtrl) Perform(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	if build.BuildType == consts.AutoSelenium {
		c.TestService.Run(&build)
	}

	reply.Success("success to add build.")

	return nil
}
