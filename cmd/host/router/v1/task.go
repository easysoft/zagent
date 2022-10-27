package v1

import consts "github.com/easysoft/zv/internal/pkg/const"

type TaskResp struct {
	Status consts.DownloadStatus `json:"status"`

	ImagePath string `json:"path"`

	TaskId   int             `json:"taskId"`
	TaskType consts.TaskType `json:"taskType"`
}
