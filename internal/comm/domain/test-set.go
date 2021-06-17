package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type TestSet struct {
	TestConf

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor TestProcessor `json:"processor" yaml:"processor"`
}

type TestConf struct {
	NumberOfThreads int `json:"numberOfThreads,omitempty" yaml:"numberOfThreads,omitempty"`
	RampUpPeriod    int `json:"rampUpPeriod,omitempty" yaml:"rampUpPeriod,omitempty"` // sec

	Duration  int  `json:"duration,omitempty" yaml:"duration,omitempty"` // sec, if set, loopCount will be ignore
	LoopCount int  `json:"loopCount,omitempty" yaml:"loopCount,omitempty"`
	Forever   bool `json:"forever,omitempty" yaml:"forever,omitempty"`

	ErrorAction commConst.ErrorAction `json:"errorAction,omitempty" yaml:"errorAction,omitempty"`

	Agents []Node `json:"agents" yaml:"agents" gorm:"many2many:r_set_to_agent;"`
}
