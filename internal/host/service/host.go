package hostAgentService

import (
	"encoding/json"
	"strings"
	"time"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	hostKvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	agentService "github.com/easysoft/zagent/internal/pkg/service"
	agentTestingService "github.com/easysoft/zagent/internal/pkg/service/testing"
	requestUtils "github.com/easysoft/zagent/internal/pkg/utils/request"
	_dateUtils "github.com/easysoft/zagent/pkg/lib/date"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
)

type HostService struct {
	VmService *hostKvmService.KvmService `inject:""`

	JobService  *agentService.JobService        `inject:""`
	TestService *agentTestingService.RunService `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Check() {
	s.Register(false)
	return

	//// is running，register busy
	//if s.JobService.IsRunning() {
	//	s.Register(true)
	//	return
	//}
	//
	//// no task to run, submit free
	//if s.JobService.GetTaskSize() == 0 {
	//	s.Register(false)
	//	return
	//}
	//
	//// has task to run，register busy, then run
	//job := s.JobService.PeekJob()
	//s.Register(true)
	//
	//s.JobService.StartCreateSnapTask()
	//
	//s.PassEnvsToContainerIfNeeded(&job)
	//s.TestService.Run(&job)
	//
	//s.JobService.RemoveTask()
	//s.JobService.EndTask()
}

func (s *HostService) Register(isBusy bool) {
	host := domain.HostNode{
		Node: domain.Node{
			Ip:   agentConf.Inst.NodeIp,
			Port: agentConf.Inst.NodePort,
		},
	}

	if consts.AuthToken == "" || consts.ExpiredDate.Unix() < time.Now().Unix() { // re-apply token using secret
		host.Secret = agentConf.Inst.Secret
	}

	if isBusy {
		host.Status = consts.HostBusy
	} else {
		host.Status = consts.HostOnline
	}

	host.Vms = s.VmService.GetVms()

	s.VmService.UpdateVmMapAndDestroyTimeout(host.Vms)
	respBytes, ok := s.register(host)
	hostBytes, _ := json.Marshal(host)
	_logUtils.Info(string(hostBytes))

	if ok {
		respObj := v1.RegisterResp{}
		err := json.Unmarshal(respBytes, &respObj)
		if err == nil && respObj.Token != "" {
			respObj.TokenTime, _ = _dateUtils.UnitToDate(respObj.TokenTimeUnix)
			consts.AuthToken = respObj.Token
			consts.ExpiredDate = respObj.TokenTime
		}
	}

	if consts.AuthToken == "" {
		ok = false
	}

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, string(respBytes)))
	}
}

func (s *HostService) register(host interface{}) (resp []byte, ok bool) {
	url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/host/heartbeat")

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
