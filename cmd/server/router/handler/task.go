package handler

import (
	"encoding/json"
	v1 "github.com/easysoft/zagent/cmd/server/router/v1"
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

func (c *TaskCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	projects, total := c.TaskService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(iris.StatusOK, "请求成功",
		projects, pageNo, pageSize, total))
}

func (c *TaskCtrl) ListForSelect(ctx iris.Context) {
	projects, _ := c.TaskService.List("", "", 0, 0)

	data := map[string]interface{}{}
	data["projects"] = projects

	projectId := jwt.Get(ctx, "projectId")
	data["projects"] = projects
	data["projectId"] = projectId

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "请求成功", data))
}

func (c *TaskCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	task := c.TaskService.GetDetail(uint(id))
	buildHistories := c.HistoryService.GetBuildHistoriesByTask(task.ID)

	mp := map[string]interface{}{"task": task, "buildHistories": buildHistories}
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", mp))
	return
}

// Create
// @summary 创建测试任务
// @Description
// @Accept json
// @Produce json
// @Param task body v1.Task true "task object"
// Success 200 {object} _httpUtils.Response
// Failure 500 {object} _httpUtils.Response
// @Router /api/v1/client/task/create [post]
func (c *TaskCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	req := v1.Task{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
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
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", po))
	return
}

func (c *TaskCtrl) Update(ctx iris.Context) {
	model := model.Task{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	err := c.TaskService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", model))
}

func (c *TaskCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.TaskService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", ""))
}

func (c *TaskCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.TaskService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", ""))
}

func (c *TaskCtrl) TestWs(ctx iris.Context) {
	data := map[string]interface{}{"action": serverConst.TaskUpdate, "taskId": 1, "msg": ""}
	c.WebSocketService.Broadcast(serverConst.WsNamespace, serverConst.WsDefaultRoom, serverConst.WsEvent, data)

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", ""))
}
