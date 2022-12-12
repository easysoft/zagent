package v1

import (
	"time"

	consts "github.com/easysoft/zagent/internal/pkg/const"
)

// vm notity host being ready
type VmNotifyReq struct {
	MacAddress string `json:"macAddress"`
}

type VmNotifyResp struct {
	Token           string `json:"token"`
	Ip              string `json:"ip"`
	AgentPortOnHost int    `json:"agentPortOnHost"`
	Server          string `json:"server"`
}

type VmRegisterReq struct {
	Token           string          `json:"token"`
	MacAddress      string          `json:"macAddress"`
	Ip              string          `json:"ip"`
	AgentPortOnHost int             `json:"agentPortOnHost"`
	Status          consts.VmStatus `json:"status"` // Enums consts.VmStatus
}

type RegisterResp struct {
	Token         string    `json:"tokenSN" yaml:"token"`
	TokenTimeUnix int64     `json:"tokenTimeUnix" yaml:"tokenTimeUnix"`
	TokenTime     time.Time `json:"tokenTime" yaml:"tokenTime"`
}
