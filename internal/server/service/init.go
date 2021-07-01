package serverService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/pkg/db"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/fatih/color"
)

type InitService struct {
	SeederService *SeederService `inject:""`
}

func NewInitService() {
}

func (s *InitService) Init() {
	if !_commonUtils.IsRelease() {
		err := _db.GetInst().DB().AutoMigrate(
			model.Models...,
		)
		if err != nil {
			color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
		}

		s.SeederService.AddPerms()
	}
}
