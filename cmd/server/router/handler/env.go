package handler

import (
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type EnvCtrl struct {
	BaseCtrl

	EnvService *service.EnvService `inject:""`
}

func NewEnvCtrl() *EnvCtrl {
	return &EnvCtrl{}
}

func (c *EnvCtrl) GetMap(ctx iris.Context) {
	data, _ := c.EnvService.GetMap()
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", data))
}
