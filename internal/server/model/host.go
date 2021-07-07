package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"time"
)

type Host struct {
	BaseModel

	Name string `json:"name"`

	OsCategory consts.OsCategory `json:"osCategory"`
	OsType     consts.OsType     `json:"osType"`
	OsLang     consts.OsLang     `json:"osLang"`

	OsVersion string `json:"osVersion"`
	OsBuild   string `json:"osBuild"`
	OsBits    string `json:"osBits"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir"`

	SshPort int               `json:"sshPort"`
	VncPort int               `json:"vncPort"`
	Status  consts.HostStatus `json:"status"`

	TaskCount        int        `json:"taskCount"`
	MaxVmNum         int        `json:"maxVmNum"`
	LastRegisterTime *time.Time `json:"lastRegisterTime"`

	Backings []VmBacking `gorm:"many2many:host_backing_r;"`
	Vms      []Vm        `json:"vms" gorm:"-"`
}

func NewHost() Host {
	host := Host{}
	return host
}

func HostFromDomain(h domain.Host) (po Host) {
	po = Host{
		Name: h.Name,

		OsCategory: h.OsCategory,
		OsType:     h.OsType,
		OsLang:     h.OsLang,

		OsVersion: h.OsVersion,
		OsBuild:   h.OsBuild,
		OsBits:    h.OsBits,

		Ip:      h.Ip,
		Port:    h.Port,
		WorkDir: h.WorkDir,

		SshPort: h.SshPort,
		VncPort: h.VncPort,
		Status:  h.Status,
	}

	if po.Name == "" {
		po.Name = po.Ip
	}

	for _, v := range h.Vms {
		vm := VmFromDomain(v)
		po.Vms = append(po.Vms, vm)
	}

	return
}

func (Host) TableName() string {
	return "biz_host"
}
