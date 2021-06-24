package serverCron

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
)

type ServerCron struct {
}

func NewServerCron() *ServerCron {
	inst := &ServerCron{}
	inst.Init()
	return inst
}

func (s *ServerCron) Init() {
	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", _const.WebCheckQueueInterval),
		func() {

		},
	)
}
