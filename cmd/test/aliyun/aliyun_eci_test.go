package aliyun

import (
	"fmt"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/vendors/aliyun"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"testing"
)

func TestAliyunEci(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	commSrv := aliyun.NewAliyunCommService()
	eciSrv := aliyun.NewAliyunEciService()
	eciSrv.AliyunCommService = commSrv

	eciClient, _ := commSrv.CreateEciClient(serverConst.ALIYUN_ECI_URL, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)

	ecsUrl := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, testconst.ALIYUN_REGION)
	ecsClient, _ := commSrv.CreateEcsClient(ecsUrl, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	vpcClient, _ := commSrv.CreateVpcClient(ecsUrl, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)

	switchId, _, _ := commSrv.GetSwitch(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, vpcClient)
	_logUtils.Infof("switchId %s", switchId)

	securityGroupId, _ := commSrv.QuerySecurityGroupByVpc(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, ecsClient)
	_logUtils.Infof("securityGroupId %s", securityGroupId)

	eipId, _ := commSrv.CreateEip(testconst.ALIYUN_REGION, vpcClient)
	_logUtils.Infof("eipId %s", eipId)

	if eipId == "" || switchId == "" || securityGroupId == "" {
		msg := fmt.Sprintf("eipId (%s), switchId (%s) or securityGroupId (%s) is empty, cancel.",
			eipId, switchId, securityGroupId)
		_logUtils.Infof(msg)
		return
	}

	eciId, _ := eciSrv.CreateInst(
		"maven-testng-001",
		"registry-vpc.cn-hangzhou.aliyuncs.com/com-deeptest/maven-testng",
		"imc-bp1frwjer5gmb9tp5831",
		[]string{
			"pwd > log.txt",
			"sleep 10",

			"rm -rf ci_test_testng",
			"git clone https://gitee.com/ngtesting/ci_test_testng.git",
			"cd ci_test_testng",
			"mvn clean package > logs.txt",

			"sleep 600",
		},
		eipId, switchId, securityGroupId, testconst.ALIYUN_REGION, eciClient)

	_logUtils.Infof("eciId %s", eciId)

	eciSrv.Destroy(eciId, testconst.ALIYUN_REGION, eciClient)
	commSrv.DestroyEip(eipId, testconst.ALIYUN_REGION, vpcClient)
}
