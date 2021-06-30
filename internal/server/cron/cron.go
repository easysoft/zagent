package serverCron

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
	"github.com/easysoft/zagent/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ServerCron struct {
	ExecService *service.ExecService `inject:""`
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
			s.ExecService.CheckExec()
			s.ExecService.SetTimeout()
			s.ExecService.RetryTimeoutOrFailed()
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
