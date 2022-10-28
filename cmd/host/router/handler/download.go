package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	hostAgentService "github.com/easysoft/zv/internal/host/service"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type DownloadCtrl struct {
	DownloadService *hostAgentService.DownloadService `inject:""`
}

func NewDownloadCtrl() *DownloadCtrl {
	return &DownloadCtrl{}
}

func (c *DownloadCtrl) ListTask(ctx iris.Context) {
	data, _ := c.DownloadService.ListTask()

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", data))

	return
}

func (c *DownloadCtrl) AddTasks(ctx iris.Context) {
	req := v1.DownloadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.DownloadService.AddTasks(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", nil))
	return
}

func (c *DownloadCtrl) CancelTask(ctx iris.Context) {
	req := v1.DownloadCancelReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.DownloadService.CancelTask(req.Url)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", nil))
	return
}
