package _domain

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type RpcReq struct {
	ComputerIp   string
	ComputerPort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}

type RpcResp struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func (result *RpcResp) Pass(msg string) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = msg
}

func (result *RpcResp) Passf(str string, args ...interface{}) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RpcResp) Fail(msg string) {
	result.Code = _const.ResultFail.Int()
	result.Msg = msg
}

func (result *RpcResp) Failf(str string, args ...interface{}) {
	result.Code = _const.ResultFail.Int()
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RpcResp) IsSuccess() bool {
	return result.Code == _const.ResultSuccess.Int()
}
