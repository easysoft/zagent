package agentService

import (
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"os"
	"strings"
)

type AutomatedTestService struct {
	CommonService

	SeleniumService *SeleniumService      `inject:""`
	ScmService      *ScmService           `inject:""`
	ExecService     *AutomatedExecService `inject:""`
}

func NewAutomatedTestService() *AutomatedTestService {
	return &AutomatedTestService{}
}

func (s *AutomatedTestService) Exec(build *commDomain.Build) {
	result := commDomain.TestResult{}

	s.SetBuildWorkDir(build)

	// download res
	if build.BuildType == consts.AutoSelenium {
		s.SeleniumService.DownloadDriver(build)

		build.EnvVars += "\nDriverType=" + build.SeleniumDriverType.ToString()
		build.EnvVars += "\nDriverPath=" + build.SeleniumDriverPath
	}

	// set environment var
	setEnvVars(build)

	// get script
	s.ScmService.GetTestScript(build)
	if build.ProjectDir == "" {
		result.Fail(fmt.Sprintf("failed to get test script, %#v。", build))
		return
	}

	// exec test
	result = s.ExecService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n  cmd: \n%s",
			build.ProjectDir, build.BuildCommands))
	}

	// submit result
	result.Name = build.Name
	s.ExecService.UploadResult(*build, result)
}

func setEnvVars(build *commDomain.Build) {
	for _, env := range strings.Split(build.EnvVars, "\n") {
		arr := strings.Split(env, "=")
		if len(arr) < 2 {
			continue
		}

		name := strings.TrimSpace(arr[0])
		val := strings.TrimSpace(arr[1])
		if name == "" || val == "" {
			continue
		}

		os.Setenv(name, val)
	}
}
