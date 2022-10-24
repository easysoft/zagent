package hostHandler

import (
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/host/service/virtualbox"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type VirtualBoxCtrl struct {
	VirtualBoxService *virtualboxService.VirtualBoxService `inject:""`
}

func NewVirtualBoxCtrl() *VirtualBoxCtrl {
	return &VirtualBoxCtrl{}
}

func (c *VirtualBoxCtrl) Create(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	result, err := c.VirtualBoxService.Create(req)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail,
			fmt.Sprintf("fail to create virtualbox vm, reason: %s.", err.Error()), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to create vm", result))

	return
}

func (c *VirtualBoxCtrl) Destroy(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.VirtualBoxService.Destroy(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to destroy vm", req.VmUniqueName))
	return
}

func (c *VirtualBoxCtrl) ListTmpl(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	result, err := c.VirtualBoxService.ListTmpl(req)

	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, "fail to list vm tmpl", err))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to list vm tmpl", result))

	return
}
