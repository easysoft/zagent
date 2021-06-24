package agentHttpUtils

import (
	"encoding/json"
	"fmt"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/pkg/var"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(header commDomain.Header, url string, params []commDomain.Entity) (str string, success bool) {
	client := &http.Client{}

	urlWithParams := url + CombineUrlParams(params)

	if _var.Verbose {
		_logUtils.Info(urlWithParams)
	}

	req, reqErr := http.NewRequest("GET", urlWithParams, nil)
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return "", false
	}

	resp, respErr := client.Do(req)

	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return "", false
	}

	bytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		_logUtils.Error(readErr.Error())
		return "", false
	}
	defer resp.Body.Close()

	return string(bytes), true
}

func Post(url string, params interface{}) (interface{}, bool) {
	if _var.Verbose {
		_logUtils.Info(url)
	}
	client := &http.Client{}

	paramStr, err := json.Marshal(params)
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, false
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(paramStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _var.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}

	var result _domain.RpcResp
	json.Unmarshal(bodyStr, &result)

	defer resp.Body.Close()

	code := result.Code
	return result, code == _const.ResultSuccess.Int()
}

func GenUrl(server string, path string) string {
	server = UpdateUrl(server)
	url := fmt.Sprintf("%sapi/v1/%s", server, path)
	return url
}

func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}

func CombineUrlParams(params []commDomain.Entity) (ret string) {
	for _, param := range params {
		if ret != "" {
			ret += "&"
		}
		ret += param.Key + "=" + param.Value
	}

	ret = "?" + ret
	return
}
