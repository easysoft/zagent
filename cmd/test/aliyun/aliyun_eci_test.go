package aliyun

import (
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"testing"
)

func TestAliyunEci(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	commSrv := vendors.NewAliyunCommService()
	eciSrv := vendors.NewAliyunEciService()
	eciSrv.AliyunCommService = commSrv

	eciClient, _ := commSrv.CreateEciClient(serverConst.ALIYUN_ECI_URL, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	ecsClient, _ := commSrv.CreateEcsClient(serverConst.ALIYUN_ECI_URL, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	vpcClient, _ := commSrv.CreateVpcClient(serverConst.ALIYUN_ENS_URL, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)

	eipId, _ := commSrv.GetEip(testconst.ALIYUN_REGION, vpcClient)
	_logUtils.Infof("eipId %s", eipId)

	switchId, _, _ := commSrv.GetSwitch(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, vpcClient)
	_logUtils.Infof("switchId %s", switchId)
	securityGroupId, _ := commSrv.QuerySecurityGroupByVpc(testconst.ALIYUN_VPC, testconst.ALIYUN_REGION, ecsClient)
	_logUtils.Infof("securityGroupId %s", securityGroupId)

	eciId, _ := eciSrv.CreateInst("maven-testng-001", "maven-testng",
		"registry-vpc.cn-hangzhou.aliyuncs.com/com-deeptest/maven-testng",
		[]string{
			//"sleep 6000",
			//"rm -rf ci_test_testng",
			//"git clone https://gitee.com/ngtesting/ci_test_testng.git",
			//"cd ci_test_testng",
			//"mvn clean package > logs.txt",

			"pwd > log.txt",
			"sleep 6000",
		},
		eipId, switchId, securityGroupId, testconst.ALIYUN_REGION, eciClient)

	_logUtils.Infof("eciId %s", eciId)

	eciSrv.Destroy(eciId, testconst.ALIYUN_REGION, eciClient)
}
