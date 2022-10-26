package agentModel

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type Download struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	Url    string                  `json:"url"`
	Status consts.DownloadProgress `json:"status"`
	Retry  int                     `json:"retry"`

	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	TimeoutTime *time.Time `json:"timeoutTime"`
}

func (Download) TableName() string {
	return "biz_download"
}
