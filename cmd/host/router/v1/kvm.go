package v1

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
)

type CreateVmReq struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Os   string `json:"os" enums:"win10,ubuntu20.04"` // From 'osinfo-query os' command

	Cpu    int `json:"cpu"`
	Memory int `json:"memory"` // Unit G
	Disk   int `json:"disk"`   // Unit G

	TaskId string `json:"taskId"`
}

type CreateVmResp struct {
	Mac    string          `json:"mac"`
	Vnc    int             `json:"vnc"`
	Status consts.VmStatus `json:"status"` // Enums consts.VmStatus
}

type ExportVmReq struct {
	Vm         string `json:"vm"`
	Backing    string `json:"backing"`
	ZentaoTask int    `json:"zentaoTask"`
}

type ExportVmResp struct {
	Backing string            `json:"backing"`
	Xml     string            `json:"xml"`
	Status  consts.TaskStatus `json:"status"` // Enums consts.TaskStatus

	CompletionRate float64 `json:"completionRate"`
	Speed          float64 `json:"speed"`

	ZentaoTask int `json:"zentaoTask"`
}

type CloneVmReq struct {
	VmMacAddress string `json:"vmMacAddress"`
	VmUniqueName string `json:"vmUniqueName" example:"test-win10-x64-pro-zh_cn-clone1"`
	VmSrc        string `json:"vmSrc" example:"test-win10-x64-pro-zh_cn"`

	VmCpu        uint `json:"vmCpu" example:"3"`
	VmMemorySize uint `json:"vmMemorySize" example:"5"` // Unit G
	VmDiskSize   uint `json:"vmDiskSize" example:"60"`  // Unit G

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
	Status  consts.VmStatus `json:"status"` // Enums consts.VmStatus
}

//type KvmRespTempl struct {
//	Name string `json:"name"`
//	Type string `json:"type"`
//	UUID string `json:"uuid"`
//
//	CpuCoreNum  uint   `json:"cpuCoreNum"`
//	MemoryValue uint   `json:"memoryValue"`
//	MemoryUnit  string `json:"memoryUnit"`
//
//	OsArch        string `json:"osArch"`
//	MacAddress    string `json:"macAddress"`
//	DiskFile      string `json:"diskFile"`
//	BackingFile   string `json:"backingFile"`
//	BackingFormat string `json:"backingFormat"`
//
//	VncPost int `json:"memoryValue"`
//}
