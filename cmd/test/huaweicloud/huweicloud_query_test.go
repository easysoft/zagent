package aliyun

import (
	_const "github.com/easysoft/zv/cmd/test/_const"
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/vendors/huaweicloud"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"testing"
)

func TestHuaweiCloudQuery(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	srv := huaweicloud.NewHuaweiCloudEcsService()
	ecsClient, err := srv.CreateEcsClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	if err != nil {
		return
	}

	id := "6eca6332-5c8c-4cae-9d62-b17bcf806d2d"

	huaweiCloudService := huaweicloud.NewHuaweiCloudEcsService()
	name, status, ip, mac, err := huaweiCloudService.QueryInst(id, ecsClient)
	_logUtils.Infof("vm name %s, status %s, ip %s, mac %s", name, status, ip, mac)

	url, err := huaweiCloudService.QueryVnc(id, ecsClient)
	_logUtils.Infof("vm id %s, name %s, url %s", id, name, url)

	//err = huaweiCloudService.RemoveInst(id, ecsClient)

	_logUtils.Infof("%s, %s", id, name)
}
