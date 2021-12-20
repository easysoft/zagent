package handler

import (
	"github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/lib/http"
	"github.com/easysoft/zv/internal/server/service"
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
