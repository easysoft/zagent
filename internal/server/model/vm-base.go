package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type VmBase struct {
	BaseModel

	Name string
	Path string
	Size int

	OsCategory commConst.OsCategory
	OsType     commConst.OsType
	OsLang     commConst.OsLang

	OsVersion string
	OsBuild   string
	OsBits    string

	ResolutionHeight  int
	ResolutionWidth   int
	SuggestDiskSize   int
	SuggestMemorySize int

	SysIsoId    uint
	DriverIsoId uint
}

func (VmBase) TableName() string {
	return "biz_vm_base"
}
