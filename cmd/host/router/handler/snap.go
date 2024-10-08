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
// @Router /api/v1/snaps/addCreateSnap [post]
func (c *SnapCtrl) AddCreateSnap(ctx iris.Context) {
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

// @summary 添加回滚到快照任务
// @Accept json
// @Produce json
// @Param SnapTaskReq body []v1.SnapTaskReq true "Snap Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/snaps/addRevertSnap [post]
func (c *SnapCtrl) AddRevertSnap(ctx iris.Context) {
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

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to add revert snapshot task", nil))
	return
}

// @summary 列出虚拟机快照任务
// @Accept json
// @Produce json
// @Param vm query string true "vm name"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/snaps/listSnap [get]
func (c *SnapCtrl) ListSnap(ctx iris.Context) {
	vm := ctx.URLParam("vm")

	snaps, err := c.SnapService.ListSnap(vm)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "", snaps))

	return
}

// @summary 移除虚拟机快照任务
// @Accept json
// @Produce json
// @Param SnapTaskReq body []v1.SnapTaskReq true "Snap Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/snaps/removeSnap [post]
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
