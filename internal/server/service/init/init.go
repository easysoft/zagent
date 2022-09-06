package initService

import (
	"fmt"
	"github.com/easysoft/zv/internal/server/model"
	_db "github.com/easysoft/zv/pkg/db"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	"github.com/fatih/color"
)

type InitService struct {
	SeederService *SeederService `inject:""`
}

func NewInitService() {
}

func (s *InitService) InitDataIfNeeded() {
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
