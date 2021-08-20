package aliyun

import (
	_const "github.com/easysoft/zagent/cmd/test/_const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
	"time"
)

func TestHuaweiCloudEcs(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewHuaweiCloudEcsService()
	ecsClient, err := srv.CreateEcsClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	imgClient, err := srv.CreateImgClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	vpcClient, err := srv.CreateVpcClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)

	if err != nil {
		return
	}

	huaweiCloudService := vendors.NewHuaweiCloudEcsService()
	id, name, err := huaweiCloudService.CreateInst(
		"win10-x64-pro-zh_cn-"+_stringUtils.NewUuidWithSep(), "image-win10-x64-pro-zh_cn", ecsClient, imgClient, vpcClient)

	<-time.After(5 * time.Second)

	name, status, ip, mac, err := huaweiCloudService.QueryInst(id, ecsClient)
	_logUtils.Infof("vm name %s, status %s, ip %s, mac %s", name, status, ip, mac)

	url, err := huaweiCloudService.QueryVnc(id, ecsClient)
	_logUtils.Infof("vm id %s, name %s, url %s", id, name, url)

	//err = huaweiCloudService.RemoveInst(id, ecsClient)

	_logUtils.Infof("%s, %s", id, name)
}
