package testingService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type TestService struct {
	ExecService *ExecService `inject:""`

	SeleniumService *SeleniumService `inject:""`
	ScmService      *ScmService      `inject:""`
}

func NewService() *TestService {
	return &TestService{}
}

func (s *TestService) Run(build *commDomain.Build) {
	result := commDomain.TestResult{}
	result.Name = build.Name

	s.ExecService.SetBuildWorkDir(build)

	var err error

	// download res
	if build.BuildType == consts.AutoSelenium {
		err = s.SeleniumService.DownloadDriver(build)
		if err != nil {
			result.Failf("fail to download driver, err: %s", err.Error())
			s.ExecService.UploadResult(*build, result)
			return
		}

		build.EnvVars += "\nDriverType=" + build.BrowserType.ToString()
		build.EnvVars += "\nDriverPath=" + build.SeleniumDriverPath
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
