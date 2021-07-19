package agentTestingService

import (
	"encoding/json"
	"errors"
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/lib/shell"
	"github.com/satori/go.uuid"
	"os"
	"strings"
)

type ExecService struct {
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s *ExecService) ExcCommand(build *commDomain.Build) (err error) {
	cmdStr := build.BuildCommands
	envVars := s.parseEnvVars(build.EnvVars)

	var out []string
	out, err = _shellUtils.ExeShellWithEnvVarsAndOutputInDir(cmdStr, build.ProjectDir, envVars)

	if err != nil {
		errors.New(fmt.Sprintf("fail to exec command, out %s, err: %s", out, err.Error()))
	}

	return
}

func (s *ExecService) GetTestApp(build *commDomain.Build) _domain.RpcResp {
	result := _domain.RpcResp{}

	if strings.Index(build.AppUrl, "http://") == 0 {
		s.DownloadApp(build)
	} else {
		build.AppPath = build.AppUrl
	}

	if build.AppPath == "" {
		result.Fail(fmt.Sprintf("App获取错误，%s", build.AppUrl))
	} else {
		result.Pass("")
	}

	return result
}

func (s *ExecService) DownloadApp(build *commDomain.Build) {
	path := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.AppUrl)
	_fileUtils.Download(build.AppUrl, path)
	build.AppPath = path
}

func (s *ExecService) UploadResult(build commDomain.Build, result commDomain.TestResult) {
	zipFile := build.WorkDir + "testResult.zip"
	err := _fileUtils.ZipFiles(zipFile, build.ProjectDir, strings.Split(build.ResultFiles, ","))
	if err != nil {
		_logUtils.Errorf(_i118Utils.Sprintf("fail_to_zip_test_result",
			zipFile, build.ProjectDir, build.ResultFiles, err))
	}

	result.Payload = build

	uploadResultUrl := _httpUtils.GenUrl(agentConf.Inst.Server, "client/build/uploadResult")

	// add form field
	extraParams := map[string]string{}

	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, []string{zipFile}, extraParams)
}

func (s *ExecService) SetBuildWorkDir(build *commDomain.Build) {
	build.WorkDir = agentConf.Inst.WorkDir + uuid.NewV4().String() + _const.PthSep
	_fileUtils.MkDirIfNeeded(build.WorkDir)
}

func (s *ExecService) setEnvVars(build *commDomain.Build) (err error) {
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

func (s *ExecService) parseEnvVars(vars string) (ret []string) {
	arr := strings.Split(vars, "\n")

	for _, item := range arr {
		str := strings.TrimSpace(item)
		if str == "" {
			continue
		}

		ret = append(ret, str)
	}

	return
}
