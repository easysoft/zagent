package hostStatusService

import (
	"errors"
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	kvmService "github.com/easysoft/zv/internal/host/service/kvm"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"strings"
)

type StatusService struct {
	LibvirtService *kvmService.LibvirtService `inject:""`
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Check(req v1.CheckReq) (ret v1.CheckResp, err error) {
	services := strings.Split(req.Services, ",")

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceKvm.ToString(), services) {

		s.CheckKvm(&ret)

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
