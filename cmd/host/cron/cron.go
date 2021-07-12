package hostCron

import (
	"fmt"
	hostKvmService "github.com/easysoft/zagent/internal/agent-host/service/kvm"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"sync"
	"time"
)

type CronService struct {
	syncMap          sync.Map
	CheckHostService *hostKvmService.HostService `inject:""`
}

func NewAgentCron() *CronService {
	inst := &CronService{}
	inst.Init()
	return inst
}

func (s *CronService) Init() {
	s.syncMap.Store("isRunning", false)

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.AgentCheckInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			if isRunning.(bool) {
				_logUtils.Infof("is running, skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
				return
			}
			s.syncMap.Store("isRunning", true)

			s.CheckHostService.Register()

			s.syncMap.Store("isRunning", false)
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
