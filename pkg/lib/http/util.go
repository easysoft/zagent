package _httpUtils

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	_domain "github.com/easysoft/zagent/pkg/domain"
)

func RespData(code consts.ResultCode, msg string, objects interface{}) (r *_domain.Response) {
	r = &_domain.Response{Code: code, Msg: msg, Data: objects}
	return
}

func RespDataPagination(code consts.ResultCode, msg string, objects interface{}, pageNo, pageSize int, total int64) (r *_domain.ResponsePage) {
	r = &_domain.ResponsePage{Response: _domain.Response{Code: code, Msg: msg, Data: objects},
		PageSize: pageSize, PageNo: pageNo, TotalCount: total, TotalPage: total / int64(pageSize)}
	return
}

func RespDataFromBizErr(err error) (r *_domain.Response) {
	bizErr := err.(domain.BizErr)

	r = &_domain.Response{Code: bizErr.Code, Msg: bizErr.Msg}
	return
}
