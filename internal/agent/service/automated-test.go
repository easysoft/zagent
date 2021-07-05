package agentService

import (
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
	result.Name = build.Name

	s.SetBuildWorkDir(build)

	var err error

	// download res
	if build.BuildType == consts.AutoSelenium {
		err = s.SeleniumService.DownloadDriver(build)
		if err != nil {
			result.Failf("fail to download driver, err: %s", err.Error())
			s.ExecService.UploadResult(*build, result)
			return
		}

		build.EnvVars += "\nDriverType=" + build.SeleniumDriverType.ToString()
		build.EnvVars += "\nDriverPath=" + build.SeleniumDriverPath
	}

	// set environment var
	err = setEnvVars(build)
	if err != nil {
		result.Failf("failed to set envs, err %s。", err.Error())
		s.ExecService.UploadResult(*build, result)
		return
	}

	// get script
	err = s.ScmService.GetTestScript(build)
	if err != nil || build.ProjectDir == "" {
		result.Failf("failed to get test script, err %s。", err.Error())
		s.ExecService.UploadResult(*build, result)
		return
	}

	// exec test
	err = s.ExecService.ExcCommand(build)
	if err != nil {
		result.Failf("failed to exec testing,\n dir: %s\n  cmd: \n%s, err: %s",
			build.ProjectDir, build.BuildCommands, err.Error())
		s.ExecService.UploadResult(*build, result)
		return
	}

	// submit result
	s.ExecService.UploadResult(*build, result)
}

func setEnvVars(build *commDomain.Build) (err error) {
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

		err = os.Setenv(name, val)

		if err != nil {
			break
		}
	}

	return
}
