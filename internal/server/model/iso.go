package model

import (
	"github.com/easysoft/zv/internal/comm/const"
)

type Iso struct {
	BaseModel

	Name string `json:"name"`
	Path string `json:"path"`
	Size int    `json:"size"`

	OsPlatform consts.OsCategory `json:"osPlatform"`
	OsType     consts.OsType     `json:"osType"`
	OsLang     consts.OsLang     `json:"osLang"`

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
