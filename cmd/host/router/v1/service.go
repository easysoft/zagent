package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type ServiceCheckReq struct {
	Services string `json:"services"`
}

type ServiceCheckResp struct {
	Code       string                   `json:"code"`
	Kvm        consts.HostServiceStatus `json:"kvm"`
	Novnc      consts.HostServiceStatus `json:"novnc"`
	Websockify consts.HostServiceStatus `json:"websockify"`
}