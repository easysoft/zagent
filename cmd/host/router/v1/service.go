package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type HostServiceCheckReq struct {
	Services string `json:"services"`
}

type HostServiceCheckResp struct {
	Kvm        consts.HostServiceStatus `json:"kvm"`        // Enums consts.HostServiceStatus
	Nginx      consts.HostServiceStatus `json:"nginx"`      // Enums consts.HostServiceStatus
	Novnc      consts.HostServiceStatus `json:"novnc"`      // Enums consts.HostServiceStatus
	Websockify consts.HostServiceStatus `json:"websockify"` // Enums consts.HostServiceStatus
}
