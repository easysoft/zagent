package service

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type ValidService struct {
}

func NewValidService() *ValidService {
	return &ValidService{}
}

func (s *ValidService) Valid(model commDomain.ValidRequest) (result commDomain.ValidResp) {

	return
}
