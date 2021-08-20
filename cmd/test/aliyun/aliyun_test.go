package aliyun

import (
	"fmt"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestAliyun(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewAliyunEcsService()
	url := fmt.Sprintf("ecs-%s.aliyuncs.com", testconst.ALIYUN_REGION)
	client, err := srv.CreateClient(url, testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	if err != nil {
		return
	}

	id, name, _ := srv.CreateInst("vm-001", "tmpl-ubt20-x64-desktop-zh_cn.qcow2", client)

	vnc, _ := srv.QueryVnc(id, client)

	_logUtils.Infof("%s, %s, %s", id, name, vnc)
}
