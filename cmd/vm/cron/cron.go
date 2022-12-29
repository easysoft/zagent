package vmCron

import (
	"fmt"
	"sync"
	"time"

	v1 "github.com/easysoft/zagent/cmd/vm/router/v1"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	vmAgentService "github.com/easysoft/zagent/internal/vm/service"
	_cronUtils "github.com/easysoft/zagent/pkg/lib/cron"
	_dateUtils "github.com/easysoft/zagent/pkg/lib/date"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type CronService struct {
	syncMap       sync.Map
	VmService     *vmAgentService.VmService     `inject:""`
	StatusService *vmAgentService.StatusService `inject:""`
	ToolService   *vmAgentService.ToolService   `inject:""`
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

	_cronUtils.AddTask(
		"runZTF",
		fmt.Sprintf("@every %ds", consts.AgentCheckExecutionInterval),
		func() {
			s.runZtf()
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})

	time.AfterFunc(2*time.Second, func() { s.checkTask() })
	time.AfterFunc(2*time.Second, func() { s.runZtf() })
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

func (s *CronService) runZtf() {
	serviceStatus, _ := s.StatusService.Check(v1.VmServiceCheckReq{Services: "ztf"})
	fmt.Println(111, serviceStatus.ZtfStatus)
	if serviceStatus.ZtfStatus == consts.HostServiceReady {
		return
	} else if serviceStatus.ZtfStatus == consts.HostServiceNotAvailable {
		s.ToolService.StartToolByName("ztf")
	} else if serviceStatus.ZtfStatus == consts.HostServiceNotInstall {
		s.ToolService.Setup(v1.VmServiceInstallReq{Name: "ztf"})
	}
}
