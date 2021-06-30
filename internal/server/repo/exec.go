package repo

import (
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"strings"
)

type ExecRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewExecRepo() *ExecRepo {
	return &ExecRepo{}
}

func (r ExecRepo) Save(queue model.Queue) (err error) {
	err = r.DB.Model(&queue).Omit("StartTime", "PendingTime").Create(&queue).Error
	return
}

func (r ExecRepo) DeleteInSameGroup(groupId uint, serials []string) (err error) {
	r.DB.Where("group_id=? AND serial IN (?)", groupId, strings.Join(serials, ",")).Delete(&model.Queue{})
	return
}
