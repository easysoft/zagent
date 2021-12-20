package agentUtils

import (
	agentConf "github.com/easysoft/zv/internal/agent/conf"
	agentConst "github.com/easysoft/zv/internal/agent/utils/const"
)

func IsDeviceAgent() bool {
	return IsIosAgent() || IsAndroidAgent()
}

func IsAndroidAgent() bool {
	return agentConf.Inst.RunMode == agentConst.Android
}

func IsIosAgent() bool {
	return agentConf.Inst.RunMode == agentConst.Ios
}

func IsHostAgent() bool {
	return agentConf.Inst.RunMode == agentConst.Host
}
func IsVmAgent() bool {
	return agentConf.Inst.RunMode == agentConst.Vm
}
