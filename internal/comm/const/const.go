package consts

const (
	AgentCheckInterval    = 15 // sec
	WebCheckQueueInterval = 15 // sec

	DeviceRegisterExpireTime = 5 * 60 // sec

	WaitResPendingTimeout    = 15 * 60 // sec
	WaitForVmReadyTimeout    = 5 * 60  // sec
	WaitTestCompletedTimeout = 30 * 60 // sec
	VmLifecycleTimeout       = 60 * 60 // sec

	WaitAgentRunTaskTimeout = 30 * 60 // sec

	QueueRetryTime = 3

	DiskSizeWindows = 40 // G
	DiskSizeLinux   = 30
	DiskSizeDefault = 30

	ResDownDir    = "down"
	ResDriverDir  = "driver"
	ResDriverName = "driver"
)
