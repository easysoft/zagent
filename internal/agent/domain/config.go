package domain

type Config struct {
	Host string
	User string

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
	DirBase  string
	//DirDef   string
	//DirTempl string
}
