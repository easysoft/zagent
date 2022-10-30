package hostAgentService

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	kvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
)

type StatusService struct {
	LibvirtService *kvmService.LibvirtService `inject:""`
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Check(req v1.ServiceReq) (ret v1.CheckResp, err error) {
	services := strings.Split(req.Services, ",")

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceKvm.ToString(), services) {

		s.CheckKvm(&ret)

	} else if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceNovnc.ToString(), services) {

		s.CheckNovnc(&ret)

	} else if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceWebsockify.ToString(), services) {

		s.CheckWebsockify(&ret)

	}

	return
}

func (s *StatusService) CheckKvm(ret *v1.CheckResp) (err error) {
	ret.Kvm = consts.HostServiceNotAvailable

	defer func() {
		err1 := recover()
		if err1 != nil {
			err = errors.New(fmt.Sprintf("%v", err1))
		}
	}()

	kvmActive := s.LibvirtService.IsAlive()
	if kvmActive {
		ret.Kvm = consts.HostServiceReady
		return
	}

	out, _ := _shellUtils.ExeShell("which libvirtd") // ps -ef | grep libvirt | grep -v grep | grep -v dnsmasq
	if strings.Index(out, "libvirtd") < 0 {
		ret.Kvm = consts.HostServiceNotInstall
	}

	return
}

func (s *StatusService) CheckNovnc(ret *v1.CheckResp) (err error) {
	ret.Novnc = consts.HostServiceNotAvailable

	// get :agentConf.Inst.NodePort/novnc
	timeout := time.Second

	address := net.JoinHostPort(consts.Localhost, strconv.Itoa(agentConf.Inst.NodePort))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		_logUtils.Infof("tcp connect to %s error: %s", address, err)
	}

	if conn != nil {
		ret.Novnc = consts.HostServiceReady
		defer conn.Close()

	} else {
		pth := filepath.Join(consts.NovncDir, "index.html")
		found := _fileUtils.FileExist(pth)

		if !found {
			ret.Novnc = consts.HostServiceNotInstall
		}
	}

	return
}

func (s *StatusService) CheckWebsockify(ret *v1.CheckResp) (err error) {
	ret.Websockify = consts.HostServiceNotAvailable

	timeout := time.Second

	address := net.JoinHostPort(consts.Localhost, strconv.Itoa(consts.WebsockifyPort))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		_logUtils.Infof("tcp connect to %s error: %s", address, err)
	}

	if conn != nil {
		ret.Websockify = consts.HostServiceReady
		defer conn.Close()

	} else {
		pth := filepath.Join(consts.WebsockifyDir, "run")
		found := _fileUtils.FileExist(pth)

		if !found {
			ret.Websockify = consts.HostServiceNotInstall
		}
	}

	return
}
