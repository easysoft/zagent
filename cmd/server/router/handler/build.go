package handler

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"path/filepath"
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
	dir := filepath.Join(_const.UploadDir, _dateUtils.DateStr(time.Now()))
	_fileUtils.MkDirIfNeeded(dir)

	uploaded, n, err := ctx.UploadFormFiles(dir, beforeFileSave)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, fmt.Sprintf("操作失败, 字节%d", n), nil))
	}

	filePath := filepath.Join(dir, uploaded[0].Filename)

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "操作成功"+filePath, nil))
}
