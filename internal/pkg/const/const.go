package consts

const (
	AppName          = "zagent"
	AppNameAgent     = "agent"
	AppNameAgentHost = "host"
	AppNameAgentVm   = "vm"

	UploadMaxSize = 100000

	AgentCheckExecutionInterval = 15          // sec
	AgentCheckDownloadInterval  = 15          // sec
	WebCheckInterval            = 60          // sec
	DownloadImageTimeout        = 4 * 60 * 60 // sec
	ExportVmTimeout             = 1 * 60 * 60 // sec
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

	SshServicePost       = 22
	AgentHostServicePort = 55001
	WebsockifyPort       = 55002
	NatPortStart         = 55100
	NatPortEnd           = 55199

	AgentVmServicePort = 55201
	ZtfServicePost     = 56202
	ZdServicePost      = 56203

	FolderKvm      = "kvm"
	FolderIso      = "iso"
	FolderDownload = "download"
	FolderBacking  = "backing"
	FolderImage    = "image"
	FolderToken    = "token"

	FolderNovnc      = "novnc"
	FolderWebsockify = "websockify"

	FolderZtf = "ztf"
	FolderZd  = "zd"

	Localhost = "127.0.0.1"

	QiNiuUrl           = "https://dl.cnezsoft.com/"
	VersionDownloadUrl = QiNiuUrl + "%s/version.txt"
	PackageDownloadUrl = QiNiuUrl + "%s/%s/%s/%s.zip"

	ALIYUN_ECS_URL     = "ecs-%s.aliyuncs.com"
	ALIYUN_ECS_URL_VNC = "https://g.alicdn.com/aliyun/ecs-console-vnc2/0.0.8/index.html" +
		"?vncUrl=%s&instanceId=%s&isWindows=%t&password=%s"
	ALIYUN_ECI_URL = "eci.aliyuncs.com"
	ALIYUN_VPC_URL = "vpc.aliyuncs.com"

	HuaweiCloudUrlJobCreate  = "https://cci.%s.myhuaweicloud.cn/apis/batch/v1/namespaces/%s/jobs"
	HuaweiCloudUrlJobDestroy = "https://cci.%s.myhuaweicloud.cn/apis/batch/v1/namespaces/%s/jobs/%s"
)
