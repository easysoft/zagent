package agentTestingService

import (
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	"path/filepath"
)

type SeleniumService struct {
}

func NewSeleniumService() *SeleniumService {
	return &SeleniumService{}
}

func (s *SeleniumService) DownloadDriver(build *commDomain.Build) (err error) {
	// https://dl.cnezsoft.com/driver/chrome/windows/93/driver.exe

	fileName := consts.ResDriverName
	if _commonUtils.IsWin() {
		fileName += ".exe"
	}

	filePath := filepath.Join(agentConf.Inst.WorkDir, consts.ResDownDir, consts.ResDriverDir,
		build.BrowserType.ToString(), _commonUtils.GetOs(), build.BrowserVersion, fileName)

	url := fmt.Sprintf("%s%s/%s/%s/driver",
		consts.DriverDownloadUrl, build.BrowserType.ToString(), _commonUtils.GetOs(), build.BrowserVersion)
	if _commonUtils.IsWin() {
		url += ".exe"
	}

	if !_fileUtils.FileExist(filePath) {
		err = _fileUtils.Download(url, filePath)
	}

	if err == nil {
		build.SeleniumDriverPath = filePath
	}

	return
}
