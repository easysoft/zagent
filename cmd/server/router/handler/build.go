package handler

import (
	"encoding/json"
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	serverService "github.com/easysoft/zagent/internal/server/service"
	"github.com/easysoft/zagent/internal/server/service/kvm"
	"github.com/kataras/iris/v12"
	"github.com/mitchellh/mapstructure"
	"path/filepath"
	"time"
)

type BuildCtrl struct {
	BaseCtrl

	BuildService *serverService.BuildService `inject:""`
	VmService    kvmService.VmService        `inject:""`
}

func NewBuildCtrl() *BuildCtrl {
	return &BuildCtrl{}
}
func (c *BuildCtrl) UploadResult(ctx iris.Context) {
	dir := filepath.Join(_const.UploadDir, _dateUtils.DateStr(time.Now()))
	_fileUtils.MkDirIfNeeded(dir)

	uploaded, n, err := ctx.UploadFormFiles(dir, beforeFileSave)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, fmt.Sprintf("操作失败, 字节%d", n), nil))
	}

	filePath := filepath.Join(dir, uploaded[0].Filename)

	params := ctx.FormValues()
	var testResult domain.TestResult
	json.Unmarshal([]byte(params["result"][0]), &testResult)
	var build domain.Build
	mapstructure.Decode(testResult.Payload.(map[string]interface{}), &build)

	build.ResultPath = filePath
	jsn, _ := json.Marshal(build)
	build.ResultMsg = string(jsn)
	build.Progress = consts.ProgressCompleted
	build.Status = consts.StatusPass

	c.BuildService.SaveResult(build)
	//c.VmService.DestroyRemote(build.VmId) TODO: testing

	_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusOK,
		fmt.Sprintf("操作成功 %s %#v", filePath, params), nil))
}
