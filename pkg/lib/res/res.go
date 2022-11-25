package _resUtils

import (
	"io/ioutil"
	"strings"

	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	hostRes "github.com/easysoft/zagent/res/host"
	agentRes "github.com/easysoft/zagent/res/vm"
)

func ReadRes(path string) (ret []byte, err error) {
	isRelease := _commonUtils.IsRelease()
	if isRelease {
		if strings.Contains(path, "host") {
			ret, err = hostRes.Asset(path)
		} else {
			ret, err = agentRes.Asset(path)
		}
	} else {
		ret, err = ioutil.ReadFile(path)
	}

	return
}
