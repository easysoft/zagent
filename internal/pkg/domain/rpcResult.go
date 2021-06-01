package _domain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type RpcResult struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func (result *RpcResult) Success(msg string) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = msg
}

func (result *RpcResult) Fail(msg string) {
	result.Code = _const.ResultFail.Int()
	result.Msg = msg
}

func (result *RpcResult) IsSuccess() bool {
	return result.Code == _const.ResultSuccess.Int()
}
