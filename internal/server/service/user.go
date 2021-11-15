package serverService

import (
	"fmt"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	bizCasbin "github.com/easysoft/zagent/internal/server/biz/casbin"
	bizConst "github.com/easysoft/zagent/internal/server/biz/const"
	jwt2 "github.com/easysoft/zagent/internal/server/biz/jwt"
	"github.com/easysoft/zagent/internal/server/biz/redis"
	"github.com/easysoft/zagent/internal/server/conf"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type UserService struct {
	UserRepo  *repo.UserRepo  `inject:""`
	TokenRepo *repo.TokenRepo `inject:""`

	CasbinService *bizCasbin.CasbinService `inject:""`
}

func NewUserService() *UserService {
	return &UserService{}
}

// CheckLogin check login user
func (s *UserService) CheckLogin(ctx iris.Context, u *model.User, password string) (*model.Token, bool, string) {
	if u.ID == 0 {
		return nil, false, "用户不存在"
	} else {
		uid := strconv.FormatUint(uint64(u.ID), 10)
		if serverConf.Inst.Redis.Enable && s.TokenRepo.IsUserTokenOver(uid) {
			return nil, false, "已达到同时登录设备上限"
		}
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenStr, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			cred := bizConst.UserCredentials{
				UserId:       uid,
				LoginType:    bizConst.LoginTypeWeb,
				AuthType:     bizConst.AuthPwd,
				CreationDate: time.Now().Unix(),
				Scope:        s.TokenRepo.GetUserScope("admin"),
				Token:        tokenStr,
			}

			if serverConf.Inst.Redis.Enable {
				conn := redisUtils.GetRedisClusterClient()
				defer conn.Close()

				if err := s.TokenRepo.CacheToRedis(conn, cred, tokenStr); err != nil {
					return nil, false, err.Error()
				}
				if err := s.TokenRepo.SyncUserTokenCache(conn, cred, tokenStr); err != nil {
					return nil, false, err.Error()
				}
			} else {
				jwt2.SaveCredentials(ctx, &cred)
			}

			return &model.Token{Token: tokenStr}, true, "登录成功"
		} else {
			return nil, false, "用户名或密码错误"
		}
	}
}

// CreateUser create user
func (s *UserService) CreateUser(u *model.User) error {
	u.Password = _stringUtils.HashPassword(u.Password)
	if err := s.UserRepo.DB.Create(u).Error; err != nil {
		return err
	}

	s.addRoles(u)

	return nil
}

// UpdateUserById update user by id
func (s *UserService) UpdateUserById(id uint, nu *model.User) error {
	if len(nu.Password) > 0 {
		nu.Password = _stringUtils.HashPassword(nu.Password)
	}
	if err := s.UserRepo.UpdateObj(&model.User{}, nu, id); err != nil {
		return err
	}

	s.addRoles(nu)
	return nil
}

// addRoles add roles for user
func (s *UserService) addRoles(user *model.User) {
	if len(user.RoleIds) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := s.CasbinService.Enforcer.DeleteRolesForUser(userId); err != nil {
			color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
		}

		for _, roleId := range user.RoleIds {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := s.CasbinService.Enforcer.AddRoleForUser(userId, roleId); err != nil {
				color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
			}
		}
	}
}

func (s *UserService) UpdateRefreshToken(id uint, token string) {
	s.UserRepo.UpdateRefreshToken(id, token)
}

func (s *UserService) UpdateUserDefaultProject(userId string, projectId int) {
	s.UserRepo.UpdateUserDefaultProject(userId, projectId)
}
