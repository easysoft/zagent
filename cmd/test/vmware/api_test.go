package wsapiclient

import (
	vmwareService "github.com/easysoft/zagent/internal/server/service/vendors/vmware"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	srv := vmwareService.NewVMWareService()
	srv.Connect("https://192.168.0.56:8697", "aaron", "P@ssw0rd")

	vms, _ := srv.GetVms()
	log.Printf("%#v", vms)

	id := ""

	srv.DestroyVm(id)
}
