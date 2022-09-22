package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	kvmService "github.com/easysoft/zv/internal/host/service/kvm"
	virtualService "github.com/easysoft/zv/internal/host/service/virtual"
	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	"github.com/easysoft/zv/internal/pkg/domain"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	vmMacMap = sync.Map{}
)

type VirtualCtrl struct {
	SetupService *virtualService.SetupService `inject:""`
	KvmService   *kvmService.KvmService       `inject:""`
}

func NewVirtualCtrl() *VirtualCtrl {
	return &VirtualCtrl{}
}

func (c *VirtualCtrl) NotifyHost(ctx iris.Context) {
	req := domain.VmNotifyReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	data := domain.VmNotifyResp{
		Secret: agentConf.Inst.Secret,
	}

	// get vm ip
	vmIp, err := c.KvmService.GetVmIpByMac(req.MacAddress)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	// map vm agent port to host
	//vmAgentPortMapped, err := natHelper.GetValidPort()
	//if err != nil {
	//	ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
	//	return
	//}
	//err = natHelper.ForwardPort(vmIp, consts.AgentServicePost, vmAgentPortMapped, consts.Http)
	//if err != nil {
	//	ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
	//	return
	//}
	//
	//data.AgentPortOnHost = vmAgentPortMapped
	data.Ip = vmIp

	ctx.JSON(_httpUtils.RespData(_const.ResultPass, "success to refresh secret", data))
	return
}

func (c *VirtualCtrl) AddVmPortMap(ctx iris.Context) {
	req := v1.VmPortMapReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	resp, err := c.SetupService.AddVmPortMap(req)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(_const.ResultPass, "success", resp))
}

func (c *VirtualCtrl) RemoveVmPortMap(ctx iris.Context) {
	req := v1.VmPortMapReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	err := c.SetupService.RemoveVmPortMap(req)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(_const.ResultPass, "success", nil))
}

// GetToken
// @summary 根据VNC Port获取Token
// @Accept json
// @Produce json
// @Param port query string true "Virtual Port"
// @Success 200 {object} _httpUtils.Response{iris.Map} "code = success? 1 : 0"
// @Router /api/v1/vnc/getToken [get]
func (c *VirtualCtrl) GetToken(ctx iris.Context) {
	port := ctx.URLParam("port")

	if port == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, "no port param", nil))
		return
	}

	ret := c.SetupService.GetToken(port)
	if ret.Token == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, "token not found", nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(_const.ResultPass, "success", ret))

	return
}
