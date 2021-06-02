package commDomain

type TestEnv struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`
}
