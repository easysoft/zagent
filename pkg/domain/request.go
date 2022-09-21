package _domain

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
