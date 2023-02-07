package other

import (
	"log"
	"os"
	"testing"

	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

func TestEnv(t *testing.T) {
	os.Setenv("ABC", "123")

	str, _ := _shellUtils.ExeShell("echo $ABC")
	log.Printf("%s", str)
}
