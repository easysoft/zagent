package model

import commConst "github.com/easysoft/zagent/internal/comm/const"

type Browser struct {
	BaseModel

	Name    string                `json:"name"`
	Type    commConst.BrowserType `json:"type"`
	Version string                `json:"version"`
	Lang    commConst.OsLang      `json:"lang"`
}

func (Browser) TableName() string {
	return "biz_browser"
}
