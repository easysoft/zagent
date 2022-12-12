package vmCron

import (
	"fmt"
	"sync"
	"time"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	vmAgentService "github.com/easysoft/zagent/internal/vm/service"
	_cronUtils "github.com/easysoft/zagent/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zagent/pkg/lib/date"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type CronService struct {
	syncMap   sync.Map
	VmService *vmAgentService.VmService `inject:""`
}

func NewAgentCron() *CronService {
	inst := &CronService{}
	inst.Init()
	return inst
}

func (s *CronService) Init() {
	s.syncMap.Store("isRunning", false)
	s.syncMap.Store("lastCompletedTime", int64(0))

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.AgentCheckExecutionInterval),
		func() {
			s.checkTask()
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})

	time.AfterFunc(2*time.Second, func() { s.checkTask() })
}

func (s *CronService) checkTask() {
	isRunning, _ := s.syncMap.Load("isRunning")
	lastCompletedTime, _ := s.syncMap.Load("lastCompletedTime")

	if isRunning.(bool) || time.Now().Unix()-lastCompletedTime.(int64) < consts.AgentCheckExecutionInterval {
		_logUtils.Infof("skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
		return
	}
	s.syncMap.Store("isRunning", true)

	s.VmService.Check()

	s.syncMap.Store("isRunning", false)
	s.syncMap.Store("lastCompletedTime", time.Now().Unix())
}
