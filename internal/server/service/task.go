package service

import (
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type TaskService struct {
	TaskRepo *repo.TaskRepo `inject:""`
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) List(keywords, status string, pageNo int, pageSize int) (pos []model.Task, total int64) {
	pos, total = s.TaskRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *TaskService) Get(id uint) (po model.Task) {
	po = s.TaskRepo.Get(id)
	return
}

func (s *TaskService) Save(po *model.Task, userId uint) (err error) {
	po.UserId = userId
	err = s.TaskRepo.Save(po)

	return
}

func (s *TaskService) Update(po *model.Task) (err error) {
	err = s.TaskRepo.Update(po)

	return
}

func (s *TaskService) SetDefault(id uint) (err error) {
	err = s.TaskRepo.SetDefault(id)

	return
}

func (s *TaskService) Disable(id uint) (err error) {
	err = s.TaskRepo.Disable(id)

	return
}

func (s *TaskService) Delete(id uint) (err error) {
	err = s.TaskRepo.Delete(id)

	return
}
