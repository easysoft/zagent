package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type DownloadReq struct {
	Urls []string `json:"urls"`

	ZentaoTask int `json:"zentaoTask"`
}

type DownloadResp struct {
	Path       string            `json:"path"`
	Status     consts.TaskStatus `json:"status"`
	ZentaoTask int               `json:"zentaoTask"`
}

type DownloadCancelReq struct {
	Url string `json:"url"`
}
