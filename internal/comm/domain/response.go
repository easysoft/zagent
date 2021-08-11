package domain

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Response struct {
	Raw string `json:"raw" yaml:"raw"`

	Status      string             `json:"status" yaml:"status"`
	Code        int                `json:"code" yaml:"code"`
	PreviewType consts.PreviewType `json:"previewType" yaml:"previewType"`
	Cookies     []Cookie           `json:"cookies" yaml:"cookies"`

	ForType consts.TestType `json:"forType" yaml:"forType"` // interface or case
	ForId   uint            `json:"forId" yaml:"forId"`
}

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
