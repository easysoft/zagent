package commDomain

type KvmReq struct {
	StartAfterCreated bool `json:"startAfterCreated"`

	VmMacAddress string `json:"vmMacAddress"`
	VmBacking    string `json:"vmBacking"`
	VmTemplate   string `json:"vmTemplate"`

	VmUniqueName  string `json:"vmUniqueName"`
	VmMemorySize  int    `json:"vmMemorySize"`
	VmDiskSize    int    `json:"vmDiskSize"`
	VmCdromSys    string `json:"vmCdromSys"`
	VmCdromDriver string `json:"vmCdromDriver"`
}

type KvmResponse struct {
	Code    int    `json:"code"`
	Msg     int    `json:"msg"`
	Name    string `json:"name"`
	VncPort int    `json:"vncPort"`
	Path    string `json:"path"`
	Mac     string `json:"mac"`
}
