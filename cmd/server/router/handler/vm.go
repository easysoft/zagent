package handler

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
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

func (c *VmCtrl) Register(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := commDomain.Vm{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	rpcResp := c.AssertService.RegisterVm(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(int64(rpcResp.Code), "操作成功", rpcResp))
	return
}
