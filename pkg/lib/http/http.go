package _httpUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	authUtils "github.com/easysoft/zv/internal/pkg/utils/auth"
	_const "github.com/easysoft/zv/pkg/const"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const ()

func Get(url string) (ret []byte, err error) {
	if consts.Verbose {
		_logUtils.Infof("===DEBUG===  request: %s", url)
	}

	timeout := 3 * time.Second
	if !_const.IsRelease {
		timeout = 60 * time.Second
	}
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		_logUtils.Infof(color.RedString("get request failed, error: %s.", err.Error()))
		return
	}

	authUtils.AddBearTokenIfNeeded(req)

	resp, err := client.Do(req)
	if err != nil {
		_logUtils.Infof(color.RedString("get request failed, error: %s.", err.Error()))
		return
	}
	defer resp.Body.Close()

	if !IsSuccessCode(resp.StatusCode) {
		_logUtils.Infof(color.RedString("read response failed, StatusCode: %d.", resp.StatusCode))
		err = errors.New(resp.Status)
		return
	}

	ret, err = ioutil.ReadAll(resp.Body)
	_logUtils.Infof("===DEBUG=== response: %s", _logUtils.ConvertUnicode(ret))

	if err != nil {
		_logUtils.Infof(color.RedString("read response failed, error ", err.Error()))
		return
	}

	return
}

func Post(url string, data interface{}) (ret []byte, err error) {
	return PostOrPut(url, "POST", data)
}
func Put(url string, data interface{}) (ret []byte, err error) {
	return PostOrPut(url, "PUT", data)
}

func PostOrPut(url string, method string, data interface{}) (ret []byte, err error) {
	if consts.Verbose {
		_logUtils.Infof("===DEBUG===  request: %s", url)
	}

	timeout := 3 * time.Second
	if !_const.IsRelease {
		timeout = 60 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	dataBytes, err := json.Marshal(data)
	if consts.Verbose {
		_logUtils.Infof("===DEBUG===     data: %s", string(dataBytes))
	}

	if err != nil {
		_logUtils.Infof(color.RedString("marshal request failed, error: %s.", err.Error()))
		return
	}

	dataStr := string(dataBytes)

	req, err := http.NewRequest(method, url, strings.NewReader(dataStr))
	if err != nil {
		_logUtils.Infof(color.RedString("post request failed, error: %s.", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	authUtils.AddBearTokenIfNeeded(req)

	resp, err := client.Do(req)
	if err != nil {
		_logUtils.Infof(color.RedString("post request failed, error: %s.", err.Error()))
		return
	}

	if !IsSuccessCode(resp.StatusCode) {
		_logUtils.Infof(color.RedString("post request return '%s'.", resp.Status))
		err = errors.New(resp.Status)
		return
	}

	ret, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if consts.Verbose {
		_logUtils.Infof("===DEBUG=== response: %s", _logUtils.ConvertUnicode(ret))
	}

	if err != nil {
		_logUtils.Infof(color.RedString("read response failed, error: %s.", err.Error()))
		return
	}

	return
}

func IsSuccessCode(code int) (success bool) {
	return code >= 200 && code <= 299
}

func GenUrl(server string, path string) string {
	server = AddUrlPostFixIfNeeded(server)
	url := fmt.Sprintf("%sapi/v1/%s", server, path)
	url = AddUrlPostFixIfNeeded(url)
	return url
}

func GenUrlWithParams(pth string, params map[string]interface{}, baseUrl string) (url string) {
	uri := pth

	index := 0
	for key, val := range params {
		if index == 0 {
			uri += "?"
		} else {
			uri += "&"
		}

		uri += fmt.Sprintf("%v=%v", key, val)
		index++
	}

	url = baseUrl + uri
	url = AddUrlPostFixIfNeeded(url)

	return
}

func AddUrlPostFixIfNeeded(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}
