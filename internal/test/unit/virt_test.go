package main

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	"testing"
)

func TestServer(t *testing.T) {
	virtService := agentService.NewLibvirtService()
	virtService.GetDomain()
}
