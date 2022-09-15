package hostHandler

import (
	agentConf "github.com/easysoft/zv/internal/agent/conf"
	"github.com/easysoft/zv/internal/agent/service"
	"github.com/easysoft/zv/internal/comm/domain"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	vmMacMap = sync.Map{}
)

type SecurityCtrl struct {
	JobService *agentService.JobService `inject:""`
}

func NewSecurityCtrl() *JobCtrl {
	return &JobCtrl{}
}

func (c *SecurityCtrl) VmGetSecret(ctx iris.Context) {
	req := domain.SecurityReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	data := domain.SecurityResp{}

	mac, ok := vmMacMap.Load(req.MacAddress)
	if ok && mac != "" {
		data.Secret = agentConf.Inst.Secret
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to refresh secret", data))
	return
}
