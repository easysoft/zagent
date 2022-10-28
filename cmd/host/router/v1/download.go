package v1

type DownloadReq struct {
	Urls []string `json:"urls"`

	TaskId int `json:"taskId"`
}

type DownloadCancelReq struct {
	Url string `json:"url"`
}
