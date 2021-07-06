package consts

const (
	RegisterExpireTime  = 5  // min
	WaitToExecTime      = 60 // min
	WaitForVmLaunchTime = 60 // min
	WaitForResultTime   = 30 // min
	VmTimeout           = 20 // min

	MaxVmOnHost    = 3
	QueueRetryTime = 3

	DiskSizeWindows = 40 // G
	DiskSizeLinux   = 30
	DiskSizeDefault = 30

	ResDownDir    = "down"
	ResDriverDir  = "driver"
	ResDriverName = "driver"
)
