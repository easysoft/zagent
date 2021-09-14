package wsapiclient

import (
	vmwareService "github.com/easysoft/zagent/internal/server/service/vendors/vmware"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	srv := vmwareService.NewVMWareService()
	err := srv.Connect("https://192.168.0.56:8697", "aaron", "P@ssw0rd")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	vms, err := srv.GetVms()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	log.Printf("%#v", vms)
}
