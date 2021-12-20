package aliyun

import (
	"github.com/easysoft/zv/cmd/test/_const"
	consts "github.com/easysoft/zv/internal/comm/const"
	_logUtils "github.com/easysoft/zv/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zv/internal/pkg/lib/string"
	"github.com/easysoft/zv/internal/server/service/vendors/huaweicloud"
	"strings"
	"testing"
)

func TestHuaweiCloudCci(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := huaweicloud.NewHuaweiCloudCommService()
	cci := huaweicloud.NewHuaweiCloudCciService()

	client, _ := srv.CreateIamClient(
		testconst.HUAWEI_CLOUD_KEY, testconst.HUAWEI_CLOUD_Secret, testconst.HUAWEI_CLOUD_REGION)

	token, _ := srv.GetIamToken(
		testconst.HUAWEI_CLOUD_USER, testconst.HUAWEI_CLOUD_IAM_USER, testconst.HUAWEI_CLOUD_IAM_PASSWORD,
		client)

	image := "swr.cn-east-3.myhuaweicloud.com/tester-im/maven-testng:1.0"
	name := "maven-testng-" + _stringUtils.NewUuid()
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join([]string{
			"sleep 30",
			"rm -rf ci_test_testng",
			"git clone https://gitee.com/ngtesting/ci_test_testng.git; cd ci_test_testng",
			"mvn clean package > logs.txt",
			"sleep 600",
		}, "; "),
	}

	respCreate, err := cci.Create(name, image, cmd, token, testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE)
	_logUtils.Infof("%#v, %#v", respCreate, err)

	respDestroy, success := cci.Destroy(respCreate.Metadata.Name, token, testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE)
	_logUtils.Infof("%#v, %#v", respDestroy, success)
}
