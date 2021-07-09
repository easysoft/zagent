package handler

import (
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/biz/validate"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type RoleCtrl struct {
	RoleService *serverService.RoleService `inject:""`

	UserRepo *repo.UserRepo `inject:""`
	RoleRepo *repo.RoleRepo `inject:""`
	PermRepo *repo.PermRepo `inject:""`
}

func NewRoleCtrl() *RoleCtrl {
	return &RoleCtrl{}
}

/**
* @api {get} /admin/roles/:id 根据id获取角色信息
* @apiName 根据id获取角色信息
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 根据id获取角色信息
* @apiSampleRequest /admin/roles/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission
 */
func (c *RoleCtrl) GetRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	//id, _ := ctx.Params().GetUint("id")

	role, err := c.RoleRepo.GetRole(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	rr := c.RoleService.RoleTransform(role)
	rr.Perms = c.PermRepo.PermsTransform(c.RoleService.RolePermissions(role))
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", rr))
}

/**
* @api {post} /admin/roles/ 新建角色
* @apiName 新建角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 新建角色
* @apiSampleRequest /admin/roles/
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleCtrl) CreateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	role := new(model.Role)

	if err := ctx.ReadJSON(role); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*role)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, e, nil))
				return
			}
		}
	}

	err = c.RoleService.CreateRole(role)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}
	if role.ID == 0 {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, "操作失败", nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", c.RoleService.RoleTransform(role)))

}

/**
* @api {post} /admin/roles/:id/update 更新角色
* @apiName 更新角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 更新角色
* @apiSampleRequest /admin/roles/:id/update
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleCtrl) UpdateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	role := new(model.Role)
	if err := ctx.ReadJSON(role); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*role)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, e, nil))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	err = c.RoleService.UpdateRole(id, role)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", c.RoleService.RoleTransform(role)))

}

/**
* @api {delete} /admin/roles/:id/delete 删除角色
* @apiName 删除角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 删除角色
* @apiSampleRequest /admin/roles/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleCtrl) DeleteRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")

	err := c.RoleRepo.DeleteRoleById(id)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "删除成功", nil))
}

/**
* @api {get} /roles 获取所有的角色
* @apiName 获取所有的角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 获取所有的角色
* @apiSampleRequest /roles
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleCtrl) GetAllRoles(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	roles, count, err := c.RoleRepo.GetAllRoles(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	transform := c.RoleService.RolesTransform(roles)
	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功", map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}))
}
