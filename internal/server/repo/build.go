package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type BuildRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewBuildRepo() *BuildRepo {
	return &BuildRepo{}
}

type buildRepository struct{}

func (r BuildRepo) GetBuild(id uint) (build model.Build) {
	r.DB.Preload("Queue").Where("ID=?", id).First(&build)

	return
}

func (r BuildRepo) Save(build *model.Build) (err error) {
	err = r.DB.Model(&build).
		Omit("StartTime", "CompleteTime").
		Create(&build).Error
	return
}

func (r BuildRepo) Start(build model.Build) (err error) {
	r.DB.Model(&build).Where("id=?", build.ID).Updates(
		map[string]interface{}{"progress": commConst.ProgressInProgress, "start_time": time.Now()})
	return
}

func (r BuildRepo) Delete(build model.Build) (err error) {
	r.DB.Delete(&build)
	return
}

func (r BuildRepo) SaveResult(appiumTestTo commDomain.Build, resultPath string,
	progress commConst.BuildProgress, status commConst.BuildStatus, msg string) {

	r.DB.Model(&model.Build{}).Where("id=?", appiumTestTo.ID).Updates(
		map[string]interface{}{"progress": progress, "status": status, "resultPath": resultPath, "resultMsg": msg,
			"complete_time": time.Now()})
	return
}

func (r BuildRepo) SetTimeoutByQueueId(queueId uint) {
	r.DB.Model(&model.Build{}).Where("queue_id=?", queueId).Updates(
		map[string]interface{}{"progress": commConst.ProgressTimeout})
	return
}