package _httpUtils

import (
	"github.com/easysoft/zv/internal/comm/domain"
	_const "github.com/easysoft/zv/pkg/const"
)

type Request struct {
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
}

type Response struct {
	Code interface{} `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
type ResponsePage struct {
	Response

	PageSize   int   `json:"pageSize"`
	PageNo     int   `json:"pageNo"`
	TotalCount int64 `json:"totalCount"`
	TotalPage  int64 `json:"totalPage"`
}

func RespData(code _const.ResultCode, msg string, objects interface{}) (r *Response) {
	r = &Response{Code: code, Msg: msg, Data: objects}
	return
}

func RespDataPagination(code _const.ResultCode, msg string, objects interface{}, pageNo, pageSize int, total int64) (r *ResponsePage) {
	r = &ResponsePage{Response: Response{Code: code, Msg: msg, Data: objects},
		PageSize: pageSize, PageNo: pageNo, TotalCount: total, TotalPage: total / int64(pageSize)}
	return
}

func RespDataFromBizErr(err *domain.BizErr) (r *Response) {
	r = &Response{Code: err.Code, Msg: err.Msg}
	return
}
