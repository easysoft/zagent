package vmAgentService

import (
	v1 "github.com/easysoft/zagent/cmd/vm/router/v1"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	natHelper "github.com/easysoft/zagent/internal/pkg/utils/net"
	"net"
	"strconv"
	"strings"
	"time"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
)

type StatusService struct {
	ToolService *ToolService `inject:""`
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Check(req v1.VmServiceCheckReq) (ret v1.VmServiceCheckResp, err error) {
	services := strings.Split(req.Services, ",")

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceZtf.ToString(), services) {

		ret.ZtfStatus, ret.ZtfVersion, err = s.CheckZtf()
	}

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceZd.ToString(), services) {

		ret.ZdStatus, ret.ZdVersion, err = s.CheckZd()
	}

	return
}

func (s *StatusService) CheckZtf() (status consts.HostServiceStatus, version string, err error) {
	status = consts.HostServiceNotAvailable

	timeout := time.Second

	port, _ := natHelper.GetUsedPortByKeyword("ztf", agentConf.Inst.ZtfPort)
	address := net.JoinHostPort(consts.Localhost, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, timeout)

	version, num := s.ToolService.getOldVersion("ztf")

	if conn != nil {
		defer conn.Close()

		status = consts.HostServiceReady
	} else if num == 0 {
		status = consts.HostServiceNotInstall
	}

	return
}

func (s *StatusService) CheckZd() (status consts.HostServiceStatus, version string, err error) {
	status = consts.HostServiceNotAvailable

	timeout := time.Second

	port, _ := natHelper.GetUsedPortByKeyword("zd", agentConf.Inst.ZdPort)
	address := net.JoinHostPort(consts.Localhost, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, timeout)

	version, num := s.ToolService.getOldVersion("zd")

	if conn != nil {
		defer conn.Close()

		status = consts.HostServiceReady
	} else if num == 0 {
		status = consts.HostServiceNotInstall
	}

	return
}
