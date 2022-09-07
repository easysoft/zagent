package hostHandler

import (
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/host/service/virtualbox"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type VirtualBoxCtrl struct {
	VirtualBoxService *virtualboxService.VirtualBoxService `inject:""`
}

func NewVirtualBoxCtrl() *VirtualBoxCtrl {
	return &VirtualBoxCtrl{}
}

// Create
// @summary 创建VirtualBox虚拟机
// @Accept json
// @Produce json
// @Param virtualboxReq body v1.VirtualBoxReq true "VirtualBox Request Object"
// @Success 200 {object} _httpUtils.Response{data=v1.VirtualBoxResp} "code = success? 1 : 0"
// @Router /api/v1/virtualbox/create [post]
func (c *VirtualBoxCtrl) Create(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	result, err := c.VirtualBoxService.Create(req)
	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError,
			fmt.Sprintf("fail to create virtualbox vm, reason %s.", err.Error()), nil))
		return
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to create vm", result))

	return
}

// Destroy
// @summary 摧毁VirtualBox虚拟机
// @Accept json
// @Produce json
// @Param name path string true "VirtualBox Name"
// @Success 200 {object} _httpUtils.Response{} "code = success? 1 : 0"
// @Router /api/v1/virtualbox/{name}/destroy [post]
func (c *VirtualBoxCtrl) Destroy(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.VirtualBoxService.Destroy(req)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to destroy vm", req.VmUniqueName))
	return
}

// ListTmpl
// @summary 获取VirtualBox虚拟机模板信息
// @Produce json
// @Success 200 {object} _httpUtils.Response{data=[]v1.KvmRespTempl} "code = success? 1 : 0"
// @Router /api/v1/virtualbox/listTempl [post]
func (c *VirtualBoxCtrl) ListTmpl(ctx iris.Context) {
	req := v1.VirtualBoxReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	result, err := c.VirtualBoxService.ListTmpl(req)

	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "fail to list vm tmpl", err))
		return
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to list vm tmpl", result))

	return
}
