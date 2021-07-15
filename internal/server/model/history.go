package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
)

type History struct {
	BaseModel

	OwnerID   uint   `json:"ownerID"`
	OwnerType string `json:"ownerType"`

	Progress string `json:"progress"`
	Status   string `json:"status"`

	QueueId uint   `json:"queueId" gorm:"-"`
	NodeIp  string `json:"nodeIp" gorm:"-"`
	VncPort int    `json:"vncPort" gorm:"-"`
}

func NewHistoryPo(tp consts.EntityType, id uint, progress consts.BuildProgress, status string) History {
	po := History{
		OwnerType: tp.ToString(),
		OwnerID:   id,
		Progress:  progress.ToString(),
		Status:    status,
	}

	return po
}

func (History) TableName() string {
	return "biz_history"
}
