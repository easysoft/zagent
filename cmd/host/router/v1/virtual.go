package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type DestroyVmReq struct {
	Ip string `json:"ip"`
}

type VmPortMapReq struct {
	VmIp   string                `json:"vmIp"`
	VmPort int                   `json:"vmPort"`
	Type   consts.NatForwardType `json:"type"`
}

type VmPortMapResp struct {
	VmIp   string                `json:"vmIp"`
	VmPort int                   `json:"vmPort"`
	Type   consts.NatForwardType `json:"type"`

	HostPort      int  `json:"hostPort"`
	AlreadyMapped bool `json:"alreadyMapped"`
}

type VncTokenResp struct {
	Token string `json:"token"`
	Ip    string `json:"ip"`
	Port  string `json:"port"`
}
