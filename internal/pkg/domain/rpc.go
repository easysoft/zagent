package _domain

import _const "github.com/easysoft/zagent/internal/pkg/const"

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

func (result *RpcResp) Success(msg string) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = msg
}

func (result *RpcResp) Fail(msg string) {
	result.Code = _const.ResultFail.Int()
	result.Msg = msg
}

func (result *RpcResp) IsSuccess() bool {
	return result.Code == _const.ResultSuccess.Int()
}
