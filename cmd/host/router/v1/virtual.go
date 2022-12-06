package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type DestroyVmReq struct {
	Ip string `json:"ip"`
}

type VmPortMapReq struct {
	Ip   string                `json:"ip"`
	Port int                   `json:"port"`
	Type consts.NatForwardType `json:"type"` // Enums consts.NatForwardType
}

type VmPortMapResp struct {
	Ip   string                `json:"ip"`
	Port int                   `json:"port"`
	Type consts.NatForwardType `json:"type"` // Enums consts.NatForwardType

	HostPort      int  `json:"hostPort"`
	AlreadyMapped bool `json:"alreadyMapped"`
}

type VncTokenResp struct {
	Token string `json:"token"`
	Host  string `json:"host"`
	Port  string `json:"port"`
}
