package v1

type DownloadReq struct {
	Urls []string `json:"urls"`

	TaskId int `json:"taskId"`
}
