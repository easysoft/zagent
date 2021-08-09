package serverConst

const (
	WsNamespace   = "default"
	WsEvent       = "OnChat"
	WsDefaultRoom = "square"

	TrainingTimeout = 60 * 60 // sec

	PageSize = 15
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
