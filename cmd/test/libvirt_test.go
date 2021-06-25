package main

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"strings"
	"testing"
)

func TestLibVirt(t *testing.T) {
	_logUtils.Init(agentConst.AppName)

	agentConf.Inst.Host = "192.168.0.56"
	agentConf.Inst.User = "aaron"
	agentConf.Init()

	service := agentService.NewLibvirtService()

	/**
	src:  xml template
	base: backing file, get by vm's OsCategory properties etc.
	*/
	src := "src-win10-x64-pro-zh_cn"
	vmName := strings.Replace(src, "src-", "test-", -1) + "-" + _stringUtils.NewUuid()
	vm := commDomain.Vm{Name: vmName, DiskSize: 40000, Src: src,
		OsCategory: commConst.Windows, OsType: commConst.Win10,
		OsVersion: "x64-pro", OsLang: commConst.ZH_CN} // base image by os props

	dom, macAddress, vncPort, err := service.CreateVm(&vm)
	if err != nil {
		_logUtils.Infof("fail to create vm, err %s", err.Error())
		return
	}
	defer dom.Free()

	name, _ := dom.GetName()
	_logUtils.Infof("%s: %s, %d", name, macAddress, vncPort)
}
