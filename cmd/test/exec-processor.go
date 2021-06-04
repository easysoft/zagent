package main

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_rpcUtils "github.com/easysoft/zagent/internal/pkg/libs/rpc"
	"log"
)

func main() {
	processor := commDomain.TestProcessor{Type: _const.Simple}

	dataLoop := commDomain.TestProcessor{Type: _const.DataLoop}
	dataLoop.Src = commConst.ZenData
	dataLoop.Path = "http://127.0.0.1:8848/data?config=demo/default.yaml&F=field_format"

	interf := commDomain.TestInterface{Name: "百度搜索"}
	req := commDomain.Request{
		Method: _const.Get,
		URL:    commDomain.URL{Protocol: "http", Host: "max.demo.zentao.net/"},
	}
	req.URL.Params = append(req.URL.Params, commDomain.Entity{Key: "mode", Value: "getconfig"})

	processor.Children = append(processor.Children, dataLoop)
	dataLoop.Children = append(dataLoop.Children, interf)

	url := fmt.Sprintf("http://127.0.0.1:%d/", _const.RpcPort)
	result := _rpcUtils.Post(url, string(_const.Post), "task", "Exec", processor)

	log.Printf("%v", result)
}
