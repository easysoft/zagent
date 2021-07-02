package main

import (
	_shellUtils "github.com/easysoft/zagent/internal/pkg/lib/shell"
	"github.com/smallnest/rpcx/log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("ABC", "123")

	str, _ := _shellUtils.ExeShell("echo $ABC")
	log.Infof("%s", str)
}
