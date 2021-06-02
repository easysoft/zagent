package agentService

import (
	"github.com/easysoft/zagent/internal/agent/model"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/libs/file"
	_gitUtils "github.com/easysoft/zagent/internal/pkg/libs/git"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mholt/archiver/v3"
	"github.com/satori/go.uuid"
	"os"
	"strings"
)

type ScmService struct {
}

func NewScmService() *ScmService {
	return &ScmService{}
}

func (s *ScmService) GetTestScript(build *domain.BuildTo) _domain.RpcResp {
	if build.ScmAddress != "" {
		CheckoutCodes(build)
	} else if strings.Index(build.ScriptUrl, "http://") == 0 {
		DownloadCodes(build)
	} else {
		build.ProjectDir = _fileUtils.AddPathSepIfNeeded(build.ScriptUrl)
	}

	result := _domain.RpcResp{}
	result.Success("")
	return result
}

func CheckoutCodes(task *domain.BuildTo) {
	url := task.ScmAddress
	userName := task.ScmAccount
	password := task.ScmPassword

	projectDir := task.WorkDir + _gitUtils.GetGitProjectName(url) + _const.PthSep

	_fileUtils.MkDirIfNeeded(projectDir)

	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	if userName != "" {
		options.Auth = &http.BasicAuth{
			Username: userName,
			Password: password,
		}
	}
	_, err := git.PlainClone(projectDir, false, &options)
	if err != nil {
		return
	}

	task.ProjectDir = projectDir
}

func DownloadCodes(task *domain.BuildTo) {
	zipPath := task.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(task.ScriptUrl)
	_fileUtils.Download(task.ScriptUrl, zipPath)

	scriptFolder := _fileUtils.GetZipSingleDir(zipPath)
	if scriptFolder != "" { // single dir in zip
		task.ProjectDir = task.WorkDir + scriptFolder
		archiver.Unarchive(zipPath, task.WorkDir)
	} else { // more then one dir, unzip to a folder
		fileNameWithoutExt := _fileUtils.GetFileNameWithoutExt(zipPath)
		task.ProjectDir = task.WorkDir + fileNameWithoutExt + _const.PthSep
		archiver.Unarchive(zipPath, task.ProjectDir)
	}
}
