package model

type Environment struct {
	BaseModel

	OsCategory string `json:"osCategory"`
	OsType     string `json:"osType"`

	OsVersion string `json:"osVersion"`
	OsLang    string `json:"osLang"`

	TaskId string `json:"taskId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
