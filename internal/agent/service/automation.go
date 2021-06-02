package agentService

import (
	"fmt"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type AutomatedTestService struct {
	CommonService
	RegisterService *RegisterService `inject:""`
	ScmService      *ScmService      `inject:""`
	ExecService     *ExecService     `inject:""`
}

func NewAutomatedTestService() *AutomatedTestService {
	return &AutomatedTestService{}
}

func (s *AutomatedTestService) ExecTest(build *commDomain.Build) {
	result := _domain.RpcResp{}

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
			build.ProjectDir, build.BuildCommands))
	}

	// submit result
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
