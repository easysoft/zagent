package model

import "github.com/easysoft/zagent/internal/comm/const"

type Environment struct {
	BaseModel

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang"`

	ImageName string `json:"imageName"`
	ImageSrc  string `json:"imageSrc"`

	TaskId string `json:"taskId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
