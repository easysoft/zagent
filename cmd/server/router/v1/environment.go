package v1

import "github.com/easysoft/zagent/internal/comm/const"

type Environment struct {
	OsCategory consts.OsCategory `json:"osCategory" example:"windows"`
	OsType     consts.OsType     `json:"osType" example:"win10"`
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang" example:"zh_cn"`
}
