package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type DownloadReq struct {
	Url        string `json:"url"`
	Md5        string `json:"md5"`
	ZentaoTask int    `json:"zentaoTask"`
}

type DownloadResp struct {
	Path       string            `json:"path"`
	Status     consts.TaskStatus `json:"status"` // Enums consts.TaskStatus
	ZentaoTask int               `json:"zentaoTask"`
}

type DownloadCancelReq struct {
	Url string `json:"url"`
}
