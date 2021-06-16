package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type Result struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func (result *Result) Success(msg string) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = msg
}