package aliyun

import (
	"fmt"
	testconst "github.com/easysoft/zv/cmd/test/_const"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/pkg/vendors/aliyun"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"testing"
	"time"
)

func TestAliyunEcs(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	commSrv := aliyun.NewAliyunCommService()
	ecsSrv := aliyun.NewAliyunEcsService()

	ecsUrl := fmt.Sprintf("ecs-%s.aliyuncs.com", testconst.ALIYUN_REGION)
	client, err := commSrv.CreateEcsClient(ecsUrl, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	vpcClient, err := commSrv.CreateVpcClient(ecsUrl, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)

	switchId, _, err := commSrv.GetSwitch(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, vpcClient)
	securityGroupId, err := commSrv.QuerySecurityGroupByVpc(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, client)

	id, name, _ := ecsSrv.CreateInst("vm-001", "tmpl-ubuntu-20-desktop-x64-zh_cn",
		switchId, securityGroupId, testconst.ALIYUN_REGION, client)
	err = ecsSrv.StartInst(id, client)

	status := ""
	macAddress := ""
	for i := 0; i < 2*60; i++ {
		<-time.After(1 * time.Second)

		status, macAddress, err = ecsSrv.QueryInst(id, testconst.ALIYUN_REGION, client)

		if status == "Running" {
			break
		}
	}

	_logUtils.Infof("%s %s", status, macAddress)

	ip, err := ecsSrv.AllocateIp(id, client)
	vncPassword, _ := ecsSrv.QueryVncPassword(id, testconst.ALIYUN_REGION, client)
	vncUrl, _ := ecsSrv.QueryVncUrl(id, vncPassword, testconst.ALIYUN_REGION, false, client)

	_logUtils.Infof("%s, %s, %s, %s, %s", id, name, ip, vncUrl, vncPassword)

	<-time.After(60 * time.Second)

	err = ecsSrv.RemoveInst(id, client)
	_logUtils.Infof("%#v", err)
}
