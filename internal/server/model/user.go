package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name             string     `gorm:"not null; type:varchar(60)" json:"name" validate:"required,gte=2,lte=50" comment:"用户名"`
	Username         string     `gorm:"username;unique;not null;type:varchar(60)" json:"username" validate:"required,gte=2,lte=50"  comment:"名称"`
	Password         string     `gorm:"type:varchar(100)" json:"password"  comment:"密码"`
	Intro            string     `gorm:"not null; type:varchar(512)" json:"introduction" comment:"简介"`
	Avatar           string     `gorm:"type:longText" json:"avatar"  comment:"头像"`
	Token            string     `gorm:"type:varchar(128)" json:"token" comment:"令牌"`
	TokenUpdatedTime *time.Time `json:"tokenUpdatedTime" comment:"令牌更新时间"`

	RoleIds []uint `gorm:"-" json:"role_ids"  validate:"required" comment:"角色"`

	ProjectId uint `json:"projectId" comment:"用户工作项目"`
}

type Avatar struct {
	Avatar string `type:longText" json:"avatar" validate:"required" comment:"头像"`
}

type Token struct {
	Token      string `json:"token"`
	RememberMe bool   `json:"rememberMe"`
}

func (User) TableName() string {
	return "biz_user"
}
