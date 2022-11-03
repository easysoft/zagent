package agentModel

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type Task struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`

	// for download
	Url            string            `json:"url,omitempty"`
	Md5            string            `json:"md5,omitempty"`
	Path           string            `json:"path,omitempty"`
	Status         consts.TaskStatus `json:"status"`
	Retry          int               `json:"retry"`
	CompletionRate float64           `json:"completionRate"`
	Speed          float64           `json:"speed"`

	// for export vm
	Vm      string `json:"vm,omitempty"`
	Backing string `json:"backing,omitempty"`
	Xml     string `json:"xml,omitempty"`
	//Path    string `json:"path,omitempty"`

	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	TimeoutTime *time.Time `json:"timeoutTime,omitempty"`

	ZentaoTask int             `json:"zentaoTask"`
	TaskType   consts.TaskType `json:"taskType"`
}

func (Task) TableName() string {
	return "biz_task"
}
