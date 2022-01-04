package hostHandler

import (
	v1 "github.com/easysoft/zv/cmd/agent-host/router/v1"
	multiPassService "github.com/easysoft/zv/internal/agent-host/service/multipass"
	_httpUtils "github.com/easysoft/zv/internal/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type MultiPassCtrl struct {
	MultiPassService *multiPassService.MultiPassService `inject:""`
}

func NewMultiPassCtrl() *MultiPassCtrl {
	return &MultiPassCtrl{}
}

// List
// @summary 获取MultiPass虚拟机
// @Accept json
// @Produce json
// @Param task body v1.MultiPassReq true "MultiPass Request Object"
// @Success 200 {object} _httpUtils.Response{data=v1.MultiPassResp} "code = success? 1 : 0"
// @Router /api/v1/MultiPass/list [post]
func (c *MultiPassCtrl) List(ctx iris.Context) {
	req := v1.MultiPassReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	domains, err := c.MultiPassService.ListVm()
	if err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "fail to get MultiPass", err))
		return
	}

	mp := v1.MultiPassResp{

		Name: domains[0].Name,
		//VncPort:  strconv.Itoa(VmWareVncPort),
	}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to get MultiPass", mp))

	return
}
