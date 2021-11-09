package hostHandler

import (
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	"github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"strconv"
)

type KvmCtrl struct {
	VmService      *hostKvmService.VmService      `inject:""`
	LibvirtService *hostKvmService.LibvirtService `inject:""`
}

func NewKvmCtrl() *KvmCtrl {
	return &KvmCtrl{}
}

func (c *KvmCtrl) Create(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	dom, vmVncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)

	if err == nil {
		vmName, _ := dom.GetName()
		vm := domain.Vm{
			Name:        vmName,
			VncAddress:  strconv.Itoa(vmVncPort),
			ImagePath:   vmRawPath,
			BackingPath: vmBackingPath,
		}

		ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to create vm", vm))
	} else {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "fail to create vm", err))
	}

	return
}

func (c *KvmCtrl) Destroy(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.DestroyVmByName(req.VmUniqueName, true)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to destroy vm", req.VmUniqueName))
	return
}

func (c *KvmCtrl) Boot(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.BootVmByName(req.VmUniqueName)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to boot vm", req.VmUniqueName))
	return
}
func (c *KvmCtrl) Shutdown(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.ShutdownVmByName(req.VmUniqueName)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to shutdown vm", req.VmUniqueName))
	return
}
func (c *KvmCtrl) Reboot(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.RebootVmByName(req.VmUniqueName)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to reboot vm", req.VmUniqueName))
	return
}

func (c *KvmCtrl) Suspend(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.SuspendVmByName(req.VmUniqueName)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to suspend vm", req.VmUniqueName))
	return
}
func (c *KvmCtrl) Resume(ctx iris.Context) {
	req := domain.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.LibvirtService.ResumeVmByName(req.VmUniqueName)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to resume vm", req.VmUniqueName))
	return
}
