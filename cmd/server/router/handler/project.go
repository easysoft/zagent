package handler

import (
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/biz/jwt"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/service"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"github.com/kataras/iris/v12"
)

type ProjectCtrl struct {
	BaseCtrl

	ProjectService *service.ProjectService `inject:""`
}

func NewProjectCtrl() *ProjectCtrl {
	return &ProjectCtrl{}
}

func (c *ProjectCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	projects, total := c.ProjectService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		projects, pageNo, pageSize, total))
}

func (c *ProjectCtrl) ListForSelect(ctx iris.Context) {
	projects, _ := c.ProjectService.List("", "", 0, 0)

	data := map[string]interface{}{}
	data["projects"] = projects

	projectId := jwt.Get(ctx, "projectId")
	data["projects"] = projects
	data["projectId"] = projectId

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", data))
}

func (c *ProjectCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.ProjectService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *ProjectCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.Project{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	cred := jwt.GetCredentials(ctx)

	err := c.ProjectService.Save(&model, _stringUtils.ParseUint(cred.UserId))
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *ProjectCtrl) Update(ctx iris.Context) {
	model := model.Project{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.ProjectService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *ProjectCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.ProjectService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *ProjectCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.ProjectService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *ProjectCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.ProjectService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}
