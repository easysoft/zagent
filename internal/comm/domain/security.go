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
	ExpiredDate time.Time `json:"expiredDate" yaml:"expiredDate"`
}
