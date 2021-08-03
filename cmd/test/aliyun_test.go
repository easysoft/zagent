package main

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"testing"
)

func TestAliyun(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	id, name, _ := vendors.NewAliyunService().GetRegions(
		"LTAI5t9ABAZtGob7m7DsTTmg", "OxsHmMVOdbbaDXEjmAqPWGpkq0DLXn")

	_logUtils.Infof("found region: %s, %s", id, name)
}
