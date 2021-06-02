package agentService

import (
	"encoding/json"
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
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

func (s *ExecService) ExcCommand(build *commDomain.BuildTo) _domain.RpcResp {
	cmdStr := build.BuildCommands
	out, err := _shellUtils.ExeShellWithOutputInDir(cmdStr, build.ProjectDir)

	result := _domain.RpcResp{}
	if err == nil {
		result.Success(strings.Join(out, "\n"))
	} else {
		result.Fail(fmt.Sprintf("fail to exec command, out %s, errpr %#v", out, err))
	}

	return result
}

func (s *ExecService) UploadResult(build commDomain.BuildTo, result _domain.RpcResp) {
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

func (s *ExecService) GetTestApp(build *commDomain.BuildTo) _domain.RpcResp {
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

func (s *ExecService) DownloadApp(build *commDomain.BuildTo) {
	path := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.AppUrl)
	_fileUtils.Download(build.AppUrl, path)
	build.AppPath = path
}
