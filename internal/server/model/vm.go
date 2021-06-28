package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Vm struct {
	BaseModel

	BaseId int `json:"baseId"`
	HostId int `json:"hostId"`

	Name      string `json:"name"`
	Src       string `json:"src"`
	Base      string `json:"base"`
	ImagePath string `json:"imagePath"`
	BasePath  string `json:"basePath"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsVersion  string               `json:"osVersion"`
	OsLang     commConst.OsLang     `json:"osLang"`

	Status            commConst.VmStatus `json:"status"`
	DestroyAt         time.Time          `json:"destroyAt"`
	FirstDetectedTime time.Time          `json:"firstDetectedTime"`

	PublicIp   string `json:"publicIp"`
	PublicPort int    `json:"publicPort"`
	MacAddress string `json:"macAddress"`
	RpcPort    int    `json:"rpcPort"`
	SshPort    int    `json:"sshPort"`
	VncPort    int    `json:"vncPort"`
	WorkDir    string `json:"workDir"`

	DefPath          string `json:"defPath"`
	DiskSize         int    `json:"diskSize"`   // M
	MemorySize       int    `json:"memorySize"` // M
	CdromSys         string `json:"cdromSys"`
	CdromDriver      string `json:"cdromDriver"`
	ResolutionHeight int    `json:"resolutionHeight"`
	ResolutionWidth  int    `json:"resolutionWidth"`
}

func (Vm) TableName() string {
	return "biz_vm"
}
