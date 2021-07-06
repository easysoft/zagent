package domain

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Vm struct {
	ID        int
	BackingId int
	HostId    int

	Name        string
	Tmpl        string
	Backing     string
	ImagePath   string
	BackingPath string

	OsCategory consts.OsCategory
	OsType     consts.OsType
	OsVersion  string
	OsLang     consts.OsLang

	Status            consts.VmStatus
	DestroyAt         time.Time
	FirstDetectedTime time.Time

	PublicIp   string
	PublicPort int
	MacAddress string
	RpcPort    int
	SshPort    int
	VncPort    int
	WorkDir    string

	DefPath          string
	DiskSize         uint // M
	MemorySize       uint // M
	CdromSys         string
	CdromDriver      string
	ResolutionHeight int
	ResolutionWidth  int
}
