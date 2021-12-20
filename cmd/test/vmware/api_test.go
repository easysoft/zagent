package wsapiclient

import (
	vmwareService "github.com/easysoft/zv/internal/server/service/vendors/vmware"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	srv := vmwareService.NewVMWareService()
	srv.Connect("https://127.0.0.1:8697", "aaron", "P2ssw0rd@")

	vm, _ := srv.CreateVm("win10-pro-x64-zh_cn", "test-win10-pro-x64-zh_cn", 2, 1024*4)
	log.Printf("%#v", vm)

	srv.DestroyVm(vm.IdVM)
}
