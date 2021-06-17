package agentConst

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"path/filepath"
)

var (
	AppName    = "zagent"
	ConfigVer  = 1
	ConfigFile = filepath.Join("conf", AppName+".yaml")

	EnRes = filepath.Join("res", "messages_en.json")
	ZhRes = filepath.Join("res", "messages_zh.json")

	BrowserDriverDir = "browser_driver"
	LogDir           = fmt.Sprintf("log%s", _const.PthSep)

	BuildParamAppPath     = "${appPath}"
	BuildParamAppPackage  = "${appPackage}"
	BuildParamAppActivity = "${appActivity}"
	BuildParamAppiumPort  = "${appiumPort}"

	BuildParamSeleniumDriverPath = "${driverPath}"

	ControlActions = []string{"start", "stop", "restart", "install", "uninstall"}

	FolderKvm   = "kvm"
	FolderIso   = "iso"
	FolderImage = "image"
	FolderDef   = "def"
	FolderTempl = "templ"
)
