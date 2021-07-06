package consts

const (
	RegisterExpireTime       = 5  // min
	WaitResPendingTimeout    = 15 // min
	WaitForVmReadyTimeout    = 5  // min
	WaitTestCompletedTimeout = 30 // min
	VmLifecycleTimeout       = 20 // min

	QueueRetryTime = 3

	DiskSizeWindows = 40 // G
	DiskSizeLinux   = 30
	DiskSizeDefault = 30

	ResDownDir    = "down"
	ResDriverDir  = "driver"
	ResDriverName = "driver"
)
