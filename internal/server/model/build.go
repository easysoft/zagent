package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint `json:"queueId"`
	VmId    uint `json:"vmId"`

	BuildType consts.BuildType `json:"buildType"`

	Serial   string `json:"serial"`
	Priority int    `json:"priority"`
	NodeIp   string `json:"nodeIp"`
	NodePort int    `json:"nodePort"`

	AppiumPort     int                `json:"appiumPort"`
	BrowserType    consts.BrowserType `json:"browserType"`
	BrowserVersion string             `json:"browserVersion"`

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"ScmAddress"`
	ScmAccount  string `json:"ScmAccount"`
	ScmPassword string `json:"ScmPassword"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	StartTime    *time.Time `json:"startTime"`
	CompleteTime *time.Time `json:"completeTime"`

	ResultPath string `json:"resultPath"`
	ResultMsg  string `json:"resultMsg"`

	Progress consts.BuildProgress `json:"progress"`
	Status   consts.BuildStatus   `json:"status"`
}

func NewSeleniumBuildPo(queue Queue, vm Vm) Build {
	build := Build{
		QueueId:   queue.ID,
		VmId:      vm.ID,
		BuildType: queue.BuildType,
		Priority:  queue.Priority,
		NodeIp:    vm.NodeIp,
		NodePort:  vm.NodePort,

		BrowserType:    queue.BrowserType,
		BrowserVersion: queue.BrowserVersion,

		ScriptUrl:   queue.ScriptUrl,
		ScmAddress:  queue.ScmAddress,
		ScmAccount:  queue.ScmAccount,
		ScmPassword: queue.ScmPassword,

		EnvVars:       queue.EnvVars,
		BuildCommands: queue.BuildCommands,
		ResultFiles:   queue.ResultFiles,

		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,
	}

	return build
}

func NewAppiumBuildPo(queue Queue, dev Device) Build {
	build := Build{
		QueueId:   queue.ID,
		BuildType: queue.BuildType,
		Priority:  queue.Priority,
		NodeIp:    dev.NodeIp,
		NodePort:  dev.NodePort,

		AppUrl:      queue.AppUrl,
		ScriptUrl:   queue.ScriptUrl,
		ScmAddress:  queue.ScmAddress,
		ScmAccount:  queue.ScmAccount,
		ScmPassword: queue.ScmPassword,

		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,
	}

	return build
}

func NewBuildTo(build Build) domain.Build {
	toValue := domain.Build{
		ID:        build.ID,
		QueueId:   build.QueueId,
		VmId:      build.VmId,
		BuildType: build.BuildType,
		Serial:    build.Serial,
		Priority:  build.Priority,
		NodeIp:    build.NodeIp,
		NodePort:  build.NodePort,

		BrowserType:    build.BrowserType,
		BrowserVersion: build.BrowserVersion,

		AppUrl:      build.AppUrl,
		ScriptUrl:   build.ScriptUrl,
		ScmAddress:  build.ScmAddress,
		ScmAccount:  build.ScmAccount,
		ScmPassword: build.ScmPassword,

		EnvVars:       build.EnvVars,
		BuildCommands: build.BuildCommands,
		ResultFiles:   build.ResultFiles,

		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,
	}

	return toValue
}

func (Build) TableName() string {
	return "biz_build"
}
