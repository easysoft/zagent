package v1

type VirtualBoxReq struct {
	VmUniqueName string `json:"vmUniqueName" example:"test-win10-x64-pro-zh_cn"`
	BackingName  string `json:"backingName"`

	VmCpu        uint `json:"vmCpu"`
	VmMemorySize uint `json:"vmMemorySize"`
	VncPort      int  `json:"vncPort"`

	Prefix string `json:"prefix"`
}

type VirtualBoxResp struct {
	Name        string `json:"name"`
	MacAddress  string `json:"macAddress"`
	VncPort     int    `json:"vncPort"`
	VncPassword string `json:"vncPassword"`
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`
}
