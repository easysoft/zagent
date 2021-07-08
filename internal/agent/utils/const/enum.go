package agentConst

type RunMode string

const (
	Host RunMode = "host"
	Vm   RunMode = "vm"

	Machine RunMode = "machine"
	Android RunMode = "android"
	Ios     RunMode = "ios"
)

func (e RunMode) ToString() string {
	return string(e)
}
