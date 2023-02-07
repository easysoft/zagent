package hostAgentService

import (
	"fmt"

	agentModel "github.com/easysoft/zagent/internal/host/model"
	_db "github.com/easysoft/zagent/pkg/db"
	"github.com/fatih/color"
)

type InitService struct {
}

func NewInitService() {
}

func (s *InitService) InitModels() {
	//if !_commonUtils.IsRelease() {
	err := _db.GetInst().DB().AutoMigrate(
		agentModel.Models...,
	)
	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}
	//}
}
