package hostHandler

import (
	hostAgentService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type TaskCtrl struct {
	TaskService *hostAgentService.TaskService `inject:""`
}

func NewTaskCtrl() *TaskCtrl {
	return &TaskCtrl{}
}

// @summary 获取任务状态
// @Accept json
// @Produce json
// @Success 200 {object} _domain.Response{data=v1.ListTaskResp} "code = success | fail"
// @Router /api/v1/task/status [get]
func (c *TaskCtrl) GetStatus(ctx iris.Context) {
	data, _ := c.TaskService.ListTask()

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", data))

	return
}
