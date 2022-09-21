package model

import consts "github.com/easysoft/zv/internal/pkg/const"

type Browser struct {
	BaseModel

	Name    string             `json:"name"`
	Type    consts.BrowserType `json:"type"`
	Version string             `json:"version"`
	Lang    consts.OsLang      `json:"lang"`
}

func (Browser) TableName() string {
	return "biz_browser"
}
