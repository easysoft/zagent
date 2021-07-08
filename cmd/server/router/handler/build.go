package handler

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	"github.com/kataras/iris/v12"
	"time"
)

type BuildCtrl struct {
	BaseCtrl
	Ctx iris.Context
}

func NewBuildCtrl() *BuildCtrl {
	return &BuildCtrl{}
}
func (c *BuildCtrl) UploadResult(ctx iris.Context) {
	dir := _const.UploadDir + _dateUtils.DateStr(time.Now())
	_fileUtils.MkDirIfNeeded(dir)

	c.Ctx.UploadFormFiles("./uploads", beforeFileSave)
}
