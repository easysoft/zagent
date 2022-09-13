package kvm

import (
	agentConf "github.com/easysoft/zv/internal/agent/conf"
	"github.com/easysoft/zv/internal/comm/const"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"testing"
)

func TestLibVirt(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	agentConf.Inst.Host = "192.168.0.56"
	agentConf.Inst.User = "aaron"
	agentConf.Init(consts.AppNameAgentHost)

	//service := hostKvmService.NewLibvirtService()

	/**
	src:  xml template
	base: backing file, get by vm's OsCategory properties etc.
	*/
	//vm := domain.Vm{DiskSize: 40000,
	//	OsCategory: consts.Windows, OsType: consts.Win10,
	//	OsVersion: "x64-pro", OsLang: consts.ZH_CN,
	//	VmCpu: 2,
	//	VmMemorySize: 40000,
	//	VmDiskSize: 50000
	//}
	//
	//dom, macAddress, vncAddress, err := service.CreateVmTest(&vm)
	//if err != nil {
	//	_logUtils.Infof("fail to create vm, err %s", err.Error())
	//	return
	//}
	//defer dom.Free()
	//
	//name, _ := dom.GetName()
	//_logUtils.Infof("%s: %s, %s", name, macAddress, vncAddress)
}
