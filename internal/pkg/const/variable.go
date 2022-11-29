package consts

import (
	"fmt"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
	"path/filepath"
	"sync"
	"time"

	_const "github.com/easysoft/zagent/pkg/const"
)

var ()

var (
	ConfigFile     = filepath.Join("conf", AppNameAgent+".yaml")
	LogDir         = fmt.Sprintf("log%s", _const.PthSep)
	ControlActions = []string{"start", "stop", "restart", "install", "uninstall"}

	DownloadDir   = ""
	NovncDir      = ""
	WebsockifyDir = ""
	WorkDir       = ""

	AuthSecret  = ""
	AuthToken   = ""
	ExpiredDate = time.Now()

	Verbose          = false
	ExistVncPortMap  = sync.Map{}
	ExistHttpPortMap = sync.Map{}
	ExistSshPortMap  = sync.Map{}

	PrintLog = true

	ZtfUuid = _stringUtils.Uuid()
	ZdUuid  = _stringUtils.Uuid()
)
