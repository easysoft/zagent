package agentService

import (
	"encoding/json"
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/libs/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/libs/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	_shellUtils "github.com/easysoft/zagent/internal/pkg/libs/shell"
	"github.com/satori/go.uuid"
	"strings"
)

type ExecService struct {
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s *ExecService) ExcCommand(build *_domain.BuildTo) _domain.RpcResult {
	cmdStr := build.BuildCommands
	out, err := _shellUtils.ExeShellWithOutputInDir(cmdStr, build.ProjectDir)

	result := _domain.RpcResult{}
	if err == nil {
		result.Success(strings.Join(out, "\n"))
	} else {
		result.Fail(fmt.Sprintf("fail to exec command, out %s, errpr %#v", out, err))
	}

	return result
}

func (s *ExecService) UploadResult(build _domain.BuildTo, result _domain.RpcResult) {
	zipFile := build.WorkDir + "testResult.zip"
	err := _fileUtils.ZipFiles(zipFile, build.ProjectDir, strings.Split(build.ResultFiles, ","))
	if err != nil {
		_logUtils.Errorf("fail to zip test results, dist '%s', dir %s, files '%s', error %#v",
			zipFile, build.ProjectDir, build.ResultFiles, err)
	}

	result.Payload = build

	uploadResultUrl := _httpUtils.GenUrl(agentConf.Inst.Server, "build/upload")
	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}

func (s *ExecService) GetTestApp(build *_domain.BuildTo) _domain.RpcResult {
	result := _domain.RpcResult{}

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

func (s *ExecService) DownloadApp(build *_domain.BuildTo) {
	path := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.AppUrl)
	_fileUtils.Download(build.AppUrl, path)
	build.AppPath = path
}
