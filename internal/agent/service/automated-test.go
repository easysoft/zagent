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
	parseBuildCommand(build)
	result = s.ExecService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n  cmd: \n%s",
			build.ProjectDir, build.AutomatedTest.BuildCommands))
	}

	// submit result
	result.Name = build.AutomatedTest.Name
	s.ExecService.UploadResult(*build, result)
}

func parseBuildCommand(build *commDomain.Build) {
	// mvn clean test -Dtestng.suite=target/test-classes/baidu-test.xml
	//		 		  -DdriverPath=${driverPath}

	//dir := ""
	//if _commonUtils.IsWin() {
	//	dir = agentConst.ResPathWin
	//} else {
	//	dir = agentConst.ResPathLinux
	//}
	//build.BuildCommands = build.BuildCommands
}
