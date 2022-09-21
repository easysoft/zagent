package domain

import _const "github.com/easysoft/zv/pkg/const"

var (
	ResultVmNotFound = BizErr{-10100, "RunModeVm Not Found"}
)

type BizErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e BizErr) Error() string {
	return e.Msg
}

func NewBizErr(msg string) BizErr {
	return BizErr{Code: _const.ResultFail.Int(), Msg: msg}
}

func NewBizErrWithCode(code int, msg string) BizErr {
	return BizErr{Code: code, Msg: msg}
}
