package main

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"testing"
)

func TestVirt(t *testing.T) {
	_logUtils.Init(agentConst.AppName)

	vm := commDomain.Vm{Name: "test-win10-x64-pro-zh_cn-01",
		Src: "templ-win10-x64-pro-zh_cn", Base: "windows/win10/x64-pro-zh_cn",
		OsCategory: commConst.Windows, OsType: commConst.Win10,
		OsVersion: "x64-pro", SysLang: commConst.ZH_CN}

	virtService := agentService.NewLibvirtService()
	dom, macAddress, _ := virtService.CreateVm(&vm)
	defer dom.Free()

	virtService.StartVm(dom)

	name, _ := dom.GetName()
	_logUtils.Infof("%s: %s", name, macAddress)
}
