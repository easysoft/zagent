package serverService

import (
	"fmt"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	bizCasbin "github.com/easysoft/zagent/internal/server/biz/casbin"
	"github.com/easysoft/zagent/internal/server/biz/transformer"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/fatih/color"
	gf "github.com/snowlyg/gotransformer"
	"strconv"
	"time"
)

type RoleService struct {
	CommonService

	RoleRepo *repo.RoleRepo `inject:""`
	PermRepo *repo.PermRepo `inject:""`

	CasbinService *bizCasbin.CasbinService `inject:""`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// RolePermissions get role's permissions
func (s *RoleService) RolePermissions(role *model.Role) []*model.Permission {
	perms := s.GetPermissionsForUser(role.ID)
	var ps []*model.Permission
	for _, perm := range perms {
		if len(perm) >= 3 && len(perm[1]) > 0 && len(perm[2]) > 0 {
			search := &commDomain.Search{
				Fields: []*commDomain.Filed{
					{
						Key:       "name",
						Condition: "=",
						Value:     perm[1],
					},
					{
						Key:       "act",
						Condition: "=",
						Value:     perm[2],
					},
				},
			}
			p, err := s.PermRepo.GetPermission(search)
			if err == nil && p.ID > 0 {
				ps = append(ps, p)
			}
		}
	}
	return ps
}

func (s *RoleService) RolesTransform(roles []*model.Role) []*transformer.Role {
	var rs []*transformer.Role
	for _, role := range roles {
		r := s.RoleTransform(role)
		rs = append(rs, r)
	}
	return rs
}
func (s *RoleService) RoleTransform(role *model.Role) *transformer.Role {
	transformerRole := &transformer.Role{}
	g := gf.NewTransform(s, role, time.RFC3339)
	_ = g.Transformer()
	transformerRole.Perms = s.PermRepo.PermsTransform(s.RolePermissions(role))
	return transformerRole
}

// CreateRole create role
func (s *RoleService) CreateRole(role *model.Role) error {
	if err := s.RoleRepo.DB.Create(role).Error; err != nil {
		return err
	}

	s.addPerms(role.PermIds, role)

	return nil
}

// UpdateRole update role
func (s *RoleService) UpdateRole(id uint, nr *model.Role) error {
	if err := s.RoleRepo.UpdateObj(&model.Role{}, nr, id); err != nil {
		return err
	}

	s.addPerms(nr.PermIds, nr)

	return nil
}

func (s *RoleService) GetRolesForUser(uid uint) []string {
	uids, err := s.CasbinService.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

// addPerms add perms
func (s *RoleService) addPerms(permIds []uint, role *model.Role) {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := s.CasbinService.Enforcer.DeletePermissionsForUser(roleId); err != nil {
			color.Red(fmt.Sprintf("AppendPermsErr:%s \n", err))
		}
		var perms []model.Permission
		s.RoleRepo.DB.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := s.CasbinService.Enforcer.AddPolicy(roleId, perm.Name, perm.Act); err != nil {
				color.Red(fmt.Sprintf("AddPolicy:%s \n", err))
			}
		}
	} else {
		color.Yellow(fmt.Sprintf("没有角色：%s 权限为空 \n", role.Name))
	}
}
