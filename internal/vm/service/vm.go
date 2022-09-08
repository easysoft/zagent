package vmAgentService

import (
	"github.com/easysoft/zv/internal/agent/conf"
	agentService "github.com/easysoft/zv/internal/agent/service"
	testingService "github.com/easysoft/zv/internal/agent/service/testing"
	agentZentaoService "github.com/easysoft/zv/internal/agent/service/zentao"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_i118Utils "github.com/easysoft/zv/pkg/lib/i118"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"strings"
	"time"
)

type VmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	VmService   *VmService                 `inject:""`
	JobService  *agentService.JobService   `inject:""`
	TestService *testingService.RunService `inject:""`

	ZentaoService *agentZentaoService.ZentaoService `inject:""`
}

func NewVmService() *VmService {
	s := VmService{}
	s.TimeStamp = time.Now().Unix()
	s.VmMapVar = map[string]domain.Vm{}

	return &s
}

func (s *VmService) Check() {
	// is running，register busy
	if s.JobService.IsRunning() {
		s.VmService.Register(true)
		return
	}

	// no task to run, submit free
	if s.JobService.GetTaskSize() == 0 {
		s.VmService.Register(false)
		return
	}

	// has task to run，register busy, then run
	job := s.JobService.PeekJob()
	s.VmService.Register(true)

	s.JobService.StartTask()
	s.TestService.Run(&job)
	s.JobService.RemoveTask()
	s.JobService.EndTask()

}

func (s *VmService) Register(isBusy bool) (resp string, ok bool) {
	vm := domain.Vm{
		MacAddress: agentConf.Inst.MacAddress,
		Ip:         agentConf.Inst.NodeIp, Port: agentConf.Inst.NodePort,
		Name: agentConf.Inst.NodeName, WorkDir: agentConf.Inst.WorkDir,
	}

	if isBusy {
		vm.Status = consts.VmBusy
	} else {
		vm.Status = consts.VmReady
	}

	var url string
	if strings.Index(agentConf.Inst.Server, ":8085") > -1 {
		uri := "client/vm/register"
		url = _httpUtils.GenUrl(agentConf.Inst.Server, uri)

		bytes, err := _httpUtils.Post(url, vm)
		resp = string(bytes)
		ok = err == nil

	} else {
		uri := "api.php/v1/vm/register"
		url = s.ZentaoService.GenUrl(agentConf.Inst.Server, uri)
		s.ZentaoService.GetConfig(agentConf.Inst.Server)

		resp, ok = s.ZentaoService.Post(url, vm, true)
	}

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}

	return
}
