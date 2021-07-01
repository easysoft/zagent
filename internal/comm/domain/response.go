package domain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type Response struct {
	Raw string `json:"raw" yaml:"raw"`

	Status      string                `json:"status" yaml:"status"`
	Code        int                   `json:"code" yaml:"code"`
	PreviewType commConst.PreviewType `json:"previewType" yaml:"previewType"`
	Cookies     []Cookie              `json:"cookies" yaml:"cookies"`

	ForType commConst.TestType `json:"forType" yaml:"forType"` // interface or case
	ForId   uint               `json:"forId" yaml:"forId"`
}
