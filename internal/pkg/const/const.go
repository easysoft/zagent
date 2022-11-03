package consts

const (
	AppName          = "zagent"
	AppNameAgent     = "agent"
	AppNameAgentHost = "host"
	AppNameAgentVm   = "vm"

	AgentPort      = 8086
	WebsockifyPort = 8087
	NoVncPath      = "/novnc"

	NatPortStart = 51602
	NatPortEnd   = 51799
	VncPortStart = 5901
	VncPortEnd   = 5999

	UploadMaxSize = 100000

	AgentCheckExecutionInterval = 15          // sec
	AgentCheckDownloadInterval  = 15          // sec
	WebCheckInterval            = 60          // sec
	DownloadImageTimeout        = 4 * 60 * 60 // sec
	ExportVmTimeout             = 20          // 1 * 60 * 60 // sec
	DownloadRetry               = 3

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
	SshServicePost        = 22

	ConfigVer = 1

	BrowserDriverDir = "browser_driver"

	BuildParamAppPath     = "${appPath}"
	BuildParamAppPackage  = "${appPackage}"
	BuildParamAppActivity = "${appActivity}"
	BuildParamAppiumPort  = "${appiumPort}"

	BuildParamSeleniumDriverPath = "${driverPath}"

	FolderKvm      = "kvm"
	FolderIso      = "iso"
	FolderDownload = "download"
	FolderBacking  = "backing"
	FolderImage    = "image"
	FolderToken    = "token"

	FolderNovnc      = "novnc"
	FolderWebsockify = "websockify"

	Localhost = "127.0.0.1"
)
