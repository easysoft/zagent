package agentModel

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type Download struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	Url    string                `json:"url"`
	Md5    string                `json:"md5"`
	Path   string                `json:"path"`
	Status consts.DownloadStatus `json:"status"`
	Retry  int                   `json:"retry"`

	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	TimeoutTime *time.Time `json:"timeoutTime"`

	TaskId int `json:"taskId"`
}

func (Download) TableName() string {
	return "biz_download"
}
