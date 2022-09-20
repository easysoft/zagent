package vmHandler

import (
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type TestCtrl struct {
}

func NewTestCtrl() *TestCtrl {
	return &TestCtrl{}
}

func (c *TestCtrl) Test(ctx iris.Context) {
	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success", iris.Map{
		"code": 0,
	}))
	return
}
