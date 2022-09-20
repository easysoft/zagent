package domain

import "time"

type VmNotifyReq struct {
	MacAddress string `json:"macAddress"`
}

type VmNotifyResp struct {
	Secret string `json:"secret"`
	Ip     string `json:"ip"`
}

type RegisterResp struct {
	Token           string    `json:"token" yaml:"token"`
	ExpiredTimeUnix int64     `json:"expiredTimeUnix" yaml:"expiredTimeUnix"`
	ExpiredDate     time.Time `json:"expiredDate" yaml:"expiredDate"`
}
