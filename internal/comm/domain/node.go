package domain

import (
	"github.com/easysoft/zv/internal/comm/const"
	"time"
)

type HostNode struct {
	Node
	Status consts.HostStatus `json:"status" yaml:"status"`
	Vms    []Vm
}

type Node struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Ip      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	WorkDir string `json:"workDir" yaml:"workDir"`

	OsPlatform consts.OsCategory `json:"osPlatform" yaml:"osPlatform"`
	OsType     consts.OsType     `json:"osType" yaml:"osType"`
	SysLang    consts.OsLang     `json:"sysLang" yaml:"sysLang"`

	SshPort    int    `json:"sshPort" yaml:"sshPort"`
	VncAddress string `json:"vncAddress" yaml:"vncAddress"`

	Secret string `json:"secret" yaml:"secret"`

	LastRegisterDate time.Time `json:"lastRegisterDate" yaml:"lastRegisterDate"`
}
