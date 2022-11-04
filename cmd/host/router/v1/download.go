package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type DownloadReq struct {
	Url  string `json:"url"`
	Md5  string `json:"md5"`
	Task int    `json:"task"`
}

type DownloadResp struct {
	Path   string            `json:"path"`
	Status consts.TaskStatus `json:"status"` // Enums consts.TaskStatus
	Task   int               `json:"task"`
}

type DownloadCancelReq struct {
	Id int `json:"id"`
}
