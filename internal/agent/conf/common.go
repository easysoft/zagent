package agentConf

import (
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
)

func IsDeviceAgent() bool {
	return IsIosAgent() || IsAndroidAgent()
}

func IsAndroidAgent() bool {
	return Inst.RunMode == agentConst.Android
}

func IsIosAgent() bool {
	return Inst.RunMode == agentConst.Ios
}

func IsHostAgent() bool {
	return Inst.RunMode == agentConst.Host
}
func IsVmAgent() bool {
	return Inst.RunMode == agentConst.Vm
}
