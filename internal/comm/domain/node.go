package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type Node struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Ip      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	WorkDir string `json:"workDir" yaml:"workDir"`

	Status commConst.ServiceStatus `json:"status" yaml:"status"`

	OsPlatform commConst.OsCategory `json:"osPlatform" yaml:"osPlatform"`
	OsType     commConst.OsType     `json:"osType" yaml:"osType"`
	SysLang    commConst.OsLang     `json:"sysLang" yaml:"sysLang"`
}
