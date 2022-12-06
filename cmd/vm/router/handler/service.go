package vmHandler

import (
	v1 "github.com/easysoft/zagent/cmd/vm/router/v1"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	vmAgentService "github.com/easysoft/zagent/internal/vm/service"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type ServiceCtrl struct {
	ToolService   *vmAgentService.ToolService   `inject:""`
	StatusService *vmAgentService.StatusService `inject:""`
}

func NewCheckCtrl() *ServiceCtrl {
	return &ServiceCtrl{}
}

// @summary 检测虚拟机服务状态
// @Accept json
// @Produce json
// @Param VmServiceCheckReq body v1.VmServiceCheckReq true "Service Check Request Object"
// @Success 200 {object} v1.VmServiceCheckResp "code = success | fail"
// @Router /api/v1/service/check [post]
func (c *ServiceCtrl) Check(ctx iris.Context) {
	req := v1.VmServiceCheckReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	serviceStatus, _ := c.StatusService.Check(req)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", serviceStatus))
	return
}

// @summary 安装虚拟机服务
// @Accept json
// @Produce json
// @Param VmServiceCheckReq body v1.VmServiceInstallReq true "Service Install Request Object"
// @Success 200 {object} v1.VmServiceInstallResp "code = success | fail"
// @Router /api/v1/service/setup [post]
func (c *ServiceCtrl) Setup(ctx iris.Context) {
	req := v1.VmServiceInstallReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	resp, err := c.ToolService.Setup(req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", resp))
	return
}
