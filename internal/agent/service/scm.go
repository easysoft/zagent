package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_gitUtils "github.com/easysoft/zagent/internal/pkg/lib/git"
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

func (s *ScmService) GetTestScript(build *commDomain.Build) _domain.RpcResp {
	if build.AutomatedTest.ScmAddress != "" {
		CheckoutCodes(build)
	} else if strings.Index(build.AutomatedTest.ScriptUrl, "http://") == 0 {
		DownloadCodes(build)
	} else {
		build.ProjectDir = _fileUtils.AddPathSepIfNeeded(build.AutomatedTest.ScriptUrl)
	}

	result := _domain.RpcResp{}
	result.Success("")
	return result
}

func CheckoutCodes(task *commDomain.Build) {
	url := task.AutomatedTest.ScmAddress
	userName := task.AutomatedTest.ScmAccount
	password := task.AutomatedTest.ScmPassword

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

func DownloadCodes(task *commDomain.Build) {
	zipPath := task.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(task.AutomatedTest.ScriptUrl)
	_fileUtils.Download(task.AutomatedTest.ScriptUrl, zipPath)

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
