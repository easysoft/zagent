package commDomain

import "time"

type TestResult struct {
	TestSetId uint `json:"testSetId" yaml:"testSetId"`

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	StartTime time.Time `json:"startTime" yaml:"startTime"`
	EndTime   time.Time `json:"endTime" yaml:"endTime"`
	Duration  int       `json:"duration" yaml:"duration"` // sec

	Total  int `json:"total" yaml:"total"`
	Pass   int `json:"pass" yaml:"pass"`
	Fail   int `json:"fail" yaml:"fail"`
	Missed int `json:"missed" yaml:"missed"`
}
