package agentService

import (
	"fmt"
	agentHttpUtils "github.com/easysoft/zagent/internal/agent/utils/request"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
)

type InterfaceRequestService struct {
	CommonService
}

func NewInterfaceRequestService() *InterfaceRequestService {
	return &InterfaceRequestService{}
}

func (s *InterfaceRequestService) Request(build *commDomain.Build, interf *commDomain.TestInterface) (respStr string) {
	method := interf.Request.Method
	if method == _const.Get {
		respStr = s.Get(*interf)
	}
	interf.Raws = append(interf.Raws, respStr)

	_logUtils.Infof("exec %s request, get response %s", method, respStr)

	return
}

func (s *InterfaceRequestService) Get(interf commDomain.TestInterface) (respStr string) {
	reqObj := interf.Request
	urlObj := reqObj.URL

	url := ""
	if urlObj.Port != 0 {
		url = fmt.Sprintf("%s://%s:%d", urlObj.Protocol, urlObj.Host, urlObj.Port)
	} else {
		url = fmt.Sprintf("%s://%s", urlObj.Protocol, urlObj.Host)
	}
	url = url + urlObj.Path

	respStr, _ = agentHttpUtils.Get(reqObj.Header, url, urlObj.Params)

	return
}
