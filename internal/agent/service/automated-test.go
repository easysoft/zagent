package agentService

import (
	"fmt"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type AutomatedTestService struct {
	CommonService
	ScmService  *ScmService           `inject:""`
	ExecService *AutomatedExecService `inject:""`
}

func NewAutomatedTestService() *AutomatedTestService {
	return &AutomatedTestService{}
}

func (s *AutomatedTestService) Exec(build *commDomain.Build) {
	result := commDomain.TestResult{}

	s.SetBuildWorkDir(build)

	// get script
	s.ScmService.GetTestScript(build)
	if build.ProjectDir == "" {
		result.Fail(fmt.Sprintf("failed to get test script, %#v。", build))
		return
	}

	// exec test
	prepareEnvVars(build)

	result = s.ExecService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n  cmd: \n%s",
			build.ProjectDir, build.BuildCommands))
	}

	// submit result
	result.Name = build.Name
	s.ExecService.UploadResult(*build, result)
}

func prepareEnvVars(build *commDomain.Build) {

}
