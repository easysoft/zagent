package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint
	Queue   `sql:"-", gorm:"foreignkey:QueueId"`

	BuildType commConst.BuildType
	VmId      uint

	Serial   string
	Priority int
	NodeIp   string
	NodePort int

	AppiumPort int

	StartTime    time.Time
	CompleteTime time.Time

	Progress commConst.BuildProgress
	Status   commConst.BuildStatus
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
