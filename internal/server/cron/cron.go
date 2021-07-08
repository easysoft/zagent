package serverCron

import (
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/testing"
	"github.com/kataras/iris/v12"
	"sync"
	"time"
)

type ServerCron struct {
	syncMap     sync.Map
	ExecService *testing.ExecService `inject:""`
}

func NewServerCron() *ServerCron {
	inst := &ServerCron{}
	inst.Init()
	return inst
}

func (s *ServerCron) Init() {
	s.syncMap.Store("isRunning", false)

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.WebCheckQueueInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			if isRunning.(bool) {
				_logUtils.Infof("is running, skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
				return
			}

			s.syncMap.Store("isRunning", true)

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

			s.syncMap.Store("isRunning", false)
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
