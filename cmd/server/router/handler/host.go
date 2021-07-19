package handler

import (
	"github.com/easysoft/zagent/internal/comm/domain"
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

func (c *HostCtrl) Register(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := domain.HostNode{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	rpcResp := c.AssertService.RegisterHost(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(int64(rpcResp.Code), "操作成功", rpcResp))
	return
}
