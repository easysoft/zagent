package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type Node struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Ip      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	WorkDir string `json:"workDir" yaml:"workDir"`

	Status _const.ServiceStatus `json:"status" yaml:"status"`

	OsPlatform _const.OsCategory `json:"osPlatform" yaml:"osPlatform"`
	OsType     _const.OsType     `json:"osType" yaml:"osType"`
	SysLang    _const.SysLang    `json:"sysLang" yaml:"sysLang"`
}
