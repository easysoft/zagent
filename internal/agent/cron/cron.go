package agentCron

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"sync"
	"time"
)

type CronService struct {
	syncMap      sync.Map
	CheckService *agentService.CheckService `inject:""`
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

			s.CheckService.Check()

			s.syncMap.Store("isRunning", false)
		},
	)
}
