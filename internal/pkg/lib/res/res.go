package _resUtils

import (
	_commonUtils "github.com/easysoft/zv/internal/pkg/lib/common"
	agentRes "github.com/easysoft/zv/res/agent-vm"
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
