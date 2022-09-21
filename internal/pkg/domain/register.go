package domain

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type VmNotifyReq struct {
	MacAddress string `json:"macAddress"`
}

type VmNotifyResp struct {
	Secret          string `json:"secret"`
	Ip              string `json:"ip"`
	AgentPortOnHost int    `json:"agentPortOnHost"`
}

type RegisterResp struct {
	Token           string    `json:"token" yaml:"token"`
	ExpiredTimeUnix int64     `json:"expiredTimeUnix" yaml:"expiredTimeUnix"`
	ExpiredDate     time.Time `json:"expiredDate" yaml:"expiredDate"`
}

type VmRegisterReq struct {
	Secret          string          `json:"secret"`
	MacAddress      string          `json:"macAddress"`
	Ip              string          `json:"ip"`
	AgentPortOnHost int             `json:"agentPortOnHost"`
	Status          consts.VmStatus `json:"status"`
}
