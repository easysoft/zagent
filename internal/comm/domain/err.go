package domain

import _const "github.com/easysoft/zv/pkg/const"

var (
	ResultVmNotFound = BizErr{-10100, "Vm Not Found"}
)

type BizErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e BizErr) Error() string {
	return e.Msg
}

func NewBizErr(msg string) BizErr {
	return BizErr{_const.ResultFail.Int(), msg}
}

func NewBizErrWithCode(code int, msg string) BizErr {
	return BizErr{code, msg}
}
