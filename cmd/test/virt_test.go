package main

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"testing"
)

func TestVirt(t *testing.T) {
	_logUtils.Init(agentConst.AppName)

	virtService := agentService.NewLibvirtService()
	dom, macAddress, _ := virtService.CloneVm("win10", "win10-01")
	defer dom.Free()

	virtService.StartVm(dom)

	name, _ := dom.GetName()
	_logUtils.Infof("%s: %s", name, macAddress)

}
