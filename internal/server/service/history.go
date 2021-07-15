package serverService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type HistoryService struct {
	HistoryRepo *repo.HistoryRepo `inject:""`
	QueueRepo   *repo.QueueRepo   `inject:""`

	QueueService *QueueService `inject:""`
}

func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

func (s *HistoryService) List(tp consts.EntityType, id uint) (pos []model.History) {
	pos = s.HistoryRepo.Query(tp, id)
	return
}

func (s *HistoryService) Get(id uint) (po model.History) {
	po = s.HistoryRepo.Get(id)
	return
}
func (s *HistoryService) Create(tp consts.EntityType, id, queueId uint, progress consts.BuildProgress, status string) (err error) {
	po := model.NewHistoryPo(tp, id, queueId, progress, status)
	err = s.HistoryRepo.Save(&po)

	return
}

func (s *HistoryService) GetBuildHistoriesByTask(taskId uint) (histories []domain.BuildHistory) {
	histories = s.HistoryRepo.GetBuildHistoriesByTask(taskId)

	return
}
