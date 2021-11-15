package handler

import (
	"encoding/json"
	v1 "github.com/easysoft/zagent/cmd/server/router/v1"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/biz/jwt"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"github.com/kataras/iris/v12"
)

type TaskCtrl struct {
	BaseCtrl

	TaskService      *serverService.TaskService      `inject:""`
	HistoryService   *serverService.HistoryService   `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewTaskCtrl() *TaskCtrl {
	return &TaskCtrl{}
}

// List
// @summary 列出测试任务
// @Produce json
// @Param progress query string false "progress: consts.BuildProgress"
// @Param status query string false "status: consts.BuildStatus"
// @Param keywords query string false "keywords"
// @Param pageNo query int false "page no"
// @Param pageSize query int false "page size"
// @Success 200 {object} _httpUtils.ResponsePage{data=[]model.Task} "code = success? 1 : 0"
// @Router /api/v1/client/task/list [get]
func (c *TaskCtrl) List(ctx iris.Context) {
	disabled := ctx.URLParam("disabled")
	keywords := ctx.URLParam("keywords")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	projects, total := c.TaskService.List(keywords, disabled, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(_const.ResultSuccess, "请求成功",
		projects, pageNo, pageSize, total))
}

// Get
// @summary 获取测试任务
// @Produce json
// @Param id path int true "Task Id"
// @Success 200 {object} _httpUtils.Response{data=model.Task} "code = success? 1 : 0"
// @Router /api/v1/client/task/{id} [get]
func (c *TaskCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	task := c.TaskService.GetDetail(uint(id))
	buildHistories := c.HistoryService.GetBuildHistoriesByTask(task.ID)

	mp := map[string]interface{}{"task": task, "buildHistories": buildHistories}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", mp))
	return
}

// Create
// @summary 创建测试任务
// @Accept json
// @Produce json
// @Param task body v1.TaskReq true "Task Object"
// @Success 200 {object} _httpUtils.Response{data=model.Task} "code = success? 1 : 0"
// @Router /api/v1/client/task/create [post]
func (c *TaskCtrl) Create(ctx iris.Context) {
	req := v1.TaskReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	jsn, _ := json.Marshal(req)
	_logUtils.Infof(string(jsn))

	if c.Validate(req, ctx) {
		return
	}

	cred := jwt.GetCredentials(ctx)
	userId := uint(0)
	if cred != nil {
		userId = _stringUtils.ParseUint(cred.UserId)
	}

	po, _ := req.ToModel()
	err := c.TaskService.Save(&po, userId)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", po))
	return
}

// Update
// @summary 更新测试任务
// @Accept json
// @Produce json
// @Param task body v1.TaskReq true "Task Object"
// @Success 200 {object} _httpUtils.Response{data=model.Task} "code = success? 1 : 0"
// @Router /api/v1/client/task/create [put]
func (c *TaskCtrl) Update(ctx iris.Context) {
	model := model.Task{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	err := c.TaskService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", model))
}

// Delete
// @summary 删除测试任务
// @Accept json
// @Produce json
// @Param id path int true "Task Id"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/task/{id} [delete]
func (c *TaskCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	c.TaskService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", ""))
}

func (c *TaskCtrl) TestWs(ctx iris.Context) {
	data := map[string]interface{}{"action": serverConst.TaskUpdate, "taskId": 1, "msg": ""}
	c.WebSocketService.Broadcast(serverConst.WsNamespace, serverConst.WsDefaultRoom, serverConst.WsEvent, data)

	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", ""))
}
