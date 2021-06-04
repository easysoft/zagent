package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
)

type InterfaceRequestService struct {
	CommonService
}

func NewInterfaceRequestService() *InterfaceRequestService {
	return &InterfaceRequestService{}
}

func (s *InterfaceRequestService) Request(build *commDomain.Build, interf commDomain.TestInterface) {
	i := 0

	_logUtils.Infof("exec interface %d", i)
}
