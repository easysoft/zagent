package hostRepo

import (
	"time"

	agentModel "github.com/easysoft/zagent/internal/host/model"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"gorm.io/gorm"
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

func (r *TaskRepo) Get(id uint) (po agentModel.Task, err error) {
	r.DB.Model(&agentModel.Task{}).Preload("Environments", "NOT deleted").Where("id = ?", id).First(&po)

	return
}
func (r *TaskRepo) GetDetail(id uint) (po agentModel.Task, err error) {
	r.DB.Model(&po).
		Where("id = ?", id).First(&po)

	return
}

func (r *TaskRepo) GetByUrl(url string) (po agentModel.Task, err error) {
	r.DB.Model(&po).
		Where("url = ?", url).First(&po)

	return
}

func (r *TaskRepo) GetByMd5(md5 string) (po agentModel.Task, err error) {
	if md5 == "" {
		return
	}

	r.DB.Model(&po).
		Where("md5 = ?", md5).Order("id desc").First(&po)

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

func (r *TaskRepo) UpdateSpeed(id uint, speed float64) (err error) {
	err = r.DB.Model(&agentModel.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"speed": speed}).Error
	return
}

func (r *TaskRepo) UpdateStatus(id uint, filePath string, completionRate float64, xmlDesc string,
	status consts.TaskStatus, isStart, isEnd bool) (err error) {

	updates := map[string]interface{}{"status": status, "xml": xmlDesc}

	if filePath != "" {
		updates["path"] = filePath
	}

	if completionRate > 0 {
		updates["completion_rate"] = completionRate
	}

	if isStart {
		updates["start_time"] = time.Now()
	}
	if isEnd {
		updates["end_time"] = time.Now()
	}

	err = r.DB.Model(&agentModel.Task{}).Where("id = ?", id).
		Updates(updates).Error

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

func (r *TaskRepo) AddRetry(po agentModel.Task) (err error) {
	r.DB.Model(&agentModel.Task{}).Where("id=?", po.ID).Updates(
		map[string]interface{}{"retry": gorm.Expr("retry + ?", 1)})
	return
}
