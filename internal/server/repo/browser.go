package repo

import (
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
)

type BrowserRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewBrowserRepo() *BrowserRepo {
	return &BrowserRepo{}
}

func (r BrowserRepo) ListByBacking(backingId uint) (pos []*model.Browser) {
	r.DB.Model(&model.Browser{}).Raw("SELECT * FROM biz_browser WHERE id "+
		"IN (SELECT browser_id FROM biz_backing_browser_r WHERE vm_backing_id = ?)", backingId).
		Scan(&pos)

	return
}
