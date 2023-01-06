package v1

import consts "github.com/easysoft/zagent/internal/pkg/const"

type SnapTaskReq struct {
	Vm   string          `json:"vm"`
	Name string          `json:"name"`
	Type consts.TaskType `json:"type"`
	Task int             `json:"task"`
}

type SnapTaskResp struct {
	Path   string            `json:"path"`
	Status consts.TaskStatus `json:"status"` // Enums consts.TaskStatus
	Task   int               `json:"task"`
}

type SnapCancelReq struct {
	Id int `json:"id"`
}

type SnapItemResp struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
