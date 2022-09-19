package domain

import "time"

type SecurityReq struct {
	MacAddress string `json:"macAddress"`
}

type SecurityResp struct {
	Secret string `json:"secret"`
}

type RegisterResp struct {
	Token       string    `json:"token" yaml:"token"`
	ExpiredTime int64     `json:"expiredTime" yaml:"expiredTime"`
	ExpiredDate time.Time `json:"expiredDate" yaml:"expiredDate"`
}
