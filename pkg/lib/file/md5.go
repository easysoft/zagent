package _fileUtils

import (
	"strings"

	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

func GetMd5(pth string) (ret string, err error) {
	cmdStr := ""
	if _commonUtils.IsWin() {
		cmdStr = "CertUtil -hashfile " + pth + " MD5"
	} else {
		cmdStr = "md5sum " + pth + " | awk '{print $1}'"
	}

	ret, _ = _shellUtils.ExeSysCmd(cmdStr)

	if _commonUtils.IsWin() {
		arr := strings.Split(ret, "\n")
		if len(arr) > 1 {
			ret = arr[1]
		}
	}

	ret = strings.TrimSpace(ret)

	return
}
