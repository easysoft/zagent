package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/agent/service"
	consts "github.com/easysoft/zv/internal/comm/const"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type SecurityCtrl struct {
	JobService *agentService.JobService `inject:""`
}

func NewSecurityCtrl() *JobCtrl {
	return &JobCtrl{}
}

func (c *SecurityCtrl) RefreshSecret(ctx iris.Context) {
	req := v1.SecurityReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	consts.AuthSecret = req.Secret
	consts.IsSecretChanged = true

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to refresh secret", nil))
	return
}
