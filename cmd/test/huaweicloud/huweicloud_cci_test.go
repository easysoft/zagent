package aliyun

import (
	_const "github.com/easysoft/zagent/cmd/test/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestHuaweiCloudCci(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewHuaweiCloudCommService()
	client, err := srv.CreateIamClient(_const.HUAWEI_CLOUD_KEY, _const.HUAWEI_CLOUD_Secret, _const.HUAWEI_CLOUD_REGION)

	token, err := srv.GetIamToken(client)

	_logUtils.Infof("%s, %#v", token, err)
}
