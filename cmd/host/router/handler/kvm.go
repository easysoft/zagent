package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	natHelper "github.com/easysoft/zv/internal/agent/utils/nat"
	consts "github.com/easysoft/zv/internal/comm/const"
	hostKvmService "github.com/easysoft/zv/internal/host/service/kvm"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type KvmCtrl struct {
	VmService      *hostKvmService.KvmService     `inject:""`
	LibvirtService *hostKvmService.LibvirtService `inject:""`
}

func NewKvmCtrl() *KvmCtrl {
	return &KvmCtrl{}
}

// ListTmpl
// @summary 获取KVM虚拟机模板信息
// @Produce json
// @Success 200 {object} _httpUtils.Response{data=[]v1.KvmRespTempl} "code = success? 1 : 0"
// @Router /api/v1/kvm/listTempl [get]
func (c *KvmCtrl) ListTmpl(ctx iris.Context) {
	templs, err := c.LibvirtService.ListTmpl()

	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "fail to list vm tmpl", err))
		return
	}

	ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "success to list vm tmpl", templs))

	return
}

// Create
// @summary 创建KVM虚拟机
// @Accept json
// @Produce json
// @Param kvmReq body v1.KvmReq true "Kvm Request Object"
// @Success 200 {object} _httpUtils.Response{data=v1.KvmResp} "code = success? 1 : 0"
// @Router /api/v1/kvm/create [post]
func (c *KvmCtrl) Create(ctx iris.Context) {
	req := v1.KvmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	dom, vmVncPort, vmAgentPortMapped, vmRawPath, vmBackingPath, err := c.LibvirtService.CreateVm(&req, true)

	vmName := req.VmUniqueName
	vmStatus := consts.VmLaunch
	if err != nil || dom == nil {
		vmStatus = consts.VmFailCreate
	} else {
		vmName, _ = dom.GetName()
	}

	vm := v1.KvmResp{
		Name:        vmName,
		MacAddress:  req.VmMacAddress,
		AgentPort:   vmAgentPortMapped,
		VncPort:     vmVncPort,
		ImagePath:   vmRawPath,
		BackingPath: vmBackingPath,
		Status:      vmStatus,
	}

	ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "success to create vm", vm))

	return
}

// Clone
// @summary 克隆KVM虚拟机
// @Accept json
// @Produce json
// @Param kvmReqClone body v1.KvmReqClone true "Kvm Request Object"
// @Success 200 {object} _httpUtils.Response{data=v1.KvmResp} "code = success? 1 : 0"
// @Router /api/v1/kvm/clone [post]
func (c *KvmCtrl) Clone(ctx iris.Context) {
	req := v1.KvmReqClone{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	if req.VmSrc == "" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "request vmSrc field can not be empty.", nil))
		return
	}

	dom, vmIp, vmVncPort, vmAgentPortMapped, vmRawPath, vmBackingPath, err := c.LibvirtService.CloneVm(&req, true)

	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "fail to clone vm", err))
		return
	}

	vmName, _ := dom.GetName()
	vm := v1.KvmResp{
		Name:        vmName,
		IpAddress:   vmIp,
		MacAddress:  req.VmMacAddress,
		AgentPort:   vmAgentPortMapped,
		VncPort:     vmVncPort,
		ImagePath:   vmRawPath,
		BackingPath: vmBackingPath,
	}

	ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "success to create vm", vm))

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

	req := v1.DestroyVmReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	err = c.LibvirtService.DestroyVmByName(name, true)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	natHelper.RemoveForward(req.Ip, 0)

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

	err := c.LibvirtService.SuspendVmByName(name)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

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

	err := c.LibvirtService.ResumeVmByName(name)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to resume vm", name))
	return
}

//func (c *KvmCtrl) Boot(ctx iris.Context) {
//	req :=v1.KvmReq{}
//	if err := ctx.ReadJSON(&req); err != nil {
//		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
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
//		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
//		return
//	}
//
//	c.LibvirtService.ShutdownVmByName(req.VmUniqueName)
//
//	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to shutdown vm", req.VmUniqueName))
//	return
//}
