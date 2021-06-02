package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type InterfaceTestService struct {
	CommonService
	RegisterService *RegisterService `inject:""`
	ScmService      *ScmService      `inject:""`
	ExecService     *ExecService     `inject:""`
}

func NewInterfaceTestService() *InterfaceTestService {
	return &InterfaceTestService{}
}

func (s *InterfaceTestService) ExecTest(build *commDomain.Build) {
	result := _domain.RpcResp{}

	// TODO:

	s.ExecService.UploadResult(*build, result)
}
