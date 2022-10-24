package handler

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/service"
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
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	data := c.EnvironmentService.GetMap(model)
	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "请求成功", data))
}
