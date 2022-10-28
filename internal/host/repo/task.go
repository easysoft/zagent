package hostRepo

import (
	agentModel "github.com/easysoft/zv/internal/host/model"
	"github.com/easysoft/zv/internal/pkg/const"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"gorm.io/gorm"
	"time"
)

type TaskRepo struct {
	DB *gorm.DB `inject:""`
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) Query() (pos []agentModel.Task, err error) {
	err = r.DB.Model(&agentModel.Task{}).
		Where("NOT deleted").
		Find(&pos).Error

	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *TaskRepo) Get(id uint) (po agentModel.Task) {
	r.DB.Model(&agentModel.Task{}).Preload("Environments", "NOT deleted").Where("id = ?", id).First(&po)

	return
}
func (r *TaskRepo) GetDetail(id uint) (po agentModel.Task) {
	r.DB.Model(&po).
		Where("id = ?", id).First(&po)

	return
}

func (r *TaskRepo) GetByUrl(url string) (po agentModel.Task, err error) {
	r.DB.Model(&po).
		Where("url = ?", url).First(&po)

	return
}

func (r *TaskRepo) Save(po *agentModel.Task) (err error) {
	err = r.DB.Model(&agentModel.Task{}).Create(&po).Error
	return
}

func (r *TaskRepo) Update(po *agentModel.Task) (err error) {
	err = r.DB.Model(&agentModel.Task{}).Where("id = ?", po.ID).
		Session(&gorm.Session{FullSaveAssociations: true}).Updates(&po).Error
	return
}

func (r *TaskRepo) UpdateStatus(id uint, filePath, xmlDesc string, status consts.TaskStatus) (err error) {
	err = r.DB.Model(&agentModel.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"status": status, "end_time": time.Now()}).Error

	if filePath != "" {
		err = r.DB.Model(&agentModel.Task{}).Where("id = ?", id).
			Updates(map[string]interface{}{"path": filePath}).Error
	}

	return
}

func (r *TaskRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&agentModel.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *TaskRepo) SetFailed(po agentModel.Task) (err error) {
	r.DB.Model(&agentModel.Task{}).Where("id=?", po.ID).Updates(
		map[string]interface{}{"status": consts.Failed, "timeout_time": time.Now()})
	return
}
