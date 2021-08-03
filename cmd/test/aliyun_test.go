package main

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestAliyun(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewAliyunService()

	client, err := srv.CreateClient("ecs-cn-hangzhou.aliyuncs.com", ALIYUN_KEY, ALIYUN_Secret)
	if err != nil {
		return
	}

	id, name, _ := vendors.NewAliyunService().CreateInst("windows", "x86_64", client)

	_logUtils.Infof("%s, %s", id, name)
}
