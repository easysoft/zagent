package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type Iso struct {
	BaseModel

	Name string `json:"name"`
	Path string `json:"path"`
	Size int    `json:"size"`

	OsPlatform commConst.OsCategory `json:"osPlatform"`
	OsType     commConst.OsType     `json:"osType"`
	OsLang     commConst.OsLang     `json:"osLang"`

	OsVersion string `json:"osVersion"`
	OsBuild   string `json:"osBuild"`
	OsBits    string `json:"osBits"`

	ResolutionHeight  int `json:"resolutionHeight"`
	ResolutionWidth   int `json:"resolutionWidth"`
	SuggestDiskSize   int `json:"suggestDiskSize"`
	SuggestMemorySize int `json:"suggestMemorySize"`
}

func (Iso) TableName() string {
	return "biz_iso"
}
