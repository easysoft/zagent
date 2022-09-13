package aliyun

import (
	_const "github.com/easysoft/zv/cmd/test/_const"
	consts "github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/pkg/vendors/huaweicloud"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"testing"
	"time"
)

func TestHuaweiCloudEcs(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	srv := huaweicloud.NewHuaweiCloudEcsService()
	ecsClient, err := srv.CreateEcsClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	imgClient, err := srv.CreateImgClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	vpcClient, err := srv.CreateVpcClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)

	if err != nil {
		return
	}

	huaweiCloudService := huaweicloud.NewHuaweiCloudEcsService()
	id, name, err := huaweiCloudService.CreateInst(
		"win10-x64-pro-zh_cn-"+_stringUtils.Uuid(), "image-win10-x64-pro-zh_cn", ecsClient, imgClient, vpcClient)

	<-time.After(5 * time.Second)

	name, status, ip, mac, err := huaweiCloudService.QueryInst(id, ecsClient)
	_logUtils.Infof("vm name %s, status %s, ip %s, mac %s", name, status, ip, mac)

	url, err := huaweiCloudService.QueryVnc(id, ecsClient)
	_logUtils.Infof("vm id %s, name %s, url %s", id, name, url)

	//err = huaweiCloudService.RemoveInst(id, ecsClient)

	_logUtils.Infof("%s, %s", id, name)
}
