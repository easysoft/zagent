package serverService

import (
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
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
	categoryMap := map[string]bool{}
	typeMap := map[string]bool{}
	langMap := map[string]bool{}
	browserMap := map[string]bool{}

	categories := make([]string, 0)
	types := make([]string, 0)
	langs := make([]string, 0)
	browsers := make([]string, 0)

	backings := s.BackingRepo.ListAll()

	for _, backingItem := range backings {
		if !categoryMap[backingItem.OsCategory.ToString()] {
			categories = append(categories, backingItem.OsCategory.ToString())
			categoryMap[backingItem.OsCategory.ToString()] = true
		}

		if backingItem.OsCategory == env.OsCategory {
			if !typeMap[backingItem.OsType.ToString()] {
				types = append(types, backingItem.OsType.ToString())
				typeMap[backingItem.OsType.ToString()] = true
			}

			if backingItem.OsType == env.OsType {
				if !langMap[backingItem.OsLang.ToString()] {
					langs = append(langs, backingItem.OsLang.ToString())
					langMap[backingItem.OsLang.ToString()] = true
				}

				browserPos := s.BrowserRepo.ListByBacking(backingItem.ID)
				for _, browserItem := range browserPos {
					if !browserMap[browserItem.Type.ToString()] {
						browsers = append(browsers, browserItem.Type.ToString())
						browserMap[browserItem.Type.ToString()] = true
					}
				}
			}
		}
	}

	data = map[string]interface{}{}
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
