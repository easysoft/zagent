package aliyun

import (
	_const "github.com/easysoft/zagent/cmd/test/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestHuaweiCloud(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewHuaweiCloudService()

	ecsClient, err := srv.CreateEcsClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGIN)
	imgClient, err := srv.CreateImgClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGIN)
	vpcClient, err := srv.CreateVpcClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGIN)

	if err != nil {
		return
	}

	id, name, _ := vendors.NewHuaweiCloudService().CreateInst(
		"win10-"+_stringUtils.NewUuidWithSep(), "win10", ecsClient, imgClient, vpcClient)

	_logUtils.Infof("%s, %s", id, name)
}
