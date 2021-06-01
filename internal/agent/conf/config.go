package agentConf

type Config struct {
	Server   string `yaml:"Server"`
	NodeIp   string `yaml:"ip"`
	NodePort int    `yaml:"port"`

	Language string
	NodeName string
	WorkDir  string
	LogDir   string
}
