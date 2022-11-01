package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	hostAgentService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type DownloadCtrl struct {
	TaskService *hostAgentService.TaskService `inject:""`

	DownloadService *hostAgentService.DownloadService `inject:""`
}

func NewDownloadCtrl() *DownloadCtrl {
	return &DownloadCtrl{}
}

// @summary 添加下载任务
// @Accept json
// @Produce json
// @Param DownloadReq body v1.DownloadReq true "Download Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/download/add [post]
func (c *DownloadCtrl) Add(ctx iris.Context) {
	req := make([]v1.DownloadReq, 0)
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.DownloadService.AddTasks(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", nil))
	return
}

// @summary 强制终止下载任务
// @Accept json
// @Produce json
// @Param DownloadCancelReq body v1.DownloadCancelReq true "Cancel Download Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/download/cancel [post]
func (c *DownloadCtrl) Cancel(ctx iris.Context) {
	req := v1.DownloadCancelReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.DownloadService.CancelTask(uint(req.Id))

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", nil))
	return
}
