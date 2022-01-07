package domain

type MultiPass struct {
	Name    string `json:"name"`
	Cpus    string `json:"cpus"`
	Memory  string `json:"memory"`
	Network string `json:"network"`
	State   string `json:"state"`
	Image   string `json:"image"`
	IPv4    string `json:"i_pv_4"`
}
