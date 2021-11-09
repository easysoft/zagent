package hostHandler

import (
	vmWareService "github.com/easysoft/zagent/internal/agent-host/service/vmware"
	"github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type VmWareCtrl struct {
	VmWareService *vmWareService.VmWareService `inject:""`
}

func NewVmWareCtrl() *VmWareCtrl {
	return &VmWareCtrl{}
}

func (c *VmWareCtrl) Create(ctx iris.Context) {
	req := domain.VmWareReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	id, macAddress, err := c.VmWareService.CreateVm(&req, true)
	if err == nil {
		vm := domain.VmWareResp{
			VmId: id,
			Name: req.VmUniqueName,
			Mac:  macAddress,
			//VncAddress:  strconv.Itoa(VmWareVncPort),
		}

		ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to create VmWare vm", vm))
	} else {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "fail to create VmWare vm", err))
	}

	return
}

func (c *VmWareCtrl) Destroy(ctx iris.Context) {
	req := domain.VmWareReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.VmWareService.DestroyVm(&req, true)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to destroy VmWare vm", req.VmId))
	return
}
