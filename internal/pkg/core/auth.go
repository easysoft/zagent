package core

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	authUtils "github.com/easysoft/zagent/internal/pkg/utils/auth"
	_const "github.com/easysoft/zagent/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Auth() iris.Handler {
	return func(ctx *context.Context) {
		token := authUtils.GetTokenInAuthorization(ctx.GetHeader(_const.Authorization))

		success := false
		if token == consts.AuthToken {
			success = true
		}

		if success {
			ctx.Next()
		} else {
			ctx.StopWithJSON(http.StatusUnauthorized, _httpUtils.RespData(consts.ResultFail, "wrong token", nil))
		}
	}
}
