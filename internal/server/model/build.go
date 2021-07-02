package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint `json:"queueId"`
	Queue   `json:"queue" sql:"-" gorm:"foreignkey:QueueId"`

	BuildType commConst.BuildType `json:"buildType"`
	VmId      uint                `json:"vmId"`

	Serial   string `json:"serial"`
	Priority int    `json:"priority"`
	NodeIp   string `json:"nodeIp"`
	NodePort int    `json:"nodePort"`

	AppiumPort int `json:"appiumPort"`

	StartTime    *time.Time `json:"startTime"`
	CompleteTime *time.Time `json:"completeTime"`

	Progress commConst.BuildProgress `json:"progress"`
	Status   commConst.BuildStatus   `json:"status"`
}

func NewBuild(queueId uint, vmId uint, buildType commConst.BuildType,
	serial string, priority int, nodeIp string, nodePort int) Build {
	build := Build{
		QueueId:   queueId,
		VmId:      vmId,
		BuildType: buildType,
		Serial:    serial,
		Priority:  priority,
		NodeIp:    nodeIp,
		NodePort:  nodePort,

		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}

	return build
}

func NewBuildTo(build Build) commDomain.Build {
	toValue := commDomain.Build{
		ID:        build.ID,
		QueueId:   build.QueueId,
		BuildType: build.BuildType,
		Serial:    build.Serial,
		Priority:  build.Priority,
		NodeIp:    build.NodeIp,
		NodePort:  build.NodePort,

		AppUrl: build.AppUrl,

		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,

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
