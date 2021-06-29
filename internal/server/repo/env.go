package repo

import (
	"gorm.io/gorm"
)

type EnvRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewEnvRepo() *EnvRepo {
	return &EnvRepo{}
}

func (r *EnvRepo) GetMap() (mp map[string]interface{}, err error) {

	return
}
