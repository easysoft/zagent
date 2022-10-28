package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	vmWareService "github.com/easysoft/zv/internal/host/service/vmware"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type VmWareCtrl struct {
	VmWareService *vmWareService.VmWareService `inject:""`
}

func NewVmWareCtrl() *VmWareCtrl {
	return &VmWareCtrl{}
}

func (c *VmWareCtrl) Create(ctx iris.Context) {
	req := v1.VmWareReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	id, macAddress, err := c.VmWareService.CreateVm(&req, true)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultPass, "fail to create VmWare vm", err))
		return
	}

	vm := v1.VmWareResp{
		VmId: id,
		Name: req.VmUniqueName,
		Mac:  macAddress,
		//Vnc:  strconv.Itoa(VmWareVncPort),
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to create VmWare vm", vm))

	return
}

func (c *VmWareCtrl) Destroy(ctx iris.Context) {
	req := v1.VmWareReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.VmWareService.DestroyVm(&req, true)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to destroy VmWare vm", req.VmId))
	return
}
