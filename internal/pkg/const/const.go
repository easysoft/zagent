package consts

import (
	"fmt"
	_const "github.com/easysoft/zv/pkg/const"
	"path/filepath"
)

const (
	AppName          = "zagent"
	AppNameServer    = "server"
	AppNameAgent     = "agent"
	AppNameAgentHost = "host"
	AppNameAgentVm   = "vm"

	ServerPort = 8085
	AgentPort  = 8086

	WebsockifyPort = 51600
	NoVncPort      = 51601
	NatPortStart   = 51602
	NatPortEnd     = 51799
	VncPortStart   = 51800
	VncPortEnd     = 51999

	AgentCheckInterval = 5  // sec
	WebCheckInterval   = 60 // sec

	DeviceRegisterExpireTime = 5 * 60 // sec

	// Must > WaitResPendingTimeout
	WaitResPendingTimeout   = 60 * 60 // sec
	WaitResReadyTimeout     = 60 * 60 // sec
	WaitRunCompletedTimeout = 60 * 60 // sec
	WaitVmLifecycleTimeout  = 60 * 60 // sec
	WaitAgentRunTaskTimeout = 30 * 60 // sec

	QueueRetryTime = 3

	DiskSizeWindows = 40 // G
	DiskSizeLinux   = 30
	DiskSizeDefault = 30

	DriverDownloadUrl = "https://dl.cnezsoft.com/driver/"
	ResDownDir        = "down"
	ResDriverDir      = "driver"
	ResDriverName     = "driver"

	KvmHostIpInNatNetwork = "192.168.122.1"
	AgentServicePost      = 8086

	ConfigVer = 1

	BrowserDriverDir = "browser_driver"

	BuildParamAppPath     = "${appPath}"
	BuildParamAppPackage  = "${appPackage}"
	BuildParamAppActivity = "${appActivity}"
	BuildParamAppiumPort  = "${appiumPort}"

	BuildParamSeleniumDriverPath = "${driverPath}"

	FolderKvm     = "kvm"
	FolderIso     = "iso"
	FolderBacking = "backing"
	FolderImage   = "image"
	FolderToken   = "token"
)

var (
	ConfigFile     = filepath.Join("conf", AppNameAgent+".yaml")
	LogDir         = fmt.Sprintf("log%s", _const.PthSep)
	ControlActions = []string{"start", "stop", "restart", "install", "uninstall"}
)
