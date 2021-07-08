package main

import (
	"encoding/json"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"testing"
)

func TestUpload(t *testing.T) {
	_logUtils.Init(agentConst.AppName)
	result := commDomain.TestResult{Name: "Result Name"}

	zipFile := "/Users/aaron/testResult.zip"

	result.Payload = "Payload..."
	uploadResultUrl := _httpUtils.GenUrl("http://localhost:8085/", "client/build/uploadResult")

	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}
