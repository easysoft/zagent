package kvm

import (
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"testing"
)

func TestLibVirt(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	agentConf.Inst.Host = "192.168.0.56"
	agentConf.Inst.User = "aaron"
	agentConf.Init()

	service := hostKvmService.NewLibvirtService()

	/**
	src:  xml template
	base: backing file, get by vm's OsCategory properties etc.
	*/
	vm := domain.Vm{DiskSize: 40000,
		OsCategory: consts.Windows, OsType: consts.Win10,
		OsVersion: "x64-pro", OsLang: consts.ZH_CN}

	dom, macAddress, vncAddress, err := service.CreateVmTest(&vm)
	if err != nil {
		_logUtils.Infof("fail to create vm, err %s", err.Error())
		return
	}
	defer dom.Free()

	name, _ := dom.GetName()
	_logUtils.Infof("%s: %s, %s", name, macAddress, vncAddress)
}