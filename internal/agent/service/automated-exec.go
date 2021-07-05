package agentService

import (
	"encoding/json"
	"errors"
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/lib/shell"
	"github.com/satori/go.uuid"
	"strings"
)

type AutomatedExecService struct {
}

func NewAutomatedExecService() *AutomatedExecService {
	return &AutomatedExecService{}
}

func (s *AutomatedExecService) ExcCommand(build *commDomain.Build) (err error) {
	cmdStr := build.BuildCommands
	var out []string
	out, err = _shellUtils.ExeShellWithOutputInDir(cmdStr, build.ProjectDir)

	if err != nil {
		errors.New(fmt.Sprintf("fail to exec command, out %s, err: %s", out, err.Error()))
	}

	return
}

func (s *AutomatedExecService) GetTestApp(build *commDomain.Build) _domain.RpcResp {
	result := _domain.RpcResp{}

	if strings.Index(build.AppUrl, "http://") == 0 {
		s.DownloadApp(build)
	} else {
		build.AppPath = build.AppUrl
	}

	if build.AppPath == "" {
		result.Fail(fmt.Sprintf("App获取错误，%s", build.AppUrl))
	} else {
		result.Success("")
	}

	return result
}

func (s *AutomatedExecService) DownloadApp(build *commDomain.Build) {
	path := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.AppUrl)
	_fileUtils.Download(build.AppUrl, path)
	build.AppPath = path
}

func (s *AutomatedExecService) UploadResult(build commDomain.Build, result commDomain.TestResult) {
	zipFile := build.WorkDir + "testResult.zip"
	err := _fileUtils.ZipFiles(zipFile, build.ProjectDir, strings.Split(build.ResultFiles, ","))
	if err != nil {
		_logUtils.Errorf(_i118Utils.Sprintf("fail_to_zip_test_result",
			zipFile, build.ProjectDir, build.ResultFiles, err))
	}

	result.Payload = build

	uploadResultUrl := _httpUtils.GenUrl(agentConf.Inst.Server, "build/upload")
	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}
