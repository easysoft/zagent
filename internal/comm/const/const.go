package consts

const (
	AppName          = "zv"
	AppNameServer    = "server"
	AppNameAgent     = "agent"
	AppNameAgentHost = "host"
	AppNameAgentVm   = "vm"

	ServerPort = 8085
	AgentPort  = 8086

	AgentCheckInterval = 15 // sec
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
)
