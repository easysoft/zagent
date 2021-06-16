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
	virtService.GetDomain()
}
