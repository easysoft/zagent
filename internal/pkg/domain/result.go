package domain

import consts "github.com/easysoft/zv/internal/pkg/const"

type Result struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func (result *Result) Success(msg string) {
	result.Code = consts.ResultPass.String()
	result.Msg = msg
}
