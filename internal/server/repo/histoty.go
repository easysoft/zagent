package repo

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
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

func (r *HistoryRepo) GetBuildHistoriesByTask(taskId uint) (histories []domain.BuildHistory) {
	r.DB.Model(&model.Browser{}).Raw(
		`SELECT his.id, his.queue_id,his.progress, his.status, 
			 his.owner_type, his.owner_id, his.created_at,
             bld.result_path, vm.node_ip, vm.vnc_address 

             FROM biz_history his 
			 LEFT JOIN biz_build bld ON bld.id = his.owner_id 
			 LEFT JOIN biz_queue queue ON queue.id = his.queue_id
			 LEFT JOIN biz_vm vm ON vm.id = queue.vm_id 
			 WHERE his.queue_id IN (
				SELECT id FROM biz_queue 
			   WHERE task_id = ?
			 )
			 ORDER BY his.id ASC`,
		taskId).
		Scan(&histories)

	return
}
