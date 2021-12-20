package domain

import (
	"github.com/easysoft/zv/internal/comm/const"
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
