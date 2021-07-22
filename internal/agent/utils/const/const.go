package agentConst

import (
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"path/filepath"
)

var (
	ConfigVer  = 1
	ConfigFile = filepath.Join("conf", consts.AppNameAgent+".yaml")

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

	FolderKvm     = "kvm"
	FolderIso     = "iso"
	FolderBacking = "base"
	FolderImage   = "image"
	//FolderDef   = "def"
	//FolderTempl = "templ"
)
