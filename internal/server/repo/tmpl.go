package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
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
	r.DB.Where("id = ?", id).First(&po)

	return
}

func (r TmplRepo) QueryByOs(osCategory commConst.OsCategory, osType commConst.OsType, osLang commConst.OsLang) (
	templId uint, found bool) {

	asserts := make([]commDomain.VmAssert, 0)
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
