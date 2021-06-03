package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type InterfaceTestService struct {
	CommonService
	ExecService *InterfaceExecService `inject:""`
}

func NewInterfaceTestService() *InterfaceTestService {
	return &InterfaceTestService{}
}

func (s *InterfaceTestService) ExecScenario(build *commDomain.Build) {
	result := commDomain.TestResult{}

	scenario := build.TestScenario

	result.Name = scenario.Name
	s.ExecService.UploadResult(*build, result)
}

func (s *InterfaceTestService) ExecSet(build *commDomain.Build) {
	result := commDomain.TestResult{}

	set := build.TestSet

	// TODO:

	result.Name = set.Name
	s.ExecService.UploadResult(*build, result)
}
