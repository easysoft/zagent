package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Host struct {
	BaseModel

	Name string `json:"name"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsLang     commConst.OsLang     `json:"osLang"`

	OsVersion string `json:"osVersion"`
	OsBuild   string `json:"osBuild"`
	OsBits    string `json:"osBits"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir"`

	SshPort int                  `json:"sshPort"`
	VncPort int                  `json:"vncPort"`
	Status  commConst.HostStatus `json:"status"`

	TaskCount        int        `json:"taskCount"`
	LastRegisterDate *time.Time `json:"lastRegisterDate"`
}

func NewHost() Host {
	host := Host{}
	return host
}

func (Host) TableName() string {
	return "biz_host"
}
