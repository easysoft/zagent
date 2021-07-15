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
	s.syncMap.Store("lastCompletedTime", int64(0))

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.WebCheckInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			lastCompletedTime, _ := s.syncMap.Load("lastCompletedTime")

			if isRunning.(bool) || time.Now().Unix()-lastCompletedTime.(int64) < consts.WebCheckInterval {
				_logUtils.Infof("skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
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
					process = consts.ProgressTimeout
						||
			    	status = consts.StatusFail
			*/
			s.ExecService.CheckRetry()

			s.syncMap.Store("isRunning", false)
			s.syncMap.Store("lastCompletedTime", time.Now().Unix())
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
