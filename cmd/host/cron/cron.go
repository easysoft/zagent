package hostCron

import (
	"fmt"
	"sync"
	"time"

	hostAgentService "github.com/easysoft/zagent/internal/host/service"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_cronUtils "github.com/easysoft/zagent/pkg/lib/cron"
	"github.com/kataras/iris/v12"
)

type CronService struct {
	syncMap     sync.Map
	HostService *hostAgentService.HostService `inject:""`

	TaskService *hostAgentService.TaskService `inject:""`
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
		"execution",
		fmt.Sprintf("@every %ds", consts.AgentCheckExecutionInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			lastCompletedTime, _ := s.syncMap.Load("lastCompletedTime")

			if isRunning.(bool) || time.Now().Unix()-lastCompletedTime.(int64) < consts.AgentCheckExecutionInterval {
				//_logUtils.Infof("skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
				return
			}
			s.syncMap.Store("isRunning", true)

			//
			s.HostService.Check()

			s.syncMap.Store("isRunning", false)
			s.syncMap.Store("lastCompletedTime", time.Now().Unix())
		},
	)

	_cronUtils.AddTask(
		"download",
		fmt.Sprintf("@every %ds", consts.AgentCheckDownloadInterval),
		func() {
			s.TaskService.CheckTask()
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
