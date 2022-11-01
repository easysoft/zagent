package hostAgentService

import (
	"strings"

	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	kvmService "github.com/easysoft/zv/internal/host/service/kvm"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_checkUtils "github.com/easysoft/zv/internal/pkg/utils/check"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
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

		ret.Kvm, err = _checkUtils.CheckKvm()

	}

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceNovnc.ToString(), services) {

		ret.Novnc, err = _checkUtils.CheckNovnc()

	}

	if _stringUtils.StrInArr(consts.ServiceAll.ToString(), services) ||
		_stringUtils.StrInArr(consts.ServiceWebsockify.ToString(), services) {

		ret.Websockify, err = _checkUtils.CheckWebsockify()

	}

	return
}
