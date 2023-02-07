package kvm

import (
	"log"
	"testing"

	hostKvmService "github.com/easysoft/zagent/internal/host/service/kvm"
)

func TestLibVirt(t *testing.T) {
	service := hostKvmService.NewLibvirtService()

	err := service.StartVmByName("test-win10")
	log.Print(err.Error())

	/**

	src:  xml template
	base: backing file, get by vm's OsCategory properties etc.
	*/
	//vm := domain.RunModeVm{DiskSize: 40000,
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
