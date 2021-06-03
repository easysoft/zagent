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

type ServiceStatus string

const (
	ServiceOffline ServiceStatus = "offline"
	ServiceOnline  ServiceStatus = "online"
	ServiceActive  ServiceStatus = "active"
	ServiceBusy    ServiceStatus = "busy"
)

type BuildType string

const (
	InterfaceScenario BuildType = "interface_scenario"
	InterfaceSet      BuildType = "interface_set"

	AutomatedTest BuildType = "automated_test"
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

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Post   HttpMethod = "POST"
	Put    HttpMethod = "PUT"
	Delete HttpMethod = "DELETE"
)

type AuthType string

const (
	None        AuthType = ""
	BasicAuth   AuthType = "basicAuth"
	BearerToken AuthType = "bearerToken"
	OAuth2      AuthType = "oauth2"
)

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

type CodeChallengeMethod string

const (
	SHA256 CodeChallengeMethod = "sha256"
	Plain  CodeChallengeMethod = "plain"
)

type TestType string

const (
	Interface TestType = "interface"
	Case      TestType = "case"
	Scenario  TestType = "scenario"
)

type PreviewType string

const (
	Auto PreviewType = "auto"
	Json PreviewType = "json"
	Html PreviewType = "html"
	Xml  PreviewType = "xml"
	Text PreviewType = "text"
)

type ProcessorType string

const (
	// logic
	Simple ProcessorType = "simple"
	Once   TestType      = "once"
	Each   TestType      = "each"

	If        TestType      = "if"
	Loop      ProcessorType = "loop"
	Switch    TestType      = "switch"
	RandomOne TestType      = "randomOne"
	RandomAll TestType      = "randomAll"
	Variable  TestType      = "variable"

	// struct
	Include TestType = "include"

	// action
	CookieRetrieve TestType = "cookieRetrieve"
)

type ErrorAction string

const (
	ActionContinue        ErrorAction = "continue"
	ActionStartNextThread ErrorAction = "start_next_thread"
	ActionLoop            ErrorAction = "loop"
	ActionStopThread      ErrorAction = "stop_thread"
	ActionStopTest        ErrorAction = "stop_test"
	ActionStopTestNow     ErrorAction = "stop_test_now"
)
