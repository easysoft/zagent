package net

import (
	"encoding/json"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/lib/file"
	"github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	"testing"
)

func TestUpload(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)
	result := domain.TestResult{Name: "Result Name"}
	build := domain.Build{ID: 100, Name: "Result Name"}

	zipFile := "/Users/aaron/testResult.zip"

	result.Payload = build
	uploadResultUrl := _httpUtils.GenUrl("http://localhost:8085/", "client/build/uploadResult")

	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}
