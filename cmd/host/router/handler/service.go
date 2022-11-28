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

// @summary 检测宿主机服务状态
// @Accept json
// @Produce json
// @Param HostServiceCheckReq body v1.HostServiceCheckReq true "Service Check Request Object"
// @Success 200 {object} v1.HostServiceCheckResp "code = success | fail"
// @Router /api/v1/service/check [post]
func (c *ServiceCtrl) CheckService(ctx iris.Context) {
	req := v1.HostServiceCheckReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	serviceStatus, _ := c.StatusService.Check(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", serviceStatus))
	return
}
