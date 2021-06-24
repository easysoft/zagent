package service

import (
	"github.com/easysoft/zagent/internal/server/repo"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermService() *PermService {
	return &PermService{}
}
