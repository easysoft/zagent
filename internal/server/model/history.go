package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
)

type History struct {
	BaseModel

	TaskId  uint `json:"taskId"`
	QueueId uint `json:"queueId"`
	BuildId uint `json:"buildId"`
	VmId    uint `json:"vmId"`

	Type consts.EntityType `json:"type"`

	Progress string `json:"progress"`
	Status   string `json:"status"`
}

func NewHistoryPo(tp consts.EntityType, id uint, progress consts.BuildProgress, status string) History {
	po := History{
		Progress: progress.ToString(),
		Status:   status,
	}

	if tp == consts.Task {
		po.TaskId = id
	} else if tp == consts.Queue {
		po.QueueId = id
	} else if tp == consts.Build {
		po.BuildId = id
	} else if tp == consts.Vm {
		po.VmId = id
	}

	return po
}

func (History) TableName() string {
	return "biz_history"
}
