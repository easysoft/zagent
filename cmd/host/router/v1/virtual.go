package v1

import consts "github.com/easysoft/zv/internal/comm/const"

type VmPortMapReq struct {
	VmIp   string                `json:"vmIp"`
	VmPort int                   `json:"vmPort"`
	HostIp string                `json:"hostIp"`
	Type   consts.NatForwardType `json:"type"`
}

type VmPortMapResp struct {
	VmIp     string                `json:"vmIp"`
	VmPort   int                   `json:"vmPort"`
	HostIp   string                `json:"hostIp"`
	HostPort int                   `json:"hostPort"`
	Type     consts.NatForwardType `json:"type"`
}

type VncTokenResp struct {
	Token string `json:"token"`
	Ip    string `json:"ip"`
	Port  string `json:"port"`
}
