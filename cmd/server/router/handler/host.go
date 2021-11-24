package handler

import (
	v1 "github.com/easysoft/zagent/cmd/server/router/v1"
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

// Register
// @summary 向服务器注册宿主机
// @Accept json
// @Produce json
// @Param task body v1.HostReq true "Host Object"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/host/register [post]
func (c *HostCtrl) Register(ctx iris.Context) {
	model := v1.HostReq{}
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
