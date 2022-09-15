package v1

type VmPortMapReq struct {
	VmIp   string `json:"vmIp"`
	VmPort int    `json:"vmPort"`
	HostIp string `json:"hostIp"`
}

type VmPortMapResp struct {
	VmIp     string `json:"vmIp"`
	VmPort   int    `json:"vmPort"`
	HostIp   string `json:"hostIp"`
	HostPort int    `json:"hostPort"`
}

type VncTokenResp struct {
	Token string `json:"token"`
	Ip    string `json:"ip"`
	Port  string `json:"port"`
}
