package main

import (
	"encoding/json"
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"testing"
)

func TestUpload(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)
	result := commDomain.TestResult{Name: "Result Name"}
	build := commDomain.Build{ID: 100, Name: "Result Name"}

	zipFile := "/Users/aaron/testResult.zip"

	result.Payload = build
	uploadResultUrl := _httpUtils.GenUrl("http://localhost:8085/", "client/build/uploadResult")

	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}
