package repo

import (
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
)

type TmplRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewTmplRepo() *TmplRepo {
	return &TmplRepo{}
}

func (r *TmplRepo) Get(id uint) (po model.VmTmpl) {
	r.DB.Where("id = ?", id).First(&po)

	return
}
