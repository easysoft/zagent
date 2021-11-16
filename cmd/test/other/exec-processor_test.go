package other

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/lib/rpc"
	"log"
	"testing"
)

func TestExecProcessor(t *testing.T) {
	build := domain.Build{BuildType: consts.InterfaceTest}

	processor := domain.TestProcessor{Type: consts.Simple}

	dataLoop := domain.TestProcessor{Type: consts.DataLoop}
	dataLoop.Src = consts.ZenData
	dataLoop.Path = "http://127.0.0.1:8848/data?config=demo/default.yaml&F=field_format"

	req := domain.Request{
		Method: _const.Get,
		URL:    domain.URL{Protocol: "http", Host: "max.demo.zentao.net/"},
	}
	req.URL.Params = append(req.URL.Params, domain.Entity{Key: "mode", Value: "getconfig"})

	interf := domain.TestInterface{Name: "禅道配置", Request: req}

	dataLoop.Children = append(dataLoop.Children, interf)
	processor.Children = append(processor.Children, dataLoop)
	build.TestScenario.Processor = processor

	url := fmt.Sprintf("http://127.0.0.1:%d/", consts.AgentPort)
	result := _rpcUtils.Post(url, string(_const.Post), "task", "RunRemote", build)

	log.Printf("%v", result)
}
