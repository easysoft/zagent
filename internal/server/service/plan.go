package service

import (
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type PlanService struct {
	PlanRepo *repo.PlanRepo `inject:""`
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (s *PlanService) List(keywords, status string, pageNo int, pageSize int) (pos []model.Plan, total int64) {
	pos, total = s.PlanRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *PlanService) Get(id uint) (po model.Plan) {
	po = s.PlanRepo.Get(id)
	return
}

func (s *PlanService) Save(po *model.Plan, userId uint) (err error) {
	po.UserId = userId
	err = s.PlanRepo.Save(po)

	return
}

func (s *PlanService) Update(po *model.Plan) (err error) {
	err = s.PlanRepo.Update(po)

	return
}

func (s *PlanService) SetDefault(id uint) (err error) {
	err = s.PlanRepo.SetDefault(id)

	return
}

func (s *PlanService) Disable(id uint) (err error) {
	err = s.PlanRepo.Disable(id)

	return
}

func (s *PlanService) Delete(id uint) (err error) {
	err = s.PlanRepo.Delete(id)

	return
}
