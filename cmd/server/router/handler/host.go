package handler

import (
	"encoding/json"
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
	serverService "github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
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
// @Param task body v1.HostRegisterReq true "Host Object"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/host/register [post]
func (c *HostCtrl) Register(ctx iris.Context) {
	req := v1.HostRegisterReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	str, _ := json.Marshal(req)
	_logUtils.Infof("%v", str)

	success := c.AssertService.RegisterHost(req)
	code := _const.ResultFail
	if success {
		code = _const.ResultSuccess
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(code, "操作成功", ""))
	return
}
