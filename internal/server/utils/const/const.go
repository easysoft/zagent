package serverConst

type VmPlatform string

const (
	AppName    = "server"
	NluVersion = "2.0"

	WsNamespace   = "default"
	WsEvent       = "OnChat"
	WsDefaultRoom = "square"

	TrainingTimeout = 60 * 60 // sec

	PageSize             = 15
	KvmNative VmPlatform = "kvmNative"
	Pve       VmPlatform = "pve"

	DockerNative ContainerPlatform = "dockerNative"
	Portainer    ContainerPlatform = "portainer"
)

type WsEventAction string

const (
	TaskUpdate WsEventAction = "task_update"
)

func (e WsEventAction) ToString() string {
	return string(e)
}

var (
	SlotTypeAbbrMap = map[string]string{"synonym": "syn", "lookup": "lkp", "regex": "rgx"}
)

type ContainerPlatform string
