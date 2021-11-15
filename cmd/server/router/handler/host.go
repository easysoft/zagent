package handler

import (
	"github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	serverService "github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type HostCtrl struct {
	BaseCtrl

	AssertService *serverService.AssertService `inject:""`
}

func NewHostCtrl() *HostCtrl {
	return &HostCtrl{}
}

func (c *HostCtrl) Register(ctx iris.Context) {
	model := domain.HostNode{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	success := c.AssertService.RegisterHost(model)
	code := _const.ResultFail
	if success {
		code = _const.ResultSuccess
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(code, "操作成功", ""))
	return
}
