package commConst

type HostStatus string

const (
	HostActive  HostStatus = "active"
	HostOffline HostStatus = "offline"
)

func (e HostStatus) ToString() string {
	return string(e)
}

type VmStatus string

const (
	VmCreated       VmStatus = "created"
	VmLaunch        VmStatus = "launch"
	VmRunning       VmStatus = "running"
	VmActive        VmStatus = "active"
	VmBusy          VmStatus = "busy"
	VmDestroy       VmStatus = "destroy"
	VmFailToCreate  VmStatus = "fail_to_create"
	VmFailToDestroy VmStatus = "fail_to_destroy"
	VmUnknown       VmStatus = "unknown"
)

func (e VmStatus) ToString() string {
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

type BuildProgress string

const (
	ProgressCreated    BuildProgress = "created"
	ProgressLaunchVm   BuildProgress = "launch_vm"
	ProgressPending    BuildProgress = "pending"
	ProgressInProgress BuildProgress = "in_progress"
	ProgressTimeout    BuildProgress = "timeout"
	ProgressCompleted  BuildProgress = "completed"
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

type ServiceStatus string

const (
	ServiceOffline ServiceStatus = "offline"
	ServiceOnline  ServiceStatus = "online"
	ServiceActive  ServiceStatus = "active"
	ServiceBusy    ServiceStatus = "busy"
)

func (e ServiceStatus) ToString() string {
	return string(e)
}

type BuildType string

const (
	InterfaceScenario BuildType = "interface_scenario"
	InterfaceSet      BuildType = "interface_set"

	AutomatedTest BuildType = "automated_test"
)

func (e BuildType) ToString() string {
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
)

func (e OsLang) ToString() string {
	return string(e)
}

type AuthType string

const (
	None        AuthType = ""
	BasicAuth   AuthType = "basicAuth"
	BearerToken AuthType = "bearerToken"
	OAuth2      AuthType = "oauth2"
)

func (e AuthType) ToString() string {
	return string(e)
}

type OAuth2TypeGrantType string

const (
	AuthCode                         OAuth2TypeGrantType = "authCode"
	AuthCodeWithPKCE                 OAuth2TypeGrantType = "authCodeWithPKCE" // pkce: Proof Key for Code Exchange
	Implicit                         OAuth2TypeGrantType = "implicit"
	ResourceOwnerPasswordCredentials OAuth2TypeGrantType = "resourceOwnerPasswordCredentials"
	ClientCredentials                OAuth2TypeGrantType = "clientCredentials"
)

type OAuth2ClientAuthType string

const (
	AsBasicAuthHeader OAuth2ClientAuthType = "asBasicAuthHeader"
	CredentInBody     OAuth2ClientAuthType = "credentInBody"
)

func (e OAuth2ClientAuthType) ToString() string {
	return string(e)
}

type CodeChallengeMethod string

const (
	SHA256 CodeChallengeMethod = "sha256"
	Plain  CodeChallengeMethod = "plain"
)

func (e CodeChallengeMethod) ToString() string {
	return string(e)
}

type TestType string

const (
	Interface TestType = "interface"
	Case      TestType = "case"
	Scenario  TestType = "scenario"
)

func (e TestType) ToString() string {
	return string(e)
}

type PreviewType string

const (
	Auto PreviewType = "auto"
	Json PreviewType = "json"
	Html PreviewType = "html"
	Xml  PreviewType = "xml"
	Text PreviewType = "text"
)

func (e PreviewType) ToString() string {
	return string(e)
}

type ProcessorType string

const (
	Simple    ProcessorType = "simple"
	DataLoop  ProcessorType = "data_loop"
	Extractor ProcessorType = "extractor"
)

func (e ProcessorType) ToString() string {
	return string(e)
}

type ErrorAction string

const (
	ActionContinue        ErrorAction = "continue"
	ActionStartNextThread ErrorAction = "start_next_thread"
	ActionLoop            ErrorAction = "loop"
	ActionStopThread      ErrorAction = "stop_thread"
	ActionStopTest        ErrorAction = "stop_test"
	ActionStopTestNow     ErrorAction = "stop_test_now"
)

func (e ErrorAction) ToString() string {
	return string(e)
}

type DataSource string

const (
	ZenData DataSource = "zendata"
	CSV     DataSource = "csv"
	Excel   DataSource = "excel"
)

func (e DataSource) ToString() string {
	return string(e)
}

type ExtractorType string

const (
	Value       ExtractorType = "value"
	XPath       ExtractorType = "xpath"
	JSONPath    ExtractorType = "json_path"
	CssSelector ExtractorType = "sss_selector"
)

func (e ExtractorType) ToString() string {
	return string(e)
}

type ExtractorSource string

const (
	Body   ExtractorSource = "body"
	Header ExtractorSource = "header"
)

func (e ExtractorSource) ToString() string {
	return string(e)
}
