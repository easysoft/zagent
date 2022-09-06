package _domain

import (
	"fmt"
	_const "github.com/easysoft/zv/pkg/const"
)

type RemoteReq struct {
	ComputerIp   string
	ComputerPort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}

type RemoteResp struct {
	Code    _const.ResultCode `json:"code"`
	Msg     string            `json:"msg"`
	Payload interface{}       `json:"payload"`
}

func (result *RemoteResp) Pass(msg string) {
	result.Code = _const.ResultSuccess
	result.Msg = msg
}

func (result *RemoteResp) Passf(str string, args ...interface{}) {
	result.Code = _const.ResultSuccess
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RemoteResp) Fail(msg string) {
	result.Code = _const.ResultFail
	result.Msg = msg
}

func (result *RemoteResp) Failf(str string, args ...interface{}) {
	result.Code = _const.ResultFail
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RemoteResp) IsSuccess() bool {
	return result.Code == _const.ResultSuccess
}
