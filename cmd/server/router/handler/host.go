package handler

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type HostCtrl struct {
	BaseCtrl

	HostService *serverService.HostService `inject:""`
}

func NewHostCtrl() *HostCtrl {
	return &HostCtrl{}
}

func (c *HostCtrl) Register(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := commDomain.Host{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	rpcResp := c.HostService.Register(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", rpcResp))
	return
}
