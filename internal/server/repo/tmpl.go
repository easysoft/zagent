package repo

import (
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
)

type TmplRepo struct {
	BaseRepo
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewTmplRepo() *TmplRepo {
	return &TmplRepo{}
}

func (r *TmplRepo) Get(id uint) (po model.VmTmpl) {
	r.DB.Model(&model.VmTmpl{}).Where("id = ?", id).First(&po)

	return
}

func (r TmplRepo) QueryByOs(osCategory consts.OsCategory, osType consts.OsType, osLang consts.OsLang) (
	templId uint, found bool) {

	asserts := make([]domain.VmAssert, 0)
	r.DB.Model(model.VmTmpl{}).
		Where("NOT disabled AND NOT deleted").Order("id ASC").
		Find(&asserts)

	templIds := make([]uint, 0)
	templIds, found = r.FindAssetByOs(osCategory, osType, osLang, asserts, nil)
	if len(templIds) > 0 {
		templId = templIds[0]
		found = true
	}

	return
}
