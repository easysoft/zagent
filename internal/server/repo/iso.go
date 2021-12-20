package repo

import (
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
)

type IsoRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func (r IsoRepo) Get(id uint) (iso model.Iso) {
	r.DB.Model(&model.Iso{}).Where("id=?", id).First(&iso)
	return
}
