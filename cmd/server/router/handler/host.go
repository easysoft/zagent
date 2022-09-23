package handler

import (
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
	"github.com/easysoft/zv/internal/pkg/domain"
	serverService "github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

type HostCtrl struct {
	BaseCtrl

	AssertService *serverService.AssertService `inject:""`
}

func NewHostCtrl() *HostCtrl {
	return &HostCtrl{}
}

// Register
// @summary 向服务器注册宿主机
// @Accept json
// @Produce json
// @Param task body v1.HostRegisterReq true "RunModeHost Object"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/host/register [post]
func (c *HostCtrl) Register(ctx iris.Context) {
	req := v1.HostRegisterReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	_logUtils.Infof("%v", ctx.Request().RemoteAddr)

	success := c.AssertService.RegisterHost(req)
	if !success {
		ctx.StopWithJSON(http.StatusInternalServerError, "register fail")
		return
	}

	data := domain.RegisterResp{
		Token:           "123",
		ExpiredTimeUnix: time.Now().Unix() + 24*3600,
	}
	_, _ = ctx.JSON(data)

	return
}
