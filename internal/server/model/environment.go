package model

import consts "github.com/easysoft/zv/internal/pkg/const"

type Environment struct {
	BaseModel

	OsCategory consts.OsCategory `json:"osCategory" example:"windows"`
	OsType     consts.OsType     `json:"osType" example:"win10"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang" example:"zh_cn"`

	ImageName string `json:"imageName"`
	ImageSrc  string `json:"imageSrc"`

	TaskId string `json:"taskId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
