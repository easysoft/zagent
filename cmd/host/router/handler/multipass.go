package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	multiPassService "github.com/easysoft/zagent/internal/host/service/multipass"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type MultiPassCtrl struct {
	MultiPassService *multiPassService.MultiPassService `inject:""`
}

func NewMultiPassCtrl() *MultiPassCtrl {
	return &MultiPassCtrl{}
}

func (c *MultiPassCtrl) List(ctx iris.Context) {
	domains, err := c.MultiPassService.GetVmList()
	if domains == nil || err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultPass, "fail to find MultiPass vm", err))
		return
	}
	var mps []v1.MultiPassResp
	mp := v1.MultiPassResp{}
	for _, v := range domains {
		mp.Name = v.Name
		mp.Memory = v.Memory
		mp.Cpus = v.Cpus
		mp.IPv4 = v.IPv4
		mp.State = v.State
		mps = append(mps, mp)
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to get MultiPass vm", mps))

	return
}

func (c *MultiPassCtrl) Create(ctx iris.Context) {
	req := v1.MultiPassReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}
	domains, err := c.MultiPassService.CreateVm(&req, false)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, "fail to create vm", err))
		return
	}

	vm := v1.MultiPassResp{
		Name: domains.Name, //Vnc:  strconv.Itoa(MultiPassVncPort),
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to create MultiPass vm", vm))

	return
}

func (c *MultiPassCtrl) Destroy(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.MultiPassService.DestroyVm(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to destroy MultiPass vm", name))
	return
}

func (c *MultiPassCtrl) Reboot(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.MultiPassService.RebootVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to reboot vm", name))
	return
}

func (c *MultiPassCtrl) Suspend(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.MultiPassService.SuspendVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to suspend vm", name))
	return
}

func (c *MultiPassCtrl) Resume(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.MultiPassService.ResumeVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to resume vm", name))
	return
}

func (c *MultiPassCtrl) GetToken(ctx iris.Context) {
	port := ctx.URLParam("port")

	if port == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "no port param", nil))
		return
	}

	ret := c.MultiPassService.GetToken(port)
	if ret.Token == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "token not found", nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", ret))

	return
}
