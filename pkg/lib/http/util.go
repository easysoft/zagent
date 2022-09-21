package _httpUtils

import (
	"github.com/easysoft/zv/internal/pkg/domain"
	_const "github.com/easysoft/zv/pkg/const"
	_domain "github.com/easysoft/zv/pkg/domain"
)

func RespData(code _const.ResultCode, msg string, objects interface{}) (r *_domain.Response) {
	r = &_domain.Response{Code: code, Msg: msg, Data: objects}
	return
}

func RespDataPagination(code _const.ResultCode, msg string, objects interface{}, pageNo, pageSize int, total int64) (r *_domain.ResponsePage) {
	r = &_domain.ResponsePage{Response: _domain.Response{Code: code, Msg: msg, Data: objects},
		PageSize: pageSize, PageNo: pageNo, TotalCount: total, TotalPage: total / int64(pageSize)}
	return
}

func RespDataFromBizErr(err *domain.BizErr) (r *_domain.Response) {
	r = &_domain.Response{Code: err.Code, Msg: err.Msg}
	return
}
