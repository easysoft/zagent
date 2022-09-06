package handler

import (
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type EnvironmentCtrl struct {
	BaseCtrl

	EnvironmentService *serverService.EnvironmentService `inject:""`
}

func NewEnvironmentCtrl() *EnvironmentCtrl {
	return &EnvironmentCtrl{}
}

func (c *EnvironmentCtrl) GetData(ctx iris.Context) {
	model := model.Environment{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	data := c.EnvironmentService.GetMap(model)
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "请求成功", data))
}
