package repo

import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type ProjectRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.Project, total int64) {
	query := r.DB.Select("*").Order("id ASC")
	if status == "true" {
		query = query.Where("NOT disabled")
	} else if status == "false" {
		query = query.Where("disabled")
	}

	if keywords != "" {
		query = query.Where("name LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}
	query = query.Where("NOT deleted")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.Project{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *ProjectRepo) Get(id uint) (po model.Project) {
	r.DB.Where("id = ?", id).First(&po)

	return
}

func (r *ProjectRepo) Save(po *model.Project) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *ProjectRepo) Update(po *model.Project) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *ProjectRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.Project{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.Project{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *ProjectRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *ProjectRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}
