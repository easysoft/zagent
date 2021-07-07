package agentCron

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_cronUtils "github.com/easysoft/zagent/internal/pkg/lib/cron"
)

type CronService struct {
	CheckService *agentService.CheckService `inject:""`
}

func NewAgentCron() *CronService {
	inst := &CronService{}
	inst.Init()
	return inst
}

func (s *CronService) Init() {
	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.AgentCheckInterval),
		func() {
			s.CheckService.Check()
		},
	)
}
