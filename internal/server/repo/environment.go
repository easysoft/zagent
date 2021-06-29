package repo

import (
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
)

type EnvironmentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewEnvRepo() *EnvironmentRepo {
	return &EnvironmentRepo{}
}

func (r *EnvironmentRepo) GetMap() (mp map[string]interface{}, err error) {

	return
}

func (r *EnvironmentRepo) Get(id int) (po model.Environment) {
	r.DB.Where("id = ?", id).First(&po)

	return
}
