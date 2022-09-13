package hostAgentService

import (
	"encoding/json"
	v1 "github.com/easysoft/zv/cmd/server/router/v1"
	agentConf "github.com/easysoft/zv/internal/agent/conf"
	agentService "github.com/easysoft/zv/internal/agent/service"
	testingService "github.com/easysoft/zv/internal/agent/service/testing"
	agentZentaoService "github.com/easysoft/zv/internal/agent/service/zentao"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	hostKvmService "github.com/easysoft/zv/internal/host/service/kvm"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_i118Utils "github.com/easysoft/zv/pkg/lib/i118"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"strings"
)

type HostService struct {
	VmService *hostKvmService.VmService `inject:""`

	JobService  *agentService.JobService   `inject:""`
	TestService *testingService.RunService `inject:""`

	ZentaoService *agentZentaoService.ZentaoService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Check() {
	// is running，register busy
	if s.JobService.IsRunning() {
		s.Register(true)
		return
	}

	// no task to run, submit free
	if s.JobService.GetTaskSize() == 0 {
		s.Register(false)
		return
	}

	// has task to run，register busy, then run
	job := s.JobService.PeekJob()
	s.Register(true)

	s.JobService.StartTask()

	s.PassEnvsToContainerIfNeeded(&job)
	s.TestService.Run(&job)

	s.JobService.RemoveTask()
	s.JobService.EndTask()
}

func (s *HostService) Register(isBusy bool) {
	host := domain.HostNode{
		Node: domain.Node{
			Ip:     agentConf.Inst.NodeIp,
			Port:   agentConf.Inst.NodePort,
			Secret: agentConf.Inst.Secret,
		},
	}
	if isBusy {
		host.Status = consts.HostBusy
	} else {
		host.Status = consts.HostOnline
	}

	host.Vms = s.VmService.GetVms()
	s.VmService.UpdateVmMapAndDestroyTimeout(host.Vms)

	respBytes, ok := s.register(host)

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
}

func (s *HostService) register(host interface{}) (resp []byte, ok bool) {
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

func (s *HostService) PassEnvsToContainerIfNeeded(build *domain.Build) {
	str := "docker run"
	if strings.Index(build.BuildCommands, "docker run") > -1 {
		newStr := str
		for _, env := range strings.Split(build.EnvVars, "\n") {
			newStr += " -e " + env + " "
		}

		build.BuildCommands = strings.Replace(build.BuildCommands, str, newStr, -1)
	}
}
