package testingService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	"path"
	"path/filepath"
	"strings"
)

type SeleniumService struct {
}

func NewSeleniumService() *SeleniumService {
	return &SeleniumService{}
}

func (s *SeleniumService) DownloadDriver(build *commDomain.Build) (err error) {
	// http://127.0.0.1:8085/down/driver/chrome/windows/92/chrome.exe

	fileName := consts.ResDriverName
	if _commonUtils.IsWin() {
		fileName += ".exe"
	}

	relatePath := path.Join(consts.ResDownDir, consts.ResDriverDir,
		build.BrowserType.ToString(), _commonUtils.GetOs(), build.BrowserVersion, fileName)

	url := agentConf.Inst.Server + relatePath
	filePath := filepath.Join(agentConf.Inst.WorkDir, strings.Replace(relatePath, "/", _const.PthSep, -1))

	if !_fileUtils.FileExist(filePath) {
		err = _fileUtils.Download(url, filePath)
	}

	if err == nil {
		build.SeleniumDriverPath = filePath
	}

	return
}
