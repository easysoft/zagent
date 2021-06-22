package domain

import agentConst "github.com/easysoft/zagent/internal/agent/utils/const"

type Config struct {
	Host string
	User string

	RunMode  agentConst.RunMode `yaml:"runMode"`
	Server   string             `yaml:"Server"`
	NodeIp   string             `yaml:"ip"`
	NodePort int                `yaml:"port"`

	Language string
	NodeName string
	WorkDir  string
	LogDir   string

	DirKvm   string
	DirIso   string
	DirImage string
	DirBase  string
	//DirDef   string
	//DirTempl string
}
