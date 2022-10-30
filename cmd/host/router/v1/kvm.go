package v1

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
)

type CreateVmReq struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Os   string `json:"os"`

	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
	Disk   int `json:"disk"`

	TaskId string `json:"taskId"`
}

type CreateVmResp struct {
	Mac    string          `json:"mac"`
	Vnc    int             `json:"vnc"`
	Status consts.VmStatus `json:"status"`
}

type ExportVmReq struct {
	Vm         string `json:"vm"`
	Backing    string `json:"backing"`
	ZentaoTask int    `json:"zentaoTask"`
}

type ExportVmResp struct {
	Backing    string            `json:"backing"`
	Xml        string            `json:"xml"`
	Status     consts.TaskStatus `json:"status"`
	ZentaoTask int               `json:"zentaoTask"`
}

type KvmReq struct {
	VmMacAddress string `json:"vmMacAddress"`
	VmTemplate   string `json:"vmTemplate"`
	VmBacking    string `json:"vmBacking"`

	VmUniqueName string `json:"vmUniqueName" example:"test-win10-x64-pro-zh_cn"`
	VmCpu        uint   `json:"vmCpu" example:"3"`
	VmMemorySize uint   `json:"vmMemorySize" example:"5120000"`
	VmDiskSize   uint   `json:"vmDiskSize" example:"30000"`
	//VmCdromSys    string `json:"vmCdromSys"`
	//VmCdromDriver string `json:"vmCdromDriver"`

	OsCategory consts.OsCategory `json:"osCategory" example:"windows"` // Enums consts.OsCategory
	OsType     consts.OsType     `json:"osType" example:"win10"`       // Enums consts.OsType
	OsVersion  string            `json:"osVersion" example:"x64-pro"`
	OsLang     consts.OsLang     `json:"osLang" example:"zh_cn"` // Enums consts.OsLang

	//StartAfterCreated bool `json:"startAfterCreated"`
}

type KvmReqClone struct {
	VmMacAddress string `json:"vmMacAddress"`
	VmUniqueName string `json:"vmUniqueName" example:"test-win10-x64-pro-zh_cn-clone1"`
	VmSrc        string `json:"vmSrc" example:"test-win10-x64-pro-zh_cn"`

	VmCpu        uint `json:"vmCpu" example:"3"`
	VmMemorySize uint `json:"vmMemorySize" example:"5120000"`
	VmDiskSize   uint `json:"vmDiskSize" example:"30000"`

	//StartAfterCreated bool `json:"startAfterCreated"`
}

type KvmResp struct {
	Name    string          `json:"name"`
	Ip      string          `json:"ip"`
	Mac     string          `json:"mac"`
	Agent   int             `json:"agent"`
	Vnc     int             `json:"vnc"`
	VncUrl  string          `json:"vncUrl"`
	Image   string          `json:"image"`
	Backing string          `json:"backing"`
	Status  consts.VmStatus `json:"status"`
}
type KvmRespTempl struct {
	Name string `json:"name"`
	Type string `json:"type"`
	UUID string `json:"uuid"`

	CpuCoreNum  uint   `json:"cpuCoreNum"`
	MemoryValue uint   `json:"memoryValue"`
	MemoryUnit  string `json:"memoryUnit"`

	OsArch        string `json:"osArch"`
	MacAddress    string `json:"macAddress"`
	DiskFile      string `json:"diskFile"`
	BackingFile   string `json:"backingFile"`
	BackingFormat string `json:"backingFormat"`

	VncPost int `json:"memoryValue"`
}
