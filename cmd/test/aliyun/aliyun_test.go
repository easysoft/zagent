package aliyun

import (
	"fmt"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
	"time"
)

func TestAliyun(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewAliyunEcsService()
	url := fmt.Sprintf("ecs-%s.aliyuncs.com", testconst.ALIYUN_REGION)
	client, err := srv.CreateClient(url, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)

	id, name, _ := srv.CreateInst("vm-001", "ubuntu-20-desktop-x64-zh_cn", client)
	err = srv.StartInst(id, client)

	status := ""
	macAddress := ""
	for i := 0; i < 2*60; i++ {
		<-time.After(1 * time.Second)

		status, macAddress, err = srv.QueryInst(id, testconst.ALIYUN_REGION, client)

		if status == "Running" {
			break
		}
	}

	_logUtils.Infof("%s %s", status, macAddress)

	ip, err := srv.AllocateIp(id, client)
	vncPassword, _ := srv.QueryVncPassword(id, testconst.ALIYUN_REGION, client)
	vncUrl, _ := srv.QueryVncUrl(id, vncPassword, testconst.ALIYUN_REGION, false, client)

	_logUtils.Infof("%s, %s, %s, %s, %s", id, name, ip, vncUrl, vncPassword)

	<-time.After(60 * time.Second)

	err = srv.RemoveInst(id, client)
	_logUtils.Infof("%#v", err)
}
