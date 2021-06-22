package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type HostNode struct {
	Node
	HostStatus commConst.HostStatus `json:"hostStatus" yaml:"hostStatus"`
	Vms        []Vm
}

type VmNode struct {
	Node
	VmStatus      commConst.VmStatus      `json:"vmStatus" yaml:"vmStatus"`
	ServiceStatus commConst.ServiceStatus `json:"serviceStatus" yaml:"serviceStatus"`
	taskCount     int
}

type DeviceNode struct {
	Node
	ServiceStatus commConst.ServiceStatus `json:"serviceStatus" yaml:"serviceStatus"`
	taskCount     int
}

type Node struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Ip      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	WorkDir string `json:"workDir" yaml:"workDir"`

	OsPlatform commConst.OsCategory `json:"osPlatform" yaml:"osPlatform"`
	OsType     commConst.OsType     `json:"osType" yaml:"osType"`
	SysLang    commConst.OsLang     `json:"sysLang" yaml:"sysLang"`

	SshPort int `json:"sshPort" yaml:"sshPort"`
	VncPort int `json:"vncPort" yaml:"vncPort"`

	LastRegisterDate time.Time `json:"lastRegisterDate" yaml:"lastRegisterDate"`
}
