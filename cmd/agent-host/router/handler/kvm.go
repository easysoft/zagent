package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/agent-host/router/v1"
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	"github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
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

// Create
// @summary 创建KVM虚拟机
// @Accept json
// @Produce json
// @Param task body v1.KvmReq true "Kvm Request Object"
// @Success 200 {object} _httpUtils.Response{data=domain.Vm} "code = success? 1 : 0"
// @Router /api/v1/kvm/create [post]
func (c *KvmCtrl) Create(ctx iris.Context) {
	req := v1.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	dom, vmVncPort, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)

	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "fail to create vm", err))
		return
	}

	vmName, _ := dom.GetName()
	vm := domain.Vm{
		Name:        vmName,
		VncPort:     strconv.Itoa(vmVncPort),
		ImagePath:   vmRawPath,
		BackingPath: vmBackingPath,
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to create vm", vm))

	return
}

// Destroy
// @summary 摧毁KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _httpUtils.Response{} "code = success? 1 : 0"
// @Router /api/v1/kvm/{name}/destroy [post]
func (c *KvmCtrl) Destroy(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.DestroyVmByName(name, true)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to destroy vm", name))
	return
}

// Reboot
// @summary 重启KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _httpUtils.Response{} "code = success? 1 : 0"
// @Router /api/v1/kvm/{name}/reboot [post]
func (c *KvmCtrl) Reboot(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.RebootVmByName(name)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to reboot vm", name))
	return
}

// Suspend
// @summary 暂停KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _httpUtils.Response{} "code = success? 1 : 0"
// @Router /api/v1/kvm/{name}/suspend [post]
func (c *KvmCtrl) Suspend(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.SuspendVmByName(name)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to suspend vm", name))
	return
}

// Resume
// @summary 恢复KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _httpUtils.Response{} "code = success? 1 : 0"
// @Router /api/v1/kvm/{name}/resume [post]
func (c *KvmCtrl) Resume(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.ResumeVmByName(name)

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to resume vm", name))
	return
}

//func (c *KvmCtrl) Boot(ctx iris.Context) {
//	req :=v1.KvmReq{}
//	if err := ctx.ReadJSON(&req); err != nil {
//		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
//		return
//	}
//
//	c.LibvirtService.BootVmByName(req.VmUniqueName)
//
//	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to boot vm", req.VmUniqueName))
//	return
//}
//func (c *KvmCtrl) Shutdown(ctx iris.Context) {
//	req :=v1.KvmReq{}
//	if err := ctx.ReadJSON(&req); err != nil {
//		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
//		return
//	}
//
//	c.LibvirtService.ShutdownVmByName(req.VmUniqueName)
//
//	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to shutdown vm", req.VmUniqueName))
//	return
//}
