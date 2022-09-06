package _httpUtils

import _const "github.com/easysoft/zv/pkg/const"

type Request struct {
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
	Aa       int `json:"aa"`
}

type Response struct {
	Code _const.ResultCode `json:"code"`
	Msg  string            `json:"message"`
	Data interface{}       `json:"data"`
}
type ResponsePage struct {
	Response

	PageSize   int   `json:"pageSize"`
	PageNo     int   `json:"pageNo"`
	TotalCount int64 `json:"totalCount"`
	TotalPage  int64 `json:"totalPage"`
}

func ApiRes(code _const.ResultCode, msg string, objects interface{}) (r *Response) {
	r = &Response{Code: code, Msg: msg, Data: objects}
	return
}
func ApiResPage(code _const.ResultCode, msg string, objects interface{}, pageNo, pageSize int, total int64) (r *ResponsePage) {
	r = &ResponsePage{Response: Response{Code: code, Msg: msg, Data: objects},
		PageSize: pageSize, PageNo: pageNo, TotalCount: total, TotalPage: total / int64(pageSize)}
	return
}
