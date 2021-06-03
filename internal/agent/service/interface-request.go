package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type InterfaceRequestService struct {
	CommonService
}

func NewInterfaceRequestService() *InterfaceRequestService {
	return &InterfaceRequestService{}
}

func (s *InterfaceRequestService) Request(build *commDomain.Build, interf commDomain.TestInterface) {

}
