package v1

type VirtualBoxReq struct {
	VmUniqueName string `json:"vmUniqueName" example:"test-win10-x64-pro-zh_cn"`
	BackingName  string `json:"backingName"`

	Bridge       string `json:"bridge"`
	VmCpu        uint   `json:"vmCpu"`
	VmMemorySize uint   `json:"vmMemorySize"`
	VncPort      int    `json:"vncPort"`

	CloudIamUser     string `json:"cloudIamUser"`
	CloudIamPassword string `json:"cloudIamPassword"`

	Prefix string `json:"prefix"`
}

type VirtualBoxResp struct {
	Name        string `json:"name"`
	MacAddress  string `json:"macAddress"`
	VncPort     int    `json:"vncPort"`
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`
}
