package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
)

type History struct {
	BaseModel

	OwnerID   uint
	OwnerType string

	Progress string `json:"progress"`
	Status   string `json:"status"`
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
