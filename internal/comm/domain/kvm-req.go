package domain

import commConst "github.com/easysoft/zagent/internal/comm/const"

type KvmReq struct {
	VmMacAddress   string `json:"vmMacAddress"`
	VmBackingPath  string `json:"vmBacking"`
	VmTemplateName string `json:"vmTemplate"`

	VmUniqueName  string `json:"vmUniqueName"`
	VmMemorySize  uint   `json:"vmMemorySize"`
	VmDiskSize    uint   `json:"vmDiskSize"`
	VmCdromSys    string `json:"vmCdromSys"`
	VmCdromDriver string `json:"vmCdromDriver"`

	OsCategory commConst.OsCategory `json:"osCategory"`
	OsType     commConst.OsType     `json:"osType"`
	OsVersion  string               `json:"osVersion"`
	OsLang     commConst.OsLang     `json:"osLang"`

	StartAfterCreated bool `json:"startAfterCreated"`
}

type KvmResponse struct {
	Code    int    `json:"code"`
	Msg     int    `json:"msg"`
	Name    string `json:"name"`
	VncPort int    `json:"vncPort"`
	Path    string `json:"path"`
	Mac     string `json:"mac"`
}
