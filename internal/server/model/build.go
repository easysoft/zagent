package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint `json:"queueId"`
	Queue   `json:"queue" sql:"-", gorm:"foreignkey:QueueId"`

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

func NewBuild() Build {
	build := Build{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}

	return build
}

func NewBuildDetail(queueId uint, vmId uint, buildType commConst.BuildType,
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
