package service

import (
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type EnvironmentService struct {
	BackingRepo     *repo.BackingRepo     `inject:""`
	BrowserRepo     *repo.BrowserRepo     `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{}
}

func (s *EnvironmentService) GetMap(env model.Environment) (data map[string]interface{}) {
	categories := make([]string, 0)
	types := make([]string, 0)
	langs := make([]string, 0)
	browsers := make([]string, 0)

	backings := s.BackingRepo.ListAll()

	for _, backingItem := range backings {
		categories = append(categories, backingItem.OsCategory.ToString())

		if backingItem.OsCategory == env.OsCategory {
			types = append(types, backingItem.OsType.ToString())

			if backingItem.OsType == env.OsType {
				langs = append(langs, backingItem.OsLang.ToString())

				browserPos := s.BrowserRepo.ListByBacking(backingItem.ID)
				for _, browserItem := range browserPos {
					browsers = append(browsers, browserItem.Type.ToString()+" "+browserItem.Version)
				}
			}
		}
	}

	data["categories"] = categories
	data["types"] = types
	data["langs"] = langs
	data["browsers"] = browsers

	return
}

type Nested struct {
	Level int
	Code  string
	Child []Nested
}
