package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
)

type BackingRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewBackingRepo() *BackingRepo {
	return &BackingRepo{}
}

func (r BackingRepo) Get(id uint) (image model.VmBacking) {
	r.DB.Where("id=?", id).First(&image)
	return
}

func (r BackingRepo) QueryByOs(osCategory commConst.OsCategory, osType commConst.OsType, osLang commConst.OsLang) (ids []int) {
	var db = r.DB.Model(model.VmBacking{}).Where("NOT disabled AND NOT deleted")
	if osCategory != "" {
		db.Where("osPlatform = ?", osCategory)
	}
	if osType != "" {
		db.Where("osType = ?", osType)
	}
	if osCategory != "" {
		db.Where("osLang = ?", osLang)
	}

	db.Order("id ASC, createdAt ASC").Find(&ids)

	return
}

func (r BackingRepo) QueryByBrowser(browserType commConst.BrowserType, version string) (ids []int) {

	sql := "SELECT r.backingImageId id " +
		"FROM BizBackingImageCapability_Relation r " +
		"LEFT JOIN BizBrowser browser ON browser.id = r.browserId " +
		"WHERE NOT browser.disabled AND NOT browser.deleted "
	if browserType != "" {
		sql += "AND browser.type = ? "
	}
	if version != "" {
		sql += "AND browser.version = ? "
	}

	sql += "ORDER BY id"
	r.DB.Raw(sql).Scan(&ids)

	return
}
