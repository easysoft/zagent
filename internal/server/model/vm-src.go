package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type VmSrc struct {
	BaseModel

	HostId int

	Name string

	OsCategory commConst.OsCategory
	OsType     commConst.OsType
	OsVersion  string
	OsLang     commConst.OsLang

	Status commConst.VmStatus
}

func (VmSrc) TableName() string {
	return "biz_vm_src"
}
