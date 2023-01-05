package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	hostAgentService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type SnapCtrl struct {
	TaskService *hostAgentService.TaskService `inject:""`

	SnapService *hostAgentService.SnapService `inject:""`
}

func NewSnapCtrl() *SnapCtrl {
	return &SnapCtrl{}
}

// @summary 添加创建快照任务
// @Accept json
// @Produce json
// @Param SnapTaskReq body []v1.SnapTaskReq true "Snap Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/snaps/add [post]
func (c *SnapCtrl) Add(ctx iris.Context) {
	req := make([]v1.SnapTaskReq, 0)
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err = c.SnapService.AddTasks(req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to add create snapshot task", nil))
	return
}

// @summary 强制终止新建快照任务
// @Accept json
// @Produce json
// @Param SnapCancelReq body v1.SnapCancelReq true "CancelDate Snap Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/snaps/cancel [post]
func (c *SnapCtrl) Cancel(ctx iris.Context) {
	req := v1.SnapCancelReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	c.SnapService.CancelTask(uint(req.Id))

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", nil))
	return
}

func (c *SnapCtrl) ListSnap(ctx iris.Context) {
	req := v1.SnapTaskReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	snaps, err := c.SnapService.ListSnap(req.Vm)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to create snapshot", snaps))

	return
}

func (c *SnapCtrl) RemoveSnap(ctx iris.Context) {
	req := v1.SnapTaskReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err = c.SnapService.RemoveSnap(&req)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to remove snapshot", nil))

	return
}

func (c *SnapCtrl) RevertSnap(ctx iris.Context) {
	req := v1.SnapTaskReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err = c.SnapService.RevertSnap(&req)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to remove snapshot", nil))

	return
}
