package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Host struct {
	BaseModel

	Name string

	OsCategory commConst.OsCategory
	OsType     commConst.OsType
	OsLang     commConst.OsLang

	OsVersion string
	OsBuild   string
	OsBits    string

	Ip      string
	Port    int
	WorkDir string

	SshPort int
	VncPort int
	Status  commConst.HostStatus

	taskCount        int
	LastRegisterDate time.Time
}

func NewHost() Host {
	host := Host{}
	return host
}

func (Host) TableName() string {
	return "biz_host"
}
