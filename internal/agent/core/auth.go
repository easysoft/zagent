package core

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	commonUtils "github.com/easysoft/zv/internal/pkg/utils/common"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Auth() iris.Handler {
	return func(ctx *context.Context) {
		token := commonUtils.GetTokenInAuthorization(ctx.GetHeader(_const.Authorization))

		success := false
		if token == consts.AuthToken {
			success = true
		}

		if success {
			ctx.Next()
		} else {
			ctx.StopWithJSON(http.StatusUnauthorized, _httpUtils.ApiRes(_const.ResultFail, "wrong token", nil))
		}
	}
}
