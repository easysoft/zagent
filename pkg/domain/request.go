package _domain

import consts "github.com/easysoft/zagent/internal/pkg/const"

type Request struct {
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
}

type Response struct {
	Code consts.ResultCode `json:"code"` // Enums consts.ResultCode
	Msg  string            `json:"msg,omitempty"`
	Data interface{}       `json:"data,omitempty"`
}
type ResponsePage struct {
	Response

	PageSize   int   `json:"pageSize"`
	PageNo     int   `json:"pageNo"`
	TotalCount int64 `json:"totalCount"`
	TotalPage  int64 `json:"totalPage"`
}
