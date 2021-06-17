package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Vm struct {
	Id     int
	BaseId int
	HostId int

	Name      string
	Src       string
	Base      string
	ImagePath string
	BasePath  string

	OsCategory commConst.OsCategory
	OsType     commConst.OsType
	OsVersion  string
	OsLang     commConst.OsLang

	Status            commConst.VmStatus
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
	DiskSize         int
	MemorySize       int
	CdromSys         string
	CdromDriver      string
	ResolutionHeight int
	ResolutionWidth  int
}