package domain

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
)

var (
	ResultVmNotFound = BizErr{consts.BizVmNotExist, "Vm Not Found"}
)

type BizErr struct {
	Code consts.ResultCode `json:"code"`
	Msg  string            `json:"msg"`
}

func (e BizErr) Error() string {
	return e.Msg
}

func NewBizErr(msg string) BizErr {
	return BizErr{Code: consts.ResultFail, Msg: msg}
}

func NewBizErrWithCode(code consts.ResultCode, msg string) BizErr {
	return BizErr{Code: code, Msg: msg}
}
