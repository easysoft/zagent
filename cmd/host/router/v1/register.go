package v1

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

// vm notity host being ready
type VmNotifyReq struct {
	MacAddress string `json:"macAddress"`
}

type VmNotifyResp struct {
	Secret          string `json:"secret"`
	Ip              string `json:"ip"`
	AgentPortOnHost int    `json:"agentPortOnHost"`
}

type VmRegisterReq struct {
	Secret          string          `json:"secret"`
	MacAddress      string          `json:"macAddress"`
	Ip              string          `json:"ip"`
	AgentPortOnHost int             `json:"agentPortOnHost"`
	Status          consts.VmStatus `json:"status"` // Enums consts.VmStatus
}

type RegisterResp struct {
	Token           string    `json:"token" yaml:"token"`
	ExpiredTimeUnix int64     `json:"expiredTimeUnix" yaml:"expiredTimeUnix"`
	ExpiredDate     time.Time `json:"expiredDate" yaml:"expiredDate"`
}
