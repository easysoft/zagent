package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
)

type VmTmpl struct {
	BaseModel

	HostId int `json:"hostId"`

	Name string `json:"name"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang"`

	Status consts.VmStatus `json:"status"`
}

func (VmTmpl) TableName() string {
	return "biz_vm_tmpl"
}
