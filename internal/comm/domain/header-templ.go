package commDomain

type HeaderTempl struct {
	Version   float64 `json:"version" yaml:"version"`
	Name      string  `json:"name" yaml:"name"`
	Desc      string  `json:"desc" yaml:"desc"`
	IsDefault bool    `json:"isDefault" yaml:"isDefault"`

	Items []HeaderItem `json:"items" yaml:"items"`
}

type HeaderItem struct {
	Name    string `json:"name" yaml:"name"`
	Desc    string `json:"desc" yaml:"desc"`
	Default string `json:"default" yaml:"default"`
}
