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

func (c *DownloadCtrl) Download(ctx iris.Context) {
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
