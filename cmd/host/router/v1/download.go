package v1

type DownloadReq struct {
	Urls []string `json:"urls"`

	Ids []int `json:"ids"`

	Id int `json:"id"`
}
