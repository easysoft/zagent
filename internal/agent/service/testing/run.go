package agentTestingService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type RunService struct {
	ExecService *ExecService `inject:""`

	SeleniumService *SeleniumService `inject:""`
	ScmService      *ScmService      `inject:""`
}

func NewRunService() *RunService {
	return &RunService{}
}

func (s *RunService) Run(build *commDomain.Build) {
	result := commDomain.TestResult{}
	result.Name = build.Name

	s.ExecService.SetBuildWorkDir(build)

	var err error

	// download res
	if build.BuildType == consts.SeleniumTest {
		err = s.SeleniumService.DownloadDriver(build)
		if err != nil {
			result.Failf("fail to download driver, err: %s", err.Error())
			s.ExecService.UploadResult(*build, result)
			return
		}

		build.EnvVars += "\nDriverType=" + build.BrowserType.ToString()
		build.EnvVars += "\nDriverVersion=" + build.BrowserVersion
		build.EnvVars += "\nDriverPath=" + build.SeleniumDriverPath
	}

	// set environment var
	err = s.ExecService.setEnvVars(build)
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
