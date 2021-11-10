package domain

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Vm struct {
	ID        int
	BackingId int
	HostId    int

	Name        string `json:"name"`
	Tmpl        string
	Backing     string
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`

	OsCategory consts.OsCategory
	OsType     consts.OsType
	OsVersion  string
	OsLang     consts.OsLang

	Status            consts.VmStatus
	DestroyAt         time.Time
	FirstDetectedTime time.Time

	Ip         string
	Port       int
	MacAddress string `json:"macAddress"`
	RpcPort    int
	SshPort    int
	VncPort    string `json:"vncPort"`
	WorkDir    string

	DefPath          string
	DiskSize         uint // M
	MemorySize       uint // M
	CdromSys         string
	CdromDriver      string
	ResolutionHeight int
	ResolutionWidth  int
}
