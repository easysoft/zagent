package agentInterfaceService

import (
	commDomain "github.com/easysoft/zv/internal/comm/domain"
)

type InterfaceTestService struct {
	ExecService *InterfaceExecService `inject:""`
}

func NewInterfaceTestService() *InterfaceTestService {
	return &InterfaceTestService{}
}

func (s *InterfaceTestService) ExecScenario(build *commDomain.Build) (result commDomain.TestResult) {
	scenario := build.TestScenario
	s.ExecService.ExecProcessor(build, &scenario.Processor)

	// TODO: deal with result with logs in scenario.Processor

	result.Name = scenario.Name

	return
}

func (s *InterfaceTestService) ExecSet(build *commDomain.Build, result *commDomain.TestResult) {
	set := build.TestSet

	// TODO:

	result.Name = set.Name
}
