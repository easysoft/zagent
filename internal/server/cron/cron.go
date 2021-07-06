package serverCron

import (
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
	"github.com/easysoft/zagent/internal/server/service/testing"
	"github.com/kataras/iris/v12"
)

type ServerCron struct {
	ExecService *testing.ExecService `inject:""`
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
			/**
			query queue by status:
				consts.ProgressCreated, consts.ProgressPendingRes: to create vm on host
				consts.ProgressLaunchVm:                        to exec testing on vm
			*/
			s.ExecService.CheckExec()

			/**
			query queue by progress timeout:
				consts.ProgressPendingRes: consts.WaitResPendingTimeout
				consts.ProgressLaunchVm: consts.WaitForVmReadyTimeout
				consts.ProgressRunning: consts.WaitTestCompletedTimeout
			*/
			s.ExecService.CheckTimeout()

			/**
				query queue for retry:
					process = consts.ProgressTimeout ||
			    	status = consts.StatusFail
			*/
			s.ExecService.RetryTimeoutOrFailed()
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
