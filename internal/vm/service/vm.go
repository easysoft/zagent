package vmAgentService

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	agentService "github.com/easysoft/zagent/internal/pkg/service"
	agentTestingService "github.com/easysoft/zagent/internal/pkg/service/testing"
	requestUtils "github.com/easysoft/zagent/internal/pkg/utils/request"
	_domain "github.com/easysoft/zagent/pkg/domain"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_dateUtils "github.com/easysoft/zagent/pkg/lib/date"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
)

type VmService struct {
	VmMapVar  map[string]domain.Vm
	TimeStamp int64

	VmService   *VmService                      `inject:""`
	JobService  *agentService.JobService        `inject:""`
	TestService *agentTestingService.RunService `inject:""`
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
	vm := v1.VmRegisterReq{
		MacAddress: agentConf.Inst.MacAddress,
		Ip:         agentConf.Inst.NodeIp,
	}

	if isBusy {
		vm.Status = consts.VmBusy
	} else {
		vm.Status = consts.VmReady
	}

	if consts.AuthSecret == "" || consts.ExpiredDate.Unix() < time.Now().Unix() { // re-apply token using secret
		var err error
		vm.Secret, vm.Ip, vm.AgentPortOnHost, err = s.notifyHost()
		consts.AuthSecret = vm.Secret

		if err != nil || vm.Secret == "" || vm.Ip == "" || vm.AgentPortOnHost == 0 {
			_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_notify", "error or return empty value"))
			return
		}
	}

	respBytes, ok := s.register(vm)

	if ok {
		respObj := v1.RegisterResp{}
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
	url := requestUtils.GenUrl(agentConf.Inst.Server, "api.php/v1/zanode/heartbeat")

	resp, err := _httpUtils.Post(url, host)
	ok = err == nil

	return
}

func (s *VmService) notifyHost() (secret, ip string, agentPortOnHost int, err error) {
	uri := "virtual/notifyHost"
	url := _httpUtils.GenUrl(
		fmt.Sprintf("http://%s:%d/", consts.KvmHostIpInNatNetwork, consts.AgentHostServicePort),
		uri)

	_, macObj := _commonUtils.GetIp()
	reqData := v1.VmNotifyReq{
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
	respData := v1.VmNotifyResp{}
	err = json.Unmarshal(respDataBytes, &respData)
	if err != nil {
		return
	}

	agentPortOnHost = respData.AgentPortOnHost
	secret = respData.Secret
	ip = respData.Ip
	if secret == "" {
		err = errors.New("secret is empty")
		return
	}

	return
}
