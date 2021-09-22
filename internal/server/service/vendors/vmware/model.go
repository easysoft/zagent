package vmwareService

type VmWareStatus string

const (
	On       VmWareStatus = "on"
	Off      VmWareStatus = "off"
	Shutdown VmWareStatus = "shutdown"
	Suspend  VmWareStatus = "suspend"
	Pause    VmWareStatus = "pause"
	Unpause  VmWareStatus = "unpause"
)

func (e VmWareStatus) ToString() string {
	return string(e)
}

type Vm struct {
	IdVM         string `json:"id"`
	Path         string `json:"path"`
	Denomination string `json:"displayName"`
	Description  string `json:"annotation"`
	// Image        string `json:"image"`
	CPU struct {
		Processors int `json:"processors"`
	}
	PowerStatus string `json:"power_state"`
	Memory      int    `json:"memory"`
}

type NicResp struct {
	Num  string `json:"num"`
	Nics []Nic  `json:"nics"`
}

type Nic struct {
	Index      int    `json:"index"`
	Type       string `json:"type"`
	Vmnet      string `json:"vmnet"`
	MacAddress string `json:"macAddress"`
}
