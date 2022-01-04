package domain

type MultiPass struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Image string `json:"image"`
	IPv4  string `json:"i_pv_4"`
}
