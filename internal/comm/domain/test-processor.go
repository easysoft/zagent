package domain

import (
	"github.com/easysoft/zagent/internal/comm/const"
)

type TestProcessor struct {
	DataLoopProcessor
	ExtractorProcessor

	Name     string               `json:"name" yaml:"name"`
	Comments string               `json:"comments" yaml:"comments"`
	Type     consts.ProcessorType `json:"type" yaml:"type"`

	ParentId uint `json:"parentId" yaml:"parentId"`

	// can be interface, Processor.
	Children []interface{} `json:"children" yaml:"children" gorm:"-"`

	// results
	Results []string `json:"results" yaml:"results" gorm:"-"`

	// store test data loaded by DataLoopProcessor
	DataStore map[string]interface{} `json:"dataStore" yaml:"dataStore" gorm:"-"`
}

// automated cookie management

type DataLoopProcessor struct {
	Src  consts.DataSource `json:"src,omitempty" yaml:"src,omitempty"`
	Path string            `json:"path,omitempty" yaml:"path,omitempty"`

	Loop           int    `json:"loop,omitempty" yaml:"loop,omitempty"`
	StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`
	IsRand         bool   `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce         bool   `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`
	VarNamePostfix string `json:"varNamePostfix,omitempty" yaml:"varNamePostfix,omitempty"`
}

type ExtractorProcessor struct {
	VarName string `json:"varName,omitempty" yaml:"varName,omitempty"`
	Default string `json:"default,omitempty" yaml:"default,omitempty"`

	Source consts.ExtractorSource `json:"source,omitempty" yaml:"source,omitempty"`
	Type   consts.ExtractorType   `json:"type,omitempty" yaml:"type,omitempty"`
	Expr   string                 `json:"expr,omitempty" yaml:"expr,omitempty"`
}
