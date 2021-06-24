package handler

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ValidCtrl struct {
	ValidService *service.ValidService `inject:""`
}

func NewValidCtrl() *ValidCtrl {
	return &ValidCtrl{}
}

func (c *ValidCtrl) Valid(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := commDomain.ValidRequest{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	result := c.ValidService.Valid(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", result))
}
