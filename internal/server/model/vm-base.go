package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type VmBase struct {
	BaseModel

	Name string `json:"name"`
	Path string `json:"path"`
	Size int    `json:"size"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsLang     commConst.OsLang     `json:"osLang"`

	OsVersion string `json:"osVersion"`
	OsBuild   string `json:"osBuild"`
	OsBits    string `json:"osBits"`

	ResolutionHeight  int `json:"resolutionHeight"`
	ResolutionWidth   int `json:"resolutionWidth"`
	SuggestDiskSize   int `json:"suggestDiskSize"`
	SuggestMemorySize int `json:"suggestMemorySize"`

	SysIsoId    uint `json:"sysIsoId"`
	DriverIsoId uint `json:"driverIsoId"`
}

func (VmBase) TableName() string {
	return "biz_vm_base"
}
