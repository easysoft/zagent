package _domain

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
)

type RemoteReq struct {
	ComputerIp   string
	ComputerPort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}

type RemoteResp struct {
	Code    consts.ResultCode `json:"code"`
	Msg     string            `json:"msg"`
	Payload interface{}       `json:"payload"`
}

func (result *RemoteResp) Pass(msg string) {
	result.Code = consts.ResultPass
	result.Msg = msg
}

func (result *RemoteResp) Passf(str string, args ...interface{}) {
	result.Code = consts.ResultPass
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RemoteResp) Fail(msg string) {
	result.Code = consts.ResultFail
	result.Msg = msg
}

func (result *RemoteResp) Failf(str string, args ...interface{}) {
	result.Code = consts.ResultFail
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RemoteResp) IsSuccess() bool {
	return result.Code == consts.ResultPass
}
