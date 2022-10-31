package domain

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type Build struct {
	Version float64 `json:"version,omitempty"`
	QueueId uint    `json:"queueId,omitempty"`
	VmId    uint    `json:"vmId,omitempty"`
	ID      uint    `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Desc    string  `json:"desc,omitempty"`

	BuildType  consts.BuildType  `json:"buildType,omitempty"`
	OsCateGory consts.OsCategory `json:"osCateGory,omitempty"`
	OsType     consts.OsType     `json:"osType,omitempty"`
	OsVersion  string            `json:"osVersion,omitempty"`
	OsLang     consts.OsLang     `json:"osLang,omitempty"`

	WorkDir    string `json:"workDir,omitempty"`
	ProjectDir string `json:"projectDir,omitempty"`

	Serial   string `json:"serial,omitempty"`
	Priority int    `json:"priority,omitempty"`
	NodeIp   string `json:"nodeIp,omitempty"`
	NodePort int    `json:"nodePort,omitempty"`
	DeviceIp string `json:"deviceIp,omitempty"`

	AppiumPort         int                `json:"appiumPort,omitempty"`
	BrowserType        consts.BrowserType `json:"browserType,omitempty"`
	BrowserVersion     string             `json:"browserVersion,omitempty"`
	SeleniumDriverPath string             `json:"seleniumDriverPath,omitempty"`

	AppPath     string `json:"appPath,omitempty"`
	ScriptUrl   string `json:"scriptUrl,omitempty"`
	ScmAddress  string `json:"scmAddress,omitempty"`
	ScmAccount  string `json:"scmAccount,omitempty"`
	ScmPassword string `json:"scmPassword,omitempty"`

	AppUrl          string `json:"appUrl,omitempty"`
	BuildCommands   string `json:"buildCommands,omitempty"`
	EnvVars         string `json:"envVars,omitempty"`
	ResultFiles     string `json:"resultFiles,omitempty"`
	KeepResultFiles bool   `json:"keepResultFiles,omitempty"`
	ResultPath      string `json:"resultPath,omitempty"`
	ResultMsg       string `json:"resultMsg,omitempty"`

	StartTime    *time.Time `json:"startTime,omitempty"`
	CompleteTime *time.Time `json:"completeTime,omitempty"`

	Progress consts.BuildProgress `json:"progress"`
	Status   consts.BuildStatus   `json:"status"`
}
