package model

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
)

type Container struct {
	BaseModel

	Name      string `json:"name"`
	Desc      string `json:"desc"`
	ImageName string `json:"imageName"`

	HostId   uint   `json:"hostId"`
	HostName string `json:"hostName"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang"`

	NodeIp   string `json:"nodeIp"`
	NodePort int    `json:"nodePort"`

	Histories []History `json:"histories" gorm:"polymorphic:Owner;polymorphicValue:container"`
}

func (Container) TableName() string {
	return "biz_container"
}
