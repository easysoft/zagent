package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint `json:"queueId"`
	Queue   `json:"queue" sql:"-" gorm:"foreignkey:QueueId"`

	BuildType consts.BuildType `json:"buildType"`
	VmId      uint             `json:"vmId"`

	Serial   string `json:"serial"`
	Priority int    `json:"priority"`
	NodeIp   string `json:"nodeIp"`
	NodePort int    `json:"nodePort"`

	AppiumPort int `json:"appiumPort"`

	StartTime    *time.Time `json:"startTime"`
	CompleteTime *time.Time `json:"completeTime"`

	Progress consts.BuildProgress `json:"progress"`
	Status   consts.BuildStatus   `json:"status"`
}

func NewBuild(queueId uint, vmId uint, buildType consts.BuildType,
	serial string, priority int, nodeIp string, nodePort int) Build {
	build := Build{
		QueueId:   queueId,
		VmId:      vmId,
		BuildType: buildType,
		Serial:    serial,
		Priority:  priority,
		NodeIp:    nodeIp,
		NodePort:  nodePort,

		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,
	}

	return build
}

func NewBuildTo(build Build) domain.Build {
	toValue := domain.Build{
		ID:        build.ID,
		QueueId:   build.QueueId,
		BuildType: build.BuildType,
		Serial:    build.Serial,
		Priority:  build.Priority,
		NodeIp:    build.NodeIp,
		NodePort:  build.NodePort,

		AppUrl: build.AppUrl,

		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,

		ScriptUrl:   build.ScriptUrl,
		ScmAddress:  build.ScmAddress,
		ScmAccount:  build.ScmAccount,
		ScmPassword: build.ScmPassword,

		BuildCommands:   build.BuildCommands,
		EnvVars:         build.EnvVars,
		ResultFiles:     build.ResultFiles,
		KeepResultFiles: build.KeepResultFiles,
	}

	return toValue
}
