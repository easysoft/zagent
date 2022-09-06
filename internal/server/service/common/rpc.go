package commonService

import (
	"encoding/json"
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/server/model"
	_domain "github.com/easysoft/zv/pkg/domain"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
)

type RemoteService struct {
}

func NewRpcService() *RemoteService {
	return &RemoteService{}
}

func (s RemoteService) SeleniumTest(build model.Build) (result _domain.RemoteResp) {
	buildTo := model.NewBuildTo(build)
	buildTo.BrowserType = build.BrowserType
	buildTo.BrowserVersion = build.BrowserVersion

	obj := interface{}(buildTo)
	result = s.HttpRequest(build.NodeIp, build.NodePort, "job", &obj)

	return
}

func (s RemoteService) AppiumTest(build model.Build) (result _domain.RemoteResp) {
	appiumTestTo := model.NewBuildTo(build)
	appiumTestTo.AppiumPort = build.AppiumPort

	obj := interface{}(appiumTestTo)
	result = s.HttpRequest(build.NodeIp, build.NodePort, "job", &obj)

	return
}

func (s RemoteService) UnitTest(build model.Build) (result _domain.RemoteResp) {
	buildTo := model.NewBuildTo(build)

	obj := interface{}(buildTo)
	result = s.HttpRequest(build.NodeIp, build.NodePort, "job", &obj)

	return
}

func (s RemoteService) CreateKvm(hostIp string, hostPort int, req v1.KvmReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "kvm/create", &obj)
	return
}
func (s RemoteService) DestroyKvm(hostIp string, hostPort int, req v1.KvmReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "kvm/destroy", &obj)
	return
}

func (s RemoteService) CreateVirtualBox(hostIp string, hostPort int, req v1.VirtualBoxReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "virtualbox/create", &obj)
	return
}
func (s RemoteService) DestroyVirtualBox(hostIp string, hostPort int, req v1.VirtualBoxReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "virtualbox/destroy", &obj)
	return
}

func (s RemoteService) CreateVmWare(hostIp string, hostPort int, req v1.VmWareReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "vmware/create", &obj)
	return
}
func (s RemoteService) DestroyVmWare(hostIp string, hostPort int, req v1.VmWareReq) (result _domain.RemoteResp) {
	obj := interface{}(req)
	result = s.HttpRequest(hostIp, hostPort, "vmware/destroy", &obj)
	return
}

func (s RemoteService) HttpRequest(ip string, port int, apiPath string, data *interface{}) (
	remoteResult _domain.RemoteResp) {

	url := fmt.Sprintf("http://%s:%d/api/v1/%s", ip, port, apiPath)

	bodyBytes, err := _httpUtils.Post(url, data)
	if err != nil {
		return
	}

	json.Unmarshal(bodyBytes, &remoteResult)

	return
}

//func (s RemoteService) RpcRequest(ip string, port int, apiPath string, method string, param *interface{}) (
//	remoteResult _domain.RemoteResp) {
//
//	cc := &codec.MsgpackCodec{}
//
//	data, _ := cc.Encode(param)
//	url := fmt.Sprintf("http://%s:%d/", ip, port)
//	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
//	if err != nil {
//		msg := fmt.Sprintf("Fail to create request: %s", err.Error())
//		_logUtils.Errorf(msg)
//		remoteResult.Fail(msg)
//		return
//	}
//
//	// 设置header
//	h := req.Header
//	h.Set(gateway.XServicePath, apiPath)
//	h.Set(gateway.XServiceMethod, method)
//	h.Set(gateway.XMessageID, "10000")
//	h.Set(gateway.XMessageType, "0")
//	h.Set(gateway.XSerializeType, "3")
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		msg := fmt.Sprintf("fail to call, err: %s", err.Error())
//		_logUtils.Errorf(msg)
//		remoteResult.Fail(msg)
//		return
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		msg := fmt.Sprintf("fail to read response, err: %s", err.Error())
//		_logUtils.Errorf(msg)
//		remoteResult.Fail(msg)
//
//	}
//
//	err = cc.Decode(body, &remoteResult)
//	if err != nil {
//		msg := fmt.Sprintf("fail to decode reply, err: %s", err.Error())
//		_logUtils.Errorf(msg)
//		remoteResult.Fail(msg)
//	}
//
//	msg := fmt.Sprintf("agent return %d-%s.", remoteResult.Code, remoteResult.Msg)
//	_logUtils.Info(msg)
//
//	return
//}
