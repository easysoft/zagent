package vmAgentService

import (
	"encoding/json"
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
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

func (s *VmService) Register(isBusy bool) (ok bool) {
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

	respBytes, ok := s.register(vm)

	respObj := v1.HostRegisterResp{}
	json.Unmarshal(respBytes, &respObj)
	consts.AuthToken = respObj.Data.Token

	if consts.AuthToken == "" {
		ok = false
	}

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, respBytes))
	}

	return
}

func (s *VmService) register(host interface{}) (resp []byte, ok bool) {
	var url string
	if strings.Index(agentConf.Inst.Server, ":8085") > -1 {
		uri := "client/host/register"
		url = _httpUtils.GenUrl(agentConf.Inst.Server, uri)
	} else {
		uri := "api.php/v1/host/register"
		url = s.ZentaoService.GenUrl(agentConf.Inst.Server, uri)
	}

	resp, err := _httpUtils.Post(url, host)
	ok = err == nil

	return
}
