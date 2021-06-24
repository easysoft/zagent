package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type Iso struct {
	BaseModel

	Name string
	Path string
	Size int

	OsPlatform commConst.OsCategory
	OsType     commConst.OsType
	OsLang     commConst.OsLang

	OsVersion string
	OsBuild   string
	OsBits    string

	ResolutionHeight  int
	ResolutionWidth   int
	suggestDiskSize   int
	suggestMemorySize int
}

func (Iso) TableName() string {
	return "biz_iso"
}
