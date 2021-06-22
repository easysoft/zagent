package commDomain

type ContainerInfo struct {
	Name    string `json:"name" yaml:"name"`
	Image   string `json:"image" yaml:"image"`
	ImageId string `json:"imageId" yaml:"imageId"`
	SshPort int    `json:"sshPort" yaml:"sshPort"`
}
