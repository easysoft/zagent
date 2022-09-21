package repo

import (
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type BuildRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewBuildRepo() *BuildRepo {
	return &BuildRepo{}
}

type buildRepository struct{}

func (r BuildRepo) GetBuild(id uint) (build model.Build) {
	r.DB.Model(&model.Build{}).Preload("Queue", "NOT deleted").Where("ID=?", id).First(&build)

	return
}

func (r BuildRepo) Save(build *model.Build) (err error) {
	err = r.DB.Model(&build).
		Omit("StartTime", "EndTime").
		Create(&build).Error
	return
}

func (r BuildRepo) Start(build model.Build) (err error) {
	r.DB.Model(&model.Build{}).Where("id=?", build.ID).Updates(
		map[string]interface{}{"progress": consts.ProgressRunning, "start_time": time.Now()})
	return
}

func (r BuildRepo) Delete(build model.Build) (err error) {
	r.DB.Model(&model.Build{}).Delete(&build)
	return
}

func (r BuildRepo) SaveResult(build domain.Build) (err error) {
	err = r.DB.Model(&model.Build{}).Where("id=?", build.ID).Updates(
		map[string]interface{}{"progress": build.Progress, "status": build.Status,
			"result_path": build.ResultPath, "result_msg": build.ResultMsg,
			"end_time": time.Now()}).Error
	return
}

func (r BuildRepo) SetTimeoutByQueueId(queueId uint) {
	r.DB.Model(&model.Build{}).Where("queue_id=?", queueId).Updates(
		map[string]interface{}{"progress": consts.ProgressTimeout})
	return
}
