package serverCron

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/server/service/testing"
	_cronUtils "github.com/easysoft/zv/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zv/pkg/lib/date"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
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
				consts.ProgressCreated, consts.ProgressResPending: to create vm on host
				consts.ProgressResLaunched:                        to exec testing on vm
			*/
			s.ExecService.QueryForExec()

			/**
			query queue by progress timeout:
				consts.ProgressResPending: consts.WaitResPendingTimeout
				consts.ProgressResLaunched: consts.WaitResReadyTimeout
				consts.ProgressRunning: consts.WaitRunCompletedTimeout
			*/
			s.ExecService.QueryForTimeout()

			/**
				query queue for retry:
					process = consts.ProgressTimeout
						||
			    	status = consts.StatusFail
			*/
			s.ExecService.QueryForRetry()

			s.syncMap.Store("isRunning", false)
			s.syncMap.Store("lastCompletedTime", time.Now().Unix())
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
