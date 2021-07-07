package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Vm struct {
	BaseModel

	Name      string `json:"name"`
	Desc      string `json:"desc"`
	ImagePath string `json:"imagePath"`

	HostId   uint   `json:"hostId"`
	HostName string `json:"hostName"`

	TmplId   uint   `json:"tmplId"`
	TmplName string `json:"tmplName"`

	BackingId   uint   `json:"backingId"`
	BackingPath string `json:"backingPath"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang"`

	Status    consts.VmStatus `json:"status"`
	DestroyAt *time.Time      `json:"destroyAt"`

	FirstDetectedTime *time.Time `json:"firstDetectedTime"`
	LastRegisterTime  *time.Time `json:"lastRegisterTime"`

	PublicIp   string `json:"publicIp"`
	PublicPort int    `json:"publicPort"`
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
}

func GenKvmReq(po Vm) (req domain.KvmReq) {
	req = domain.KvmReq{
		VmMacAddress: po.MacAddress, VmUniqueName: po.Name,
		VmBackingPath: po.BackingPath, VmTemplateName: po.TmplName,

		OsCategory: po.OsCategory,
		OsType:     po.OsType,
		OsVersion:  po.OsVersion,
		OsLang:     po.OsLang,

		VmDiskSize: po.DiskSize, VmMemorySize: po.MemorySize,
		VmCdromSys: po.CdromSys, VmCdromDriver: po.CdromDriver}

	return
}

func VmFromDomain(v domain.Vm) (po Vm) {
	po = Vm{
		Status:    v.Status,
		DestroyAt: &v.DestroyAt,

		PublicIp:   v.PublicIp,
		PublicPort: v.PublicPort,
		MacAddress: v.MacAddress,
		RpcPort:    v.RpcPort,
		SshPort:    v.SshPort,
		VncPort:    v.VncPort,
		WorkDir:    v.WorkDir,
	}

	return
}

func (Vm) TableName() string {
	return "biz_vm"
}
