package agentZentaoService

import (
	"encoding/json"
	"github.com/ajg/form"
	"github.com/bitly/go-simplejson"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_var "github.com/easysoft/zagent/internal/pkg/var"
	"github.com/fatih/color"
	"github.com/yosssi/gohtml"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	RequestTypePathInfo = "PATH_INFO"
)

var (
	requestType   = ""
	zenTaoVersion = ""
	sessionVar    = ""
	sessionId     = ""
	requestFix    = ""
)

type ZentaoResponse struct {
	Status string
	Data   string
}

type ZentaoService struct {
}

func NewZentaoService() *ZentaoService {
	return &ZentaoService{}
}

func (s *ZentaoService) GetConfig(baseUrl string) bool {
	url := baseUrl + "?mode=getconfig"
	body, ok := s.Get(url)
	if !ok {
		return false
	}

	json, _ := simplejson.NewJson([]byte(body))
	requestType, _ = json.Get("requestType").String()
	zenTaoVersion, _ = json.Get("version").String()
	sessionId, _ = json.Get("sessionID").String()
	sessionVar, _ = json.Get("sessionVar").String()
	requestFix, _ = json.Get("requestFix").String()

	// check site path by calling login interface
	uri := ""
	if requestType == RequestTypePathInfo {
		uri = "user-login.json"
	} else {
		uri = "index.php?m=user&f=login&t=json"
	}
	url = baseUrl + uri
	body, ok = s.Get(url)
	if !ok {
		return false
	}

	return true
}

func (s *ZentaoService) Get(url string) (string, bool) {
	client := &http.Client{}

	if requestType == RequestTypePathInfo {
		url = url + "?" + sessionVar + "=" + sessionId
	} else {
		url = url + "&" + sessionVar + "=" + sessionId
	}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		if _var.Verbose {
			_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return")+reqErr.Error(), color.FgRed)
		}
		return "", false
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		if _var.Verbose {
			_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return")+respErr.Error(), color.FgRed)
		}
		return "", false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _var.Verbose {
		_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return") + _logUtils.ConvertUnicode(bodyStr))
	}

	var bodyJson ZentaoResponse
	jsonErr := json.Unmarshal(bodyStr, &bodyJson)
	if jsonErr != nil {
		if strings.Index(string(bodyStr), "<html>") > -1 {
			if _var.Verbose {
				_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return") + " HTML - " +
					gohtml.FormatWithLineNo(string(bodyStr)))
			}
			return "", false
		} else {
			if _var.Verbose {
				_logUtils.Infof(jsonErr.Error(), color.FgRed)
			}
			return "", false
		}
	}

	defer resp.Body.Close()

	status := bodyJson.Status
	if status == "" { // 非嵌套结构
		return string(bodyStr), true
	} else { // 嵌套结构
		dataStr := bodyJson.Data
		return dataStr, status == "success"
	}
}

func (s *ZentaoService) Post(url string, params interface{}, useFormFormat bool) (string, bool) {
	if requestType == RequestTypePathInfo {
		url = url + "?" + sessionVar + "=" + sessionId
	} else {
		url = url + "&" + sessionVar + "=" + sessionId
	}
	url = url + "&XDEBUG_SESSION_START=PHPSTORM"

	jsonStr, _ := json.Marshal(params)
	if _var.Verbose {
		_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_address") + url)
	}

	client := &http.Client{}

	val := string(jsonStr)
	if useFormFormat {
		val, _ = form.EncodeToString(params)
		// convert data to post fomat
		re3, _ := regexp.Compile(`([^&]*?)=`)
		val = re3.ReplaceAllStringFunc(string(val), replacePostData)
	}

	if _var.Verbose {
		_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_params") + val)
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(val))
	if reqErr != nil {
		if _var.Verbose {
			_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return")+reqErr.Error(), color.FgRed)
		}
		return "", false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := client.Do(req)
	if respErr != nil {
		if _var.Verbose {
			_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return")+respErr.Error(), color.FgRed)
		}
		return "", false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _var.Verbose {
		_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return") + _logUtils.ConvertUnicode(bodyStr))
	}

	var bodyJson ZentaoResponse
	jsonErr := json.Unmarshal(bodyStr, &bodyJson)
	if jsonErr != nil {
		if strings.Index(string(bodyStr), "<html>") > -1 { // some api return a html
			if _var.Verbose {
				_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return") + " HTML - " +
					gohtml.FormatWithLineNo(string(bodyStr)))
			}
			return string(bodyStr), true
		} else {
			if _var.Verbose {
				_logUtils.Infof(_i118Utils.I118Prt.Sprintf("server_return")+jsonErr.Error(), color.FgRed)
			}
			return "", false
		}
	}

	defer resp.Body.Close()

	status := bodyJson.Status
	if status == "" { // 非嵌套结构
		return string(bodyStr), true
	} else { // 嵌套结构
		dataStr := bodyJson.Data
		return dataStr, status == "success"
	}
}

func (s *ZentaoService) replacePostData(str string) string {
	str = strings.ToLower(str[:1]) + str[1:]

	if strings.Index(str, ".") > -1 {
		str = strings.Replace(str, ".", "[", -1)
		str = strings.Replace(str, "=", "]=", -1)
	}
	return str
}

func replacePostData(str string) string {
	str = strings.ToLower(str[:1]) + str[1:]

	if strings.Index(str, ".") > -1 {
		str = strings.Replace(str, ".", "[", -1)
		str = strings.Replace(str, "=", "]=", -1)
	}
	return str
}
