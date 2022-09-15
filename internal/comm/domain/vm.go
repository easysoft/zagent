package domain

import (
	"github.com/easysoft/zv/internal/comm/const"
	"time"
)

type Vm struct {
	ID        int `json:"id"`
	BackingId int `json:"backingId"`
	HostId    int `json:"hostId"`

	Name        string `json:"name"`
	Tmpl        string `json:"Tmpl"`
	Backing     string `json:"backing"`
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang"`

	Status            consts.VmStatus `json:"status"`
	DestroyAt         time.Time       `json:"destroyAt"`
	FirstDetectedTime time.Time       `json:"firstDetectedTime"`

	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	MacAddress string `json:"macAddress"`
	RpcPort    int    `json:"rpcPort"`
	SshPort    int    `json:"sshPort"`
	VncPort    int    `json:"vncPort"`
	WorkDir    string `json:"workDir"`

	DefPath          string `json:"defPath"`
	DiskSize         uint   `json:"diskSize"`   // M
	MemorySize       uint   `json:"memorySize"` // M
	CdromSys         string `json:"cdromSys"`
	CdromDriver      string `json:"cdromDriver"`
	ResolutionHeight int    `json:"resolutionHeight"`
	ResolutionWidth  int    `json:"resolutionWidth"`

	Secret string `json:"secret" yaml:"secret"`
}
