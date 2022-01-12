package v1

type MultiPassReq struct {
	VmId string `json:"vmId"` // for destroy

	VmUniqueName string `json:"vmName"`
	VmMemory     string `json:"vmMemory"`
	Cpus         string `json:"cpus"`
	Disk         string `json:"disk"`
	FilePath     string `json:"vmRauPath"`

	UserName string `json:"userName"`
	Password string `json:"password"`
}

type MultiPassResp struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`

	Name    string `json:"name"`
	Cpus    string `json:"cpus"`
	Memory  string `json:"memory"`
	Network string `json:"network"`
	State   string `json:"state"`
	Image   string `json:"image"`
	IPv4    string `json:"ipv4"`
}

type MultiPassParam struct {
	Processors int `json:"processors"`
	Memory     int `json:"memory"`
}
