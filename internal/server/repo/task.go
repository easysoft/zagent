package repo

import (
	"github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type TaskRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.Task, total int64) {
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
	err = r.DB.Model(&model.Task{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *TaskRepo) Get(id uint) (po model.Task) {
	r.DB.Preload("Environments").Where("id = ?", id).First(&po)

	return
}

func (r *TaskRepo) Save(po *model.Task) (err error) {
	err = r.DB.Model(&po).Create(&po).Error
	return
}

func (r *TaskRepo) Update(po *model.Task) (err error) {
	err = r.DB.Where("task_id = ?", po.ID).Delete(&model.Environment{}).Error
	err = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&po).Error
	return
}

func (r *TaskRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.Task{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.Task{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *TaskRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *TaskRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *TaskRepo) SetProgress(taskId uint, progress consts.BuildProgress) (err error) {
	var data map[string]interface{}
	if progress == consts.ProgressInProgress {
		data = map[string]interface{}{"progress": progress, "start_time": time.Now()}
	} else {
		data = map[string]interface{}{"progress": progress, "pending_time": time.Now()}
	}

	r.DB.Model(model.Task{}).Where("id=?", taskId).Updates(data)
	return
}

func (r *TaskRepo) SetResult(taskId uint, progress consts.BuildProgress, status consts.BuildStatus) (err error) {
	var data = map[string]interface{}{"progress": progress, "result": status, "updatedTime": time.Now()}
	r.DB.Model(model.Task{}).Where("id=?", taskId).Updates(data)
	return
}
