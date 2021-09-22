package domain

type VmWareReq struct {
	VmId         string `json:"vmId"` // for destroy
	BackingName  string `json:"backingName"`
	VmUniqueName string `json:"vmName"`

	Processors uint `json:"processors"`
	Memory     uint `json:"memory"`

	UserName string `json:"userName"`
	Password string `json:"password"`

	StartAfterCreated bool `json:"startAfterCreated"`
}

type VmWareResp struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`

	VmId       string `json:"vmId"`
	Name       string `json:"name"`
	Mac        string `json:"mac"`
	VncAddress string `json:"vncAddress"`
}

type VmWareParam struct {
	Processors int `json:"processors"`
	Memory     int `json:"memory"`
}
