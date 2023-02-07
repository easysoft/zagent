package domain

import (
	"fmt"
	"time"

	consts "github.com/easysoft/zagent/internal/pkg/const"
)

type TestResult struct {
	TestSetId uint `json:"testSetId" yaml:"testSetId"`

	Version float64           `json:"version" yaml:"version"`
	Name    string            `json:"name" yaml:"name"`
	Code    consts.ResultCode `json:"code"`
	Msg     string            `json:"msg"`

	StartTime time.Time `json:"startTime" yaml:"startTime"`
	EndTime   time.Time `json:"endTime" yaml:"endTime"`
	Duration  int       `json:"duration" yaml:"duration"` // sec

	TotalNum  int `json:"totalNum" yaml:"totalNum"`
	PassNum   int `json:"passNum" yaml:"passNum"`
	FailNum   int `json:"failNum" yaml:"failNum"`
	MissedNum int `json:"missedNum" yaml:"missedNum"`

	Payload interface{} `json:"payload"`
}

func (result *TestResult) Pass(msg string) {
	result.Code = consts.ResultPass
	result.Msg = msg
}
func (result *TestResult) Passf(str string, args ...interface{}) {
	result.Code = consts.ResultPass
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestResult) Fail(msg string) {
	result.Code = consts.ResultFail
	result.Msg = msg
}

func (result *TestResult) Failf(str string, args ...interface{}) {
	result.Code = consts.ResultFail
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestResult) IsSuccess() bool {
	return result.Code == consts.ResultPass
}
