package domain

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Host struct {
	Name string `json:"name"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsLang     consts.OsLang     `json:"osLang"`

	OsVersion string `json:"osVersion"`
	OsBuild   string `json:"osBuild"`
	OsBits    string `json:"osBits"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir"`

	SshPort    int               `json:"sshPort"`
	VncAddress string            `json:"vncAddress"`
	Status     consts.HostStatus `json:"status"`

	TaskCount        int        `json:"taskCount"`
	LastRegisterDate *time.Time `json:"lastRegisterDate"`

	Vms []Vm `json:"vms" gorm:"-"`
}
