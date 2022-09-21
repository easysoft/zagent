package handler

import (
	"github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/biz/const"
	jwt2 "github.com/easysoft/zv/internal/server/biz/jwt"
	"github.com/easysoft/zv/internal/server/biz/redis"
	"github.com/easysoft/zv/internal/server/biz/validate"
	"github.com/easysoft/zv/internal/server/conf"
	"github.com/easysoft/zv/internal/server/repo"
	"github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type AccountCtrl struct {
	UserService *serverService.UserService `inject:""`

	UserRepo  *repo.UserRepo  `inject:""`
	TokenRepo *repo.TokenRepo `inject:""`
	RoleRepo  *repo.RoleRepo  `inject:""`
	PermRepo  *repo.PermRepo  `inject:""`
}

func NewAccountCtrl() *AccountCtrl {
	return &AccountCtrl{}
}

/**
* @api {post} /admin/login 用户登陆
* @apiName 用户登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户登陆
* @apiSampleRequest /admin/login
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *AccountCtrl) UserLogin(ctx iris.Context) {
	req := new(validate.LoginRequest)

	if err := ctx.ReadJSON(req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, e, nil))
				return
			}
		}
	}

	ctx.Application().Logger().Infof("%s 登录系统", req.Username)

	search := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     req.Username,
			},
		},
	}
	user, err := c.UserRepo.GetUser(search)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	response, success, msg := c.UserService.CheckLogin(ctx, user, req.Password)
	if success {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, msg, response))
		return
	}
	response.RememberMe = req.RememberMe

	refreshToken := ""
	if success && req.RememberMe {
		refreshToken = response.Token
	}

	c.UserService.UpdateRefreshToken(user.ID, refreshToken)

	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, msg, response))
}

func (c *AccountCtrl) UserLogout(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)

	var (
		credentials *bizConst.UserCredentials
		err         error
	)
	if serverConf.Inst.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = c.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil {
			_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
			return
		}
		if credentials != nil {
			if err := c.TokenRepo.DelUserTokenCache(conn, *credentials, value.Raw); err != nil {
				_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
				return
			}
		}
	} else {
		credentials = jwt2.GetCredentials(ctx)
		if credentials == nil {
			_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
			return
		} else {
			jwt2.RemoveCredentials(ctx)
		}
	}

	ctx.Application().Logger().Infof("%d 退出系统", credentials.UserId)
	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "退出", nil))
}

func (c *AccountCtrl) UserExpire(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	sess, err := c.TokenRepo.GetRedisSession(conn, value.Raw)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}
	if sess != nil {
		if err := c.TokenRepo.UpdateUserTokenCacheExpire(conn, *sess, value.Raw); err != nil {
			_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
			return
		}
	}

	_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultPass, "", nil))
}
