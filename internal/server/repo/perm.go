package repo

import (
	"fmt"
	commDomain "github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/server/biz/transformer"
	"github.com/easysoft/zv/internal/server/model"
	gf "github.com/snowlyg/gotransformer"
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type PermRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewPermRepo() *PermRepo {
	return &PermRepo{}
}

func (r *PermRepo) NewPermission() *model.Permission {
	return &model.Permission{}
}

// GetPermission get permission
func (r *PermRepo) GetPermission(search *commDomain.Search) (*model.Permission, error) {
	t := r.NewPermission()
	err := r.Found(search).First(t).Error
	if !r.IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// DeletePermissionById del permission by id
func (r *PermRepo) DeletePermissionById(id uint) error {
	p := r.NewPermission()
	p.ID = id
	if err := r.DB.Delete(p).Error; err != nil {
		color.Red(fmt.Sprintf("DeletePermissionByIdError:%s \n", err))
		return err
	}
	return nil
}

// GetAllPermissions get all permissions
func (r *PermRepo) GetAllPermissions(s *commDomain.Search) ([]*model.Permission, int64, error) {
	var permissions []*model.Permission
	var count int64
	all := r.GetAll(&model.Permission{}, s)

	all = all.Scopes(r.Relation(s.Relations))

	if err := all.Count(&count).Error; err != nil {
		return nil, count, err
	}

	all = all.Scopes(r.Paginate(s.Offset, s.Limit))

	if err := all.Find(&permissions).Error; err != nil {
		return nil, count, err
	}

	return permissions, count, nil
}

// CreatePermission create permission
func (r *PermRepo) CreatePermission(perm *model.Permission) error {
	if err := r.DB.Create(perm).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePermission update permission
func (r *PermRepo) UpdatePermission(id uint, pj *model.Permission) error {
	if err := r.UpdateObj(&model.Permission{}, pj, id); err != nil {
		return err
	}
	return nil
}

func (r *PermRepo) PermsTransform(perms []*model.Permission) []*transformer.Permission {
	var rs []*transformer.Permission
	for _, perm := range perms {
		r := r.PermTransform(perm)
		rs = append(rs, r)
	}
	return rs
}

func (r *PermRepo) PermTransform(perm *model.Permission) *transformer.Permission {
	p := &transformer.Permission{}
	g := gf.NewTransform(r, perm, time.RFC3339)
	_ = g.Transformer()
	return p
}
