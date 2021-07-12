package repo

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/model"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"gorm.io/gorm"
)

type HistoryRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewHistoryRepo() *HistoryRepo {
	return &HistoryRepo{}
}

func (r *HistoryRepo) Query(tp consts.EntityType, id uint) (pos []model.History) {
	query := r.DB.Model(&model.History{}).
		Select("*").
		Where("type=? AND id=? AND NOT deleted", tp, id).
		Offset(0).Limit(serverConst.PageSize).Order("id DESC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *HistoryRepo) Get(id uint) (po model.History) {
	r.DB.Model(&model.History{}).Where("id = ?", id).First(&po)

	return
}

func (r *HistoryRepo) Save(po *model.History) (err error) {
	err = r.DB.Model(&model.History{}).Create(po).Error
	return
}
