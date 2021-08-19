package serverService

import (
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/biz/transformer"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermService() *PermService {
	return &PermService{}
}

func (s *PermService) GetPermission(id uint) (perm *model.Permission, err error) {
	search := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}

	perm, _ = s.PermRepo.GetPermission(search)

	return
}

func (s *PermService) CreatePermission(perm *model.Permission) (err error) {
	err = s.PermRepo.CreatePermission(perm)
	return
}

func (s *PermService) UpdatePermission(id uint, aul *model.Permission) (err error) {
	err = s.PermRepo.UpdatePermission(id, aul)
	return
}

func (s *PermService) DeletePermissionById(id uint) (err error) {
	err = s.PermRepo.DeletePermissionById(id)
	return
}

func (s *PermService) GetAllPermissions(search *domain.Search) (
	permissions []*model.Permission, count int64, err error) {

	permissions, count, err = s.PermRepo.GetAllPermissions(search)
	return
}

func (s *PermService) PermsTransform(perms []*model.Permission) (trans []*transformer.Permission) {
	trans = s.PermRepo.PermsTransform(perms)
	return
}

func (s *PermService) PermTransform(perm *model.Permission) (err error) {
	err = s.PermRepo.CreatePermission(perm)
	return
}
