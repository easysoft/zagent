package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type Node struct {
	Version float64 `yaml:"version"`
	Name    string  `yaml:"name"`
	Desc    string  `yaml:"desc"`

	Ip   string `yaml:"ip"`
	Port int    `yaml:"desc"`

	Status _const.ServiceStatus `yaml:"status"`

	OsPlatform _const.OsPlatform `yaml:"osType"`
	OsType     _const.OsType     `yaml:"osType"`
	SysLang    _const.SysLang    `yaml:"osType"`
}
