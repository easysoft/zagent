package _resUtils

import (
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	agentRes "github.com/easysoft/zagent/res/vm"
	"io/ioutil"
)

func ReadRes(path string) (ret []byte, err error) {
	isRelease := _commonUtils.IsRelease()
	if isRelease {
		ret, err = agentRes.Asset(path)
	} else {
		ret, err = ioutil.ReadFile(path)
	}

	return
}
