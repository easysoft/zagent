package commonService

import (
	"bytes"
	"fmt"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	gateway "github.com/rpcx-ecosystem/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	"io/ioutil"
	"net/http"
)

type RpcService struct {
}

func NewRpcService() *RpcService {
	return &RpcService{}
}

func (s RpcService) SeleniumTest(build model.Build) (result _domain.RpcResp) {
	seleniumTestTo := model.NewBuildTo(build)
	seleniumTestTo.BrowserType = build.Queue.BrowserType
	seleniumTestTo.BrowserVersion = build.Queue.BrowserVersion

	obj := interface{}(seleniumTestTo)
	result = s.Request(build.NodeIp, build.NodePort, "perform", "Perform", &obj)

	return
}

func (s RpcService) AppiumTest(build model.Build) (result _domain.RpcResp) {
	appiumTestTo := model.NewBuildTo(build)
	appiumTestTo.AppiumPort = build.AppiumPort

	obj := interface{}(appiumTestTo)
	result = s.Request(build.NodeIp, build.NodePort, "perform", "Perform", &obj)

	return
}

func (s RpcService) CreateVm(hostIp string, hostPort int, req commDomain.KvmReq) (result _domain.RpcResp) {
	obj := interface{}(req)
	result = s.Request(hostIp, hostPort, "vm", "Create", &obj)
	return
}

func (s RpcService) DestroyVm(hostIp string, hostPort int, req commDomain.KvmReq) (result _domain.RpcResp) {
	obj := interface{}(req)
	s.Request(hostIp, hostPort, "vm", "Destroy", &obj)

	result.Success(fmt.Sprintf("success to destroy vm via rpc %#v.", req))
	return
}

func (s RpcService) Request(ip string, port int, apiPath string, method string, param *interface{}) (
	rpcResult _domain.RpcResp) {

	cc := &codec.MsgpackCodec{}

	data, _ := cc.Encode(param)
	url := fmt.Sprintf("http://%s:%d/", ip, port)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		msg := fmt.Sprintf("Fail to create request: %s", err.Error())
		_logUtils.Errorf(msg)
		rpcResult.Fail(msg)
		return
	}

	// 设置header
	h := req.Header
	h.Set(gateway.XServicePath, apiPath)
	h.Set(gateway.XServiceMethod, method)
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		msg := fmt.Sprintf("fail to call, err: %s", err.Error())
		_logUtils.Errorf(msg)
		rpcResult.Fail(msg)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("fail to read response, err: %s", err.Error())
		_logUtils.Errorf(msg)
		rpcResult.Fail(msg)

	}

	err = cc.Decode(body, &rpcResult)
	if err != nil {
		msg := fmt.Sprintf("fail to decode reply, err: %s", err.Error())
		_logUtils.Errorf(msg)
		rpcResult.Fail(msg)
	}

	msg := fmt.Sprintf("agent return %d-%s.", rpcResult.Code, rpcResult.Msg)
	_logUtils.Info(msg)

	return
}
