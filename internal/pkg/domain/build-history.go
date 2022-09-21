package domain

import "time"

type BuildHistory struct {
	Id         uint       `json:"id"`
	Progress   string     `json:"progress"`
	Status     string     `json:"status"`
	QueueId    uint       `json:"queueId"`
	NodeIp     string     `json:"nodeIp"`
	VncAddress string     `json:"vncAddress"`
	ResultPath string     `json:"resultPath"`
	CreatedAt  *time.Time `json:"createdAt"`

	OwnerType string `json:"ownerType"`
	OwnerId   uint   `json:"ownerId"`
}
