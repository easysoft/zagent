package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
)

type BackingRepo struct {
	BaseRepo
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewBackingRepo() *BackingRepo {
	return &BackingRepo{}
}

func (r BackingRepo) ListAll() (pos []*model.VmBacking) {
	r.DB.Where("NOT Deleted AND NOT Disabled").Find(&pos)
	return
}

func (r BackingRepo) Get(id uint) (image model.VmBacking) {
	r.DB.Where("id=?", id).First(&image)
	return
}

func (r BackingRepo) QueryByOs(osCategory commConst.OsCategory, osType commConst.OsType, osLang commConst.OsLang,
	backingIdsByBrowser []int) (backingIds []int, found bool) {

	asserts := make([]commDomain.VmAssert, 0)
	r.DB.Model(model.VmBacking{}).
		Where("NOT disabled AND NOT deleted").Order("id ASC").
		Scan(&asserts)

	backingIds, found = r.FindAssetByOs(osCategory, osType, osLang, asserts, backingIdsByBrowser)

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
