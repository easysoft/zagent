package aliyun

import (
	_const "github.com/easysoft/zagent/cmd/test/_const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestHuaweiCloudQuery(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewHuaweiCloudEcsService()
	ecsClient, err := srv.CreateEcsClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)
	if err != nil {
		return
	}

	id := "6eca6332-5c8c-4cae-9d62-b17bcf806d2d"

	huaweiCloudService := vendors.NewHuaweiCloudEcsService()
	name, status, ip, mac, err := huaweiCloudService.QueryInst(id, ecsClient)
	_logUtils.Infof("vm name %s, status %s, ip %s, mac %s", name, status, ip, mac)

	url, err := huaweiCloudService.QueryVnc(id, ecsClient)
	_logUtils.Infof("vm id %s, name %s, url %s", id, name, url)

	//err = huaweiCloudService.RemoveInst(id, ecsClient)

	_logUtils.Infof("%s, %s", id, name)
}
