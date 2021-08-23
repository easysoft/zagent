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

	SshPort    int               `json:"sshPort"`
	VncAddress string            `json:"vncAddress"`
	Status     consts.HostStatus `json:"status"`

	TaskCount        int        `json:"taskCount"`
	MaxVmNum         int        `json:"maxVmNum"`
	LastRegisterTime *time.Time `json:"lastRegisterTime"`

	Backings []VmBacking `gorm:"many2many:host_backing_r;"`
	Vms      []Vm        `json:"vms" gorm:"-"`

	Priority int `json:"priority"`

	Platform consts.Platform `json:"platform"`

	CloudRegion        string `json:"cloudRegion"`
	CloudSecurityGroup string `json:"cloudSecurityGroup"`
	CloudKey           string `json:"cloudKey"`
	CloudSecret        string `json:"cloudSecret"`
	CloudNamespace     string `json:"cloudNamespace"`

	CloudUser        string `json:"cloudUser"`
	CloudIamUser     string `json:"cloudIamUser"`
	CloudIamPassword string `json:"cloudIamPassword"`
	CloudIamKey      string `json:"cloudIamKey"`
}

func NewHost() Host {
	host := Host{}
	return host
}

func HostFromDomain(h domain.HostNode) (po Host) {
	po = Host{
		Name:   h.Name,
		OsType: h.OsType,

		Ip:      h.Ip,
		Port:    h.Port,
		WorkDir: h.WorkDir,

		SshPort:    h.SshPort,
		VncAddress: h.VncAddress,
		Status:     h.Status,
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
