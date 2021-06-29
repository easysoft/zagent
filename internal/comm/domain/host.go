package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Host struct {
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

	Vms []Vm
}
