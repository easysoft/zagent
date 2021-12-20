package repo

import (
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
)

type EnvironmentRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewEnvRepo() *EnvironmentRepo {
	return &EnvironmentRepo{}
}

func (r *EnvironmentRepo) GetMap() (mp map[string]interface{}, err error) {

	return
}

func (r *EnvironmentRepo) Get(id int) (po model.Environment) {
	r.DB.Model(&model.Environment{}).Where("id = ?", id).First(&po)

	return
}
