package model

import (
	"github.com/easysoft/zv/internal/comm/const"
)

type History struct {
	BaseModel

	OwnerID   uint   `json:"ownerID"`
	OwnerType string `json:"ownerType"`

	Progress string `json:"progress"`
	Status   string `json:"status"`

	QueueId    uint   `json:"queueId"`
	NodeIp     string `json:"nodeIp" gorm:"-"`
	VncAddress string `json:"vncAddress" gorm:"-"`
}

func NewHistoryPo(tp consts.EntityType, id, queueId uint, progress consts.BuildProgress, status string) History {
	po := History{
		OwnerType: tp.ToString(),
		OwnerID:   id,
		QueueId:   queueId,
		Progress:  progress.ToString(),
		Status:    status,
	}

	return po
}

func (History) TableName() string {
	return "biz_history"
}
