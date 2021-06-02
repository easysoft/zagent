package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type TestSet struct {
	TestConf

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor TestProcessor `json:"processor" yaml:"processor"`
}

type TestConf struct {
	NumberOfThreads int `json:"numberOfThreads" yaml:"numberOfThreads"`
	RampUpPeriod    int `json:"rampUpPeriod" yaml:"rampUpPeriod"` // sec

	Duration  int  `json:"duration" yaml:"duration"` // sec, if set, loopCount will be ignore
	LoopCount int  `json:"loopCount" yaml:"loopCount"`
	Forever   bool `json:"forever" yaml:"forever"`

	ErrorAction _const.ErrorAction `json:"errorAction" yaml:"errorAction"`

	Agents []Node `json:"agents" yaml:"agents" gorm:"many2many:r_set_to_agent;"`
}
