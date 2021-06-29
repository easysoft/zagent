package service

import (
	"github.com/easysoft/zagent/internal/server/repo"
)

type EnvService struct {
	BackingRepo *repo.BackingRepo `inject:""`
	BrowserRepo *repo.BrowserRepo `inject:""`
}

func NewEnvService() *EnvService {
	return &EnvService{}
}

func (s *EnvService) GetMap() (mp map[string]interface{}, err error) {
	backings := s.BackingRepo.ListAll()

	for _, backingItem := range backings {
		browsers := s.BrowserRepo.ListByBacking(backingItem.ID)

		for _, browserItem := range browsers {

		}
	}

	return
}
