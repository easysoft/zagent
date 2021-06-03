package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
)

type TestProcessor struct {
	Name     string               `json:"name" yaml:"name"`
	Comments string               `json:"comments" yaml:"comments"`
	Type     _const.ProcessorType `json:"type" yaml:"type"`

	ParentId   uint            `json:"parentId" yaml:"parentId"`
	Children   []TestProcessor `json:"children" yaml:"children" gorm:"-"`                                       // contains processors
	Interfaces []TestInterface `json:"interfaces" yaml:"interfaces" gorm:"many2many:r_processor_to_interface;"` // refer to test cases

	DataStore map[string]interface{} `json:"dataStore" yaml:"dataStore" gorm:"-"` // store test data that loaded from file or zendata
}

// automated cookie management

type DataLoopProcessor struct {
	Source commConst.DataSource `json:"source,omitempty" yaml:"source,omitempty"`
	Path   string               `json:"path,omitempty" yaml:"path,omitempty"`

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

	Source commConst.ExtractorSource `json:"source,omitempty" yaml:"source,omitempty"`
	Type   commConst.ExtractorType   `json:"type,omitempty" yaml:"type,omitempty"`
	Expr   string                    `json:"expr,omitempty" yaml:"expr,omitempty"`
}
