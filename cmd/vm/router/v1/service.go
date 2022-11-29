package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type VmServiceCheckReq struct {
	Services string `json:"services"`
}

type VmServiceCheckResp struct {
	ZtfStatus  consts.HostServiceStatus `json:"ztfStatus"` // Enums consts.HostServiceStatus
	ZtfVersion string                   `json:"ztfVersion"`

	ZdStatus  consts.HostServiceStatus `json:"zdStatus"` // Enums consts.HostServiceStatus
	ZdVersion string                   `json:"zdVersion"`
}

type VmServiceInstallReq struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type VmServiceInstallResp struct {
	Version string `json:"version"`
}
