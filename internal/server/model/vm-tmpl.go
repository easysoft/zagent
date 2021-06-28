package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type VmTmpl struct {
	BaseModel

	HostId int `json:"hostId"`

	Name string `json:"name"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsVersion  string               `json:"osVersion"`
	OsLang     commConst.OsLang     `json:"osLang"`

	Status commConst.VmStatus `json:"status"`
}

func (VmTmpl) TableName() string {
	return "biz_vm_tmpl"
}
