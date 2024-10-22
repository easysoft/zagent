package multipass

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	multiPassService "github.com/easysoft/zagent/internal/host/service/multipass"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"log"
	"testing"
)

func TestMultiPass(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)
	agentConf.Inst.DirImage = "/home/hind3ight/kvm/image/"

	srv := multiPassService.MultiPassService{}
	req := v1.MultiPassReq{
		VmUniqueName: "test1",
		VmMemory:     4096,
		Disk:         10,
		Cpus:         1,
		ImagePath:    "ubuntu.img",
		ImgFrom:      "file",
	}

	vm, _ := srv.CreateVm(&req, false)
	log.Printf("%#v", vm)

	srv.RebootVmByName("test1")

	srv.SuspendVmByName("test1")

	srv.ResumeVmByName("test1")

	srv.DestroyVm("test1")

}
