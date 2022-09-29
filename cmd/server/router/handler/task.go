package handler

import (
	"encoding/json"
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
	"github.com/easysoft/zv/internal/server/biz/jwt"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/service"
	commonService "github.com/easysoft/zv/internal/server/service/common"
	serverConst "github.com/easysoft/zv/internal/server/utils/const"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
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

func (c *TaskCtrl) List(ctx iris.Context) {
	disabled := ctx.URLParam("disabled")
	keywords := ctx.URLParam("keywords")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	projects, total := c.TaskService.List(keywords, disabled, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.RespDataPagination(_const.ResultPass, "请求成功",
		projects, pageNo, pageSize, total))
}

func (c *TaskCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	task := c.TaskService.GetDetail(uint(id))
	buildHistories := c.HistoryService.GetBuildHistoriesByTask(task.ID)

	result := v1.TaskResp{Task: task, BuildHistories: buildHistories}
	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "操作成功", result))
	return
}

func (c *TaskCtrl) Create(ctx iris.Context) {
	req := v1.TaskReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
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
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "操作成功", po))
	return
}

func (c *TaskCtrl) Update(ctx iris.Context) {
	model := model.Task{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	err := c.TaskService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "操作成功", model))
}

func (c *TaskCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	c.TaskService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "操作成功", ""))
}

func (c *TaskCtrl) TestWs(ctx iris.Context) {
	data := map[string]interface{}{"action": serverConst.TaskUpdate, "taskId": 1, "msg": ""}
	c.WebSocketService.Broadcast(serverConst.WsNamespace, serverConst.WsDefaultRoom, serverConst.WsEvent, data)

	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "操作成功", ""))
}
