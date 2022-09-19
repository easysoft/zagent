package other

import (
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	"log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("ABC", "123")

	str, _ := _shellUtils.ExeShell("echo $ABC")
	log.Printf("%s", str)
}
