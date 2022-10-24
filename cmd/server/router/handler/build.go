package handler

import (
	"encoding/json"
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	serverService "github.com/easysoft/zv/internal/server/service"
	_const "github.com/easysoft/zv/pkg/const"
	_dateUtils "github.com/easysoft/zv/pkg/lib/date"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"github.com/mitchellh/mapstructure"
	"path/filepath"
	"time"
)

type BuildCtrl struct {
	BaseCtrl

	BuildService *serverService.BuildService `inject:""`
}

func NewBuildCtrl() *BuildCtrl {
	return &BuildCtrl{}
}
func (c *BuildCtrl) UploadResult(ctx iris.Context) {
	dir := filepath.Join(_const.UploadDir, _dateUtils.DateStr(time.Now()))
	_fileUtils.MkDirIfNeeded(dir)

	uploadedFiles, n, err := ctx.UploadFormFiles(dir, beforeFileSave)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, fmt.Sprintf("获取上传的文件错误, 字节%d", n), nil))
	}

	filePath := "N/A"
	if uploadedFiles == nil || len(uploadedFiles) == 0 {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, fmt.Sprintf("上传的文件为空, 字节%d", n), nil))
	} else {
		filePath = filepath.Join(dir, uploadedFiles[0].Filename)
	}

	params := ctx.FormValues()
	build := domain.Build{}
	arr, ok := params["result"]
	if ok && len(arr) > 0 {
		testResult := domain.TestResult{}
		json.Unmarshal([]byte(arr[0]), &testResult)

		mapstructure.Decode(testResult.Payload.(map[string]interface{}), &build)
	}

	build.ResultPath = filePath
	jsn, _ := json.Marshal(build)
	build.ResultMsg = string(jsn)
	build.Progress = consts.ProgressCompleted
	build.Status = consts.StatusPass

	c.BuildService.SaveResult(build)

	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass,
		fmt.Sprintf("操作成功 %s %#v", filePath, params), nil))
}
