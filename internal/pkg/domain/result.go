package domain

import _const "github.com/easysoft/zv/pkg/const"

type Result struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func (result *Result) Success(msg string) {
	result.Code = _const.ResultPass.Int()
	result.Msg = msg
}
