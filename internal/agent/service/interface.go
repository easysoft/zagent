package agentService

import (
	"fmt"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type InterfaceService struct {
	CommonService
	RegisterService   *RegisterService      `inject:""`
	ScmService   *ScmService      `inject:""`
	ExecService   *ExecService      `inject:""`
}

func NewInterfaceService() *InterfaceService {
	return &InterfaceService{}
}

func (s *InterfaceService) ExecTest(build *_domain.BuildTo) {
	result := _domain.RpcResult{}

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

func parseBuildCommand(build *_domain.BuildTo) {
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
