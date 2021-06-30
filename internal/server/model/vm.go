package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Vm struct {
	BaseModel

	BackingId uint `json:"baseId"`
	HostId    uint `json:"hostId"`

	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Tmpl        string `json:"tmpl"`
	Backing     string `json:"backing"`
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsVersion  string               `json:"osVersion"`
	OsLang     commConst.OsLang     `json:"osLang"`

	Status            commConst.VmStatus `json:"status"`
	DestroyAt         *time.Time         `json:"destroyAt"`
	FirstDetectedTime *time.Time         `json:"firstDetectedTime"`

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

func GenKvmReq(po Vm) (req commDomain.KvmReq) {
	req = commDomain.KvmReq{
		VmMacAddress: po.MacAddress, VmId: int(po.ID), VmUniqueName: po.Name,
		VmDiskSize: po.DiskSize, VmMemorySize: po.MemorySize,
		VmCdromSys: po.CdromSys, VmCdromDriver: po.CdromDriver, VmBackingImage: po.Backing}

	return
}

func (Vm) TableName() string {
	return "biz_vm"
}
