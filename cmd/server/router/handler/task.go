package handler

import (
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
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
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.TaskService.GetDetail(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", model))
	return
}

func (c *TaskCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.Task{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	cred := jwt.GetCredentials(ctx)

	err := c.TaskService.Save(&model, _stringUtils.ParseUint(cred.UserId))
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", model))
	return
}

func (c *TaskCtrl) Update(ctx iris.Context) {
	model := model.Task{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.TaskService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", model))
}

func (c *TaskCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.TaskService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", ""))
}

func (c *TaskCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
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
