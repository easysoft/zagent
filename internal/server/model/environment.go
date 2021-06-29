package model

import commConst "github.com/easysoft/zagent/internal/comm/const"

type Environment struct {
	BaseModel

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`

	OsVersion string           `json:"osVersion"`
	OsLang    commConst.OsLang `json:"osLang"`

	TaskId string `json:"taskId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
