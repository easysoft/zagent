package vmHandler

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type TestCtrl struct {
}

func NewTestCtrl() *TestCtrl {
	return &TestCtrl{}
}

func (c *TestCtrl) Test(ctx iris.Context) {
	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success", iris.Map{
		"code": 0,
	}))
	return
}
