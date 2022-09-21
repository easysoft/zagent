package handler

import (
	"encoding/json"
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
	"github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"time"
)

type VmCtrl struct {
	BaseCtrl

	AssertService *serverService.AssertService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

// Register
// @summary 向服务器注册虚拟机
// @Accept json
// @Produce json
// @Param task body v1.VmRegisterReq true "RunModeVm Object"
// @Success 200 {object} _httpUtils.Response "code = success? 1 : 0"
// @Router /api/v1/client/vm/register [post]
func (c *VmCtrl) Register(ctx iris.Context) {
	req := v1.VmRegisterReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(_const.ResultFail, err.Error(), nil))
		return
	}

	str, _ := json.Marshal(req)
	_logUtils.Infof("%v", str)

	//success := c.AssertService.RegisterVm(req)
	//if !success {
	//	ctx.StopWithJSON(http.StatusInternalServerError, "register fail")
	//	return
	//}

	data := domain.RegisterResp{
		Token:           "123",
		ExpiredTimeUnix: time.Now().Unix() + 24*3600,
	}
	_, _ = ctx.JSON(data)

	return
}
