package handler

import (
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type VmCtrl struct {
	BaseCtrl

	AssertService *serverService.AssertService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

// Register
// @summary 向服务器注册虚拟机
// @Accept json
// @Produce json
// @Param task body domain.Vm true "Vm Object"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/host/register [post]
func (c *VmCtrl) Register(ctx iris.Context) {
	model := domain.Vm{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	success := c.AssertService.RegisterVm(model)
	code := _const.ResultFail
	if success {
		code = _const.ResultSuccess
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(code, "操作成功", ""))
	return
}
