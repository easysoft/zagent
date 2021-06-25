package repo

import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type PlanRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewPlanRepo() *PlanRepo {
	return &PlanRepo{}
}

func (r *PlanRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.Plan, total int64) {
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
	err = r.DB.Model(&model.Plan{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *PlanRepo) Get(id uint) (po model.Plan) {
	r.DB.Where("id = ?", id).First(&po)

	return
}

func (r *PlanRepo) Save(po *model.Plan) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *PlanRepo) Update(po *model.Plan) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *PlanRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.Plan{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.Plan{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *PlanRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.Plan{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *PlanRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Plan{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}
