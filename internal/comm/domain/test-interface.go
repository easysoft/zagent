package domain

type TestInterface struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Request   Request    `json:"request" yaml:"request" gorm:"-"`
	Responses []Response `json:"responses" yaml:"responses" gorm:"-"`

	Raws []string `json:"raw" yaml:"-" gorm:"-"`
}
