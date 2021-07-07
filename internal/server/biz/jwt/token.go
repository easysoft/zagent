package jwt

import (
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	bizConst "github.com/easysoft/zagent/internal/server/biz/const"
	"github.com/easysoft/zagent/internal/server/biz/redis"
	"github.com/easysoft/zagent/internal/server/conf"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
	"time"
)

type TokenService struct {
	Middleware *JwtService      `inject:""`
	Enforcer   *casbin.Enforcer `inject:""`
	UserRepo   *repo.UserRepo   `inject:""`
	TokenRepo  *repo.TokenRepo  `inject:""`
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

// Get returns the user (&token) information for this client/request
func (m *TokenService) Get(ctx iris.Context) *jwt.Token {
	v := ctx.Values().Get(m.Middleware.Config.ContextKey)
	if v == nil {
		return nil
	}
	return v.(*jwt.Token)
}

func (m *TokenService) Serve(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	value := m.Get(ctx)

	credentials, _ := m.GetCredentials(value, ctx)
	if credentials == nil { // jwt token expired, try to refresh
		tokenStr := value.Raw

		// find user by token
		user, _ := m.UserRepo.GetByToken(tokenStr)

		if user.ID != 0 && // user exist and token not expire
			time.Now().Unix()-user.TokenUpdatedTime.Unix() < _const.UserTokenExpireTime {

			// refresh the credentials
			uid := strconv.FormatUint(uint64(user.ID), 10)
			cred := bizConst.UserCredentials{
				UserId:       uid,
				LoginType:    bizConst.LoginTypeWeb,
				AuthType:     bizConst.AuthPwd,
				CreationDate: time.Now().Unix(),
				Scope:        bizConst.AdminScope,
				Token:        tokenStr,
			}

			if serverConf.Inst.Redis.Enable {
				conn := redisUtils.GetRedisClusterClient()
				defer conn.Close()

				if err := m.TokenRepo.CacheToRedis(conn, cred, tokenStr); err != nil {
					m.Middleware.Config.ErrorHandler(ctx, err)
					return
				}
				if err := m.TokenRepo.SyncUserTokenCache(conn, cred, tokenStr); err != nil {
					m.Middleware.Config.ErrorHandler(ctx, err)
					return
				}
			} else {
				SaveCredentials(ctx, &cred)
			}
		}
	}

	// load again
	credentials, _ = m.GetCredentials(value, ctx)
	if credentials == nil {
		ctx.StopExecution()
		_, _ = ctx.JSON(_httpUtils.ApiRes(401, "", nil))
		ctx.StopExecution()
		return
	}

	ctx.Next()
}

func (m *TokenService) GetCredentials(value *jwt.Token, ctx iris.Context) (
	credentials *bizConst.UserCredentials, err error) {
	if serverConf.Inst.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = m.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil || credentials == nil {
			m.TokenRepo.UserTokenExpired(value.Raw)
			_, _ = ctx.JSON(_httpUtils.ApiRes(401, "", nil))
			ctx.StopExecution()
			return
		}
	} else {
		credentials = GetCredentials(ctx)
	}

	return
}
