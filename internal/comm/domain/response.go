package commDomain

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type Response struct {
	Raw string `json:"raw" yaml:"raw"`

	Status      string             `json:"status" yaml:"status"`
	Code        int                `json:"code" yaml:"code"`
	PreviewType _const.PreviewType `json:"previewType" yaml:"previewType"`
	Cookies     []Cookie           `json:"cookies" yaml:"cookies"`

	ForType _const.TestType `json:"forType" yaml:"forType"` // interface or case
	ForId   uint            `json:"forId" yaml:"forId"`
}
