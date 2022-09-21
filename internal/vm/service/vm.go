package vmAgentService

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/easysoft/zv/internal/agent/conf"
	agentService "github.com/easysoft/zv/internal/agent/service"
	testingService "github.com/easysoft/zv/internal/agent/service/testing"
	agentZentaoService "github.com/easysoft/zv/internal/agent/service/zentao"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	_domain "github.com/easysoft/zv/pkg/domain"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_dateUtils "github.com/easysoft/zv/pkg/lib/date"
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

	if consts.AuthSecret == "" || consts.ExpiredDate.Unix() < time.Now().Unix() { // re-apply token using secret
		var err error
		vm.Secret, vm.Ip, err = s.notifyHost()
		consts.AuthSecret = vm.Secret

		if err != nil {
			_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_notify", err.Error()))
			return
		}
	}

	respBytes, ok := s.register(vm)

	if ok {
		respObj := domain.RegisterResp{}
		err := json.Unmarshal(respBytes, &respObj)
		if err == nil && respObj.Token != "" {
			respObj.ExpiredDate, _ = _dateUtils.UnitToDate(respObj.ExpiredTimeUnix)
			consts.AuthToken = respObj.Token
			consts.ExpiredDate = respObj.ExpiredDate
		}
	}

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
		uri := "client/vm/register"
		url = _httpUtils.GenUrl(agentConf.Inst.Server, uri)
	} else {
		uri := "api.php/v1/executionNode/heartbeat"
		url = s.ZentaoService.GenUrl(agentConf.Inst.Server, uri)
	}

	resp, err := _httpUtils.Post(url, host)
	ok = err == nil

	return
}

func (s *VmService) notifyHost() (secret, ip string, err error) {
	uri := "virtual/notifyHost"
	url := _httpUtils.GenUrl(
		fmt.Sprintf("http://%s:%d/", consts.KvmHostIpInNatNetwork, consts.AgentPort),
		uri)

	_, macObj := _commonUtils.GetIp()
	reqData := domain.VmNotifyReq{
		MacAddress: macObj.String(),
	}

	bytes, err := _httpUtils.Post(url, reqData)
	if err != nil {
		return
	}

	resp := _domain.Response{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return
	}

	respDataBytes, _ := json.Marshal(resp.Data)
	respData := domain.VmNotifyResp{}
	err = json.Unmarshal(respDataBytes, &respData)
	if err != nil {
		return
	}

	secret = respData.Secret
	ip = respData.Ip
	if secret == "" {
		err = errors.New("secret is empty")
		return
	}

	return
}
