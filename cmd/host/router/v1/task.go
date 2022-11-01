package v1

import agentModel "github.com/easysoft/zagent/internal/host/model"

type ListTaskResp struct {
	Created    []agentModel.Task `json:"created"`
	InProgress []agentModel.Task `json:"inProgress"`

	Canceled  []agentModel.Task `json:"canceled"`
	Completed []agentModel.Task `json:"completed"`
	Failed    []agentModel.Task `json:"failed"`
}
