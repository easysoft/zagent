package domain

type Config struct {
	Server   string `yaml:"Server"`
	NodeIp   string `yaml:"ip"`
	NodePort int    `yaml:"port"`

	Language string
	NodeName string
	WorkDir  string
	LogDir   string

	DirKvm   string
	DirIso   string
	DirImage string
	DirDef   string
	DirTempl string
}
