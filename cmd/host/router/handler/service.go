package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	hostStatusService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type ServiceCtrl struct {
	StatusService *hostStatusService.StatusService `inject:""`
}

func NewCheckCtrl() *ServiceCtrl {
	return &ServiceCtrl{}
}

func (c *ServiceCtrl) CheckStatus(ctx iris.Context) {
	req := v1.ServiceReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	serviceStatus, _ := c.StatusService.Check(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", serviceStatus))
	return
}
