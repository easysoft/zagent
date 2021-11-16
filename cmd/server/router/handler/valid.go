package handler

import (
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ValidCtrl struct {
	ValidService *serverService.ValidService `inject:""`
}

func NewValidCtrl() *ValidCtrl {
	return &ValidCtrl{}
}

func (c *ValidCtrl) Valid(ctx iris.Context) {
	model := domain.ValidRequest{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	result := c.ValidService.Valid(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "请求成功", result))
}
