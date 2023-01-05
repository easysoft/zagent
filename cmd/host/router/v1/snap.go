package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type SnapReq struct {
	Vm   string `json:"vm"`
	Name string `json:"name"`
	Task int    `json:"task"`
}

type SnapResp struct {
	Path   string            `json:"path"`
	Status consts.TaskStatus `json:"status"` // Enums consts.TaskStatus
	Task   int               `json:"task"`
}

type SnapCancelReq struct {
	Id int `json:"id"`
}