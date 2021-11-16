package handler

import (
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/lib/convertor"
	"github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/biz/jwt"
	"github.com/easysoft/zagent/internal/server/biz/transformer"
	"github.com/easysoft/zagent/internal/server/biz/validate"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type UserCtrl struct {
	UserService *serverService.UserService `inject:""`
	RoleService *serverService.RoleService `inject:""`
	UserRepo    *repo.UserRepo             `inject:""`
	RoleRepo    *repo.RoleRepo             `inject:""`
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

/**
* @api {get} /admin/profile 获取登陆用户信息
* @apiName 获取登陆用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取登陆用户信息
* @apiSampleRequest /admin/profile
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserCtrl) GetProfile(ctx iris.Context) {
	cred := jwt.GetCredentials(ctx)
	if cred == nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(401, "not login", nil))
		return
	}

	idInt, _ := strconv.Atoi(cred.UserId)
	s := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     uint(idInt),
			},
		},
	}
	user, err := c.UserRepo.GetUser(s)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "请求成功", c.userTransform(user)))
}

func (c *UserCtrl) GetAdminInfo(ctx iris.Context) {
	user, err := c.UserRepo.GetUser(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "请求成功", map[string]string{"avatar": user.Avatar}))
}

func (c *UserCtrl) ChangeAvatar(ctx iris.Context) {
	sess := jwt.GetCredentials(ctx)
	idInt, _ := strconv.Atoi(sess.UserId)
	id := uint(idInt)

	avatar := new(model.Avatar)
	if err := ctx.ReadJSON(avatar); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*avatar)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, e, nil))
				return
			}
		}
	}

	user := c.UserRepo.NewUser()
	user.ID = id
	user.Avatar = avatar.Avatar
	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "请求成功", c.userTransform(user)))
}

/**
* @api {get} /admin/users/:id 根据id获取用户信息
* @apiName 根据id获取用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 根据id获取用户信息
* @apiSampleRequest /admin/users/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserCtrl) GetUser(ctx iris.Context) {
	//id, _ := ctx.Params().GetUint("id")

	user, err := c.UserRepo.GetUser(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", c.userTransform(user)))
}

/**
* @api {post} /admin/users/ 新建账号
* @apiName 新建账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 新建账号
* @apiSampleRequest /admin/users/
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserCtrl) CreateUser(ctx iris.Context) {
	user := new(model.User)
	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, e, nil))
				return
			}
		}
	}

	err = c.UserService.CreateUser(user)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}

	if user.ID == 0 {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "操作失败", nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", c.userTransform(user)))
	return

}

/**
* @api {post} /admin/users/:id/update 更新账号
* @apiName 更新账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 更新账号
* @apiSampleRequest /admin/users/:id/update
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserCtrl) UpdateUser(ctx iris.Context) {
	user := new(model.User)

	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, e, nil))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	if user.Username == "username" {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, "不能编辑管理员", nil))
		return
	}

	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", c.userTransform(user)))
}

/**
* @api {delete} /admin/users/:id/delete 删除用户
* @apiName 删除用户
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 删除用户
* @apiSampleRequest /admin/users/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserCtrl) DeleteUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	err := c.UserRepo.DeleteUser(id)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "删除成功", nil))
}

/**
* @api {get} /users 获取所有的账号
* @apiName 获取所有的账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取所有的账号
* @apiSampleRequest /users
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserCtrl) GetAllUsers(ctx iris.Context) {
	//name := ctx.FormValue("name")

	users, count, err := c.UserRepo.GetAllUsers(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusUnauthorized, err.Error(), nil))
		return
	}

	transform := c.usersTransform(users)
	_, _ = ctx.JSON(_httpUtils.ApiRes(_const.ResultSuccess, "操作成功", map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}))

}

func (c *UserCtrl) usersTransform(users []*model.User) []*transformer.User {
	var us []*transformer.User
	for _, user := range users {
		u := c.userTransform(user)
		us = append(us, u)
	}
	return us
}

func (c *UserCtrl) userTransform(user *model.User) *transformer.User {
	u := &transformer.User{}
	g := _convertor.NewTransform(u, user, time.RFC3339)
	_ = g.Transformer()

	roleIds := c.RoleService.GetRolesForUser(user.ID)
	var ris []int
	for _, roleId := range roleIds {
		ri, _ := strconv.Atoi(roleId)
		ris = append(ris, ri)
	}
	roles, _, err := c.RoleRepo.GetAllRoles(nil)
	if err == nil {
		u.Roles = c.RoleService.RolesTransform(roles)
	}
	return u
}
