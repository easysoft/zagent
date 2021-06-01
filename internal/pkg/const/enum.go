package _const

type ResultCode int

const (
	ResultSuccess ResultCode = 1
	ResultFail    ResultCode = 0
)

func (c ResultCode) Int() int {
	return int(c)
}

type BuildProgress string

const (
	ProgressCreated    BuildProgress = "created"
	ProgressLaunchVm   BuildProgress = "launch_vm"
	ProgressPending    BuildProgress = "pending"
	ProgressInProgress BuildProgress = "in_progress"
	ProgressTimeout    BuildProgress = "timeout"
	ProgressCompleted  BuildProgress = "completed"
)

type BuildStatus string

const (
	StatusCreated BuildStatus = "created"
	StatusPass    BuildStatus = "pass"
	StatusFail    BuildStatus = "fail"
)

type NodeStatus string

const (
	NodeActive        NodeStatus = "active"
	NodeBusy          NodeStatus = "busy"
)

type HostStatus string

type ServiceStatus string

const (
	ServiceOff    ServiceStatus = "off"
	ServiceOn     ServiceStatus = "on"
	ServiceActive ServiceStatus = "active"
	ServiceBusy   ServiceStatus = "busy"
)

type BuildType string

const (
	InterfaceTest     BuildType = "interface_test"
	UnitTest     BuildType = "unit_test"
	ZtfTest     BuildType = "ztf_test"
)

type OsPlatform string

const (
	OsWindows OsPlatform = "windows"
	OsLinux   OsPlatform = "linux"
	OsMac     OsPlatform = "mac"
)

type OsType string

const (
	Win10 OsType = "win10"
	Win7  OsType = "win7"
	WinXP OsType = "winxp"

	Ubuntu OsType = "ubuntu"
	CentOS OsType = "centos"
	Debian OsType = "debian"

	Mac OsType = "mac"
)

type SysLang string

const (
	EN_US SysLang = "en_us"
	ZH_CN SysLang = "zh_cn"
)
