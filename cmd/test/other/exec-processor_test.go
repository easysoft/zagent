package other

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_rpcUtils "github.com/easysoft/zagent/internal/pkg/lib/rpc"
	"log"
	"testing"
)

func TestExecProcessor(t *testing.T) {
	build := commDomain.Build{BuildType: commConst.InterfaceTest}

	processor := commDomain.TestProcessor{Type: commConst.Simple}

	dataLoop := commDomain.TestProcessor{Type: commConst.DataLoop}
	dataLoop.Src = commConst.ZenData
	dataLoop.Path = "http://127.0.0.1:8848/data?config=demo/default.yaml&F=field_format"

	req := commDomain.Request{
		Method: _const.Get,
		URL:    commDomain.URL{Protocol: "http", Host: "max.demo.zentao.net/"},
	}
	req.URL.Params = append(req.URL.Params, commDomain.Entity{Key: "mode", Value: "getconfig"})

	interf := commDomain.TestInterface{Name: "禅道配置", Request: req}

	dataLoop.Children = append(dataLoop.Children, interf)
	processor.Children = append(processor.Children, dataLoop)
	build.TestScenario.Processor = processor

	url := fmt.Sprintf("http://127.0.0.1:%d/", _const.RpcPort)
	result := _rpcUtils.Post(url, string(_const.Post), "task", "RunRemote", build)

	log.Printf("%v", result)
}
