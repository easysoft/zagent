package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type VmServiceCheckReq struct {
	Services string `json:"services"` // tool name, ztf or zd
}

type VmServiceCheckResp struct {
	ZtfStatus  consts.HostServiceStatus `json:"ztfStatus"` // Enums consts.HostServiceStatus
	ZtfVersion string                   `json:"ztfVersion"`

	ZdStatus  consts.HostServiceStatus `json:"zdStatus"` // Enums consts.HostServiceStatus
	ZdVersion string                   `json:"zdVersion"`
}

type VmServiceInstallReq struct {
	Name    string `json:"name"`    // tool name, ztf or zd
	Version string `json:"version"` // tool version

	Server string `json:"server"` // zentao server url
	Secret string `json:"secret"` // secret to access zentao
	Ip     string `json:"ip"`     // testing node ip, port ztf:56202, zd:56203
}

type VmServiceInstallResp struct {
	Version string `json:"version"`
}
