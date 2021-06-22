package agentConst

type RunMode string

const (
	Host    RunMode = "host"
	Vm      RunMode = "vm"
	Android RunMode = "android"
	Ios     RunMode = "ios"
)

func (e RunMode) ToString() string {
	return string(e)
}
