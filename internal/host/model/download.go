package agentModel

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type Task struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	// for download
	Url    string            `json:"url"`
	Md5    string            `json:"md5"`
	Path   string            `json:"path"`
	Status consts.TaskStatus `json:"status"`
	Retry  int               `json:"retry"`

	// for export vm
	Vm      string `json:"vm"`
	Backing string `json:"backing"`
	Xml     string `json:"xml"`
	//Path    string `json:"path"`

	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	TimeoutTime *time.Time `json:"timeoutTime"`

	ZentaoTask int             `json:"zentaoTask"`
	TaskType   consts.TaskType `json:"taskType"`
}

func (Task) TableName() string {
	return "biz_task"
}
