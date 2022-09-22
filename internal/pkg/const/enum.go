package consts

type HostCapability string

const (
	PlatformVm     HostCapability = "vm"
	PlatformDocker HostCapability = "docker"

	PlatformNative HostCapability = "native"
	PlatformCloud  HostCapability = "cloud"
)

func (e HostCapability) ToString() string {
	return string(e)
}

type HostVendor string

const (
	HostVendorVirtualBox  HostVendor = "virtualbox"
	HostVendorVmWare      HostVendor = "vmware"
	HostVendorHuaweiCloud HostVendor = "huaweicloud"
	HostVendorAliyun      HostVendor = "aliyun"

	HostVendorPve       HostVendor = "pve"
	HostVendorPortainer HostVendor = "portainer"
)

func (e HostVendor) ToString() string {
	return string(e)
}

type HostStatus string

const (
	HostOnline  HostStatus = "online"
	HostOffline HostStatus = "offline"
	HostBusy    HostStatus = "busy" // report by agent on host
)

func (e HostStatus) ToString() string {
	return string(e)
}

type VmStatus string

const (
	VmCreated    VmStatus = "created"        // set first time created
	VmLaunch     VmStatus = "launch"         // set after success to call vm creating remotely
	VmFailCreate VmStatus = "vm_fail_create" // set after fail to call vm creating remotely

	VmRunning VmStatus = "running" // report by agent on host
	VmShutOff VmStatus = "shutoff" // report by agent on host

	VmBusy  VmStatus = "busy"  // report by agent in vm
	VmReady VmStatus = "ready" // report by agent in vm

	VmUnknown     VmStatus = "unknown"      // report by agent on host, not running, destroy and shutoff
	VmDestroy     VmStatus = "destroy"      // final status
	VmDestroyFail VmStatus = "destroy_fail" // set after fail to call vm destroy remotely
)

func (e VmStatus) ToString() string {
	return string(e)
}

type BuildProgress string

const (
	// start group
	ProgressCreated BuildProgress = "created"

	// res group
	ProgressResPending     BuildProgress = "res_pending"
	ProgressResLaunched    BuildProgress = "res_launched"
	ProgressResReady       BuildProgress = "res_ready"
	ProgressResFailed      BuildProgress = "res_failed"
	ProgressResDestroy     BuildProgress = "res_destroy"
	ProgressResFailDestroy BuildProgress = "res_fail_destroy"

	// exec group
	ProgressRunning BuildProgress = "running"
	ProgressRunFail BuildProgress = "run_fail"

	// end group
	ProgressCompleted BuildProgress = "completed"
	ProgressTimeout   BuildProgress = "timeout"
	ProgressTerminal  BuildProgress = "terminal"
	ProgressCancel    BuildProgress = "cancel"
)

func (e BuildProgress) ToString() string {
	return string(e)
}

type BuildStatus string

const (
	StatusCreated BuildStatus = "created"
	StatusPass    BuildStatus = "pass"
	StatusFail    BuildStatus = "fail"
)

func (e BuildStatus) ToString() string {
	return string(e)
}

type BuildType string

const (
	ZtfTest       BuildType = "ztf"
	SeleniumTest  BuildType = "selenium"
	AppiumTest    BuildType = "appium"
	UnitTest      BuildType = "unittest"
	InterfaceTest BuildType = "interface"
)

func (e BuildType) ToString() string {
	return string(e)
}

type DeviceStatus string

const (
	DeviceOff     DeviceStatus = "off"
	DeviceOn      DeviceStatus = "on"
	DeviceActive  DeviceStatus = "active"
	DeviceBusy    DeviceStatus = "busy"
	DeviceUnknown DeviceStatus = "unknown"
)

func (e DeviceStatus) ToString() string {
	return string(e)
}

type ServiceStatus string

const (
	ServiceOffline ServiceStatus = "offline"
	ServiceOnline  ServiceStatus = "online"
	ServiceReady   ServiceStatus = "ready"
	ServiceBusy    ServiceStatus = "busy"
)

func (e ServiceStatus) ToString() string {
	return string(e)
}

type OsDevice string

const (
	Android OsDevice = "android"
	Ios     OsDevice = "ios"
	Harmony OsDevice = "harmony"
)

func (e OsDevice) ToString() string {
	return string(e)
}

type OsCategory string

const (
	Windows OsCategory = "windows"
	Linux   OsCategory = "linux"
	Mac     OsCategory = "mac"
)

func (e OsCategory) ToString() string {
	return string(e)
}

type OsType string

const (
	Win10 OsType = "win10"
	Win7  OsType = "win7"
	WinXP OsType = "winxp"

	Ubuntu OsType = "ubuntu"
	CentOS OsType = "centos"
	Debian OsType = "debian"
)

func (e OsType) ToString() string {
	return string(e)
}

type OsLang string

const (
	EN_US OsLang = "en_us"
	ZH_CN OsLang = "zh_cn"
	ZH_TW OsLang = "zh_tw"
)

func (e OsLang) ToString() string {
	return string(e)
}

type EntityType string

const (
	Task  EntityType = "task"
	Queue EntityType = "queue"
	Build EntityType = "build"
	Vm    EntityType = "vm"
)

func (e EntityType) ToString() string {
	return string(e)
}

type NatForwardType string

const (
	Http   NatForwardType = "http"
	Stream NatForwardType = "stream"
	All    NatForwardType = "*"
)

func (e NatForwardType) ToString() string {
	return string(e)
}

type BrowserType string

const (
	Chrome  BrowserType = "chrome"
	Firefox BrowserType = "firefox"
	Edge    BrowserType = "edge"
	IE      BrowserType = "ie"
)

func (e BrowserType) ToString() string {
	return string(e)
}

type RunMode string

const (
	RunModeHost RunMode = "host"
	RunModeVm   RunMode = "vm"

	RunModeMachine RunMode = "machine"
	RunModeAndroid RunMode = "android"
	RunModeIos     RunMode = "ios"
)

func (e RunMode) ToString() string {
	return string(e)
}
