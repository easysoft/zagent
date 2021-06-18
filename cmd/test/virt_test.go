package main

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/libs/string"
	"strings"
	"testing"
)

func TestVirt(t *testing.T) {
	_logUtils.Init(agentConst.AppName)
	agentConf.Init()

	virtService := agentService.NewLibvirtService()

	/**
	src:  xml template
	base: backing file, get by vm's OsCategory properties etc.
	*/
	src := "base-win10-x64-pro-zh_cn"
	vmName := strings.Replace(src, "base-", "test-", -1) + "-" + _stringUtils.NewUuid()
	vm := commDomain.Vm{Name: vmName, DiskSize: 40000, Src: src,
		OsCategory: commConst.Windows, OsType: commConst.Win10,
		OsVersion: "x64-pro", OsLang: commConst.ZH_CN}

	dom, macAddress, _ := virtService.CreateVm(&vm)
	defer dom.Free()

	name, _ := dom.GetName()
	_logUtils.Infof("%s: %s", name, macAddress)
}
