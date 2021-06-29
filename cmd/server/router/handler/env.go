package handler

import (
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type EnvCtrl struct {
	BaseCtrl

	EnvironmentService *service.EnvironmentService `inject:""`
}

func NewEnvCtrl() *EnvCtrl {
	return &EnvCtrl{}
}

func (c *EnvCtrl) GetMap(ctx iris.Context) {
	model := model.Environment{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	data := c.EnvironmentService.GetMap(model)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", data))
}
