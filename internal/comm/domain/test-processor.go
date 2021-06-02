package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type TestProcessor struct {
	Name string               `json:"name" yaml:"name"`
	Type _const.ProcessorType `json:"type" yaml:"type"`

	Times int `json:"times,omitempty" yaml:"times,omitempty"` // for loop type

	Variable string `json:"variable,omitempty" yaml:"variable,omitempty"`   // for set variable, switch
	Expr     string `json:"condition,omitempty" yaml:"condition,omitempty"` // for variable
	File     string `json:"file,omitempty" yaml:"file,omitempty"`           // for variable

	// for loop
	InputVariablePrefix string `json:"inputVariable,omitempty" yaml:"inputVariable,omitempty"`
	OutputVariable      string `json:"outVariable,omitempty" yaml:"outVariable,omitempty"`
	StartIndex          int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex            int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`

	Condition  string `json:"condition,omitempty" yaml:"condition,omitempty"` // for loop, if
	SwitchCase string `json:"variable,omitempty" yaml:"variable,omitempty"`   // for switch

	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"` // for cookieRetrieve

	ParentId   uint            `json:"parentId" yaml:"parentId"`
	Children   []TestProcessor `json:"children" yaml:"children" gorm:"-"`                                       // contains processors
	Interfaces []TestInterface `json:"interfaces" yaml:"interfaces" gorm:"many2many:r_processor_to_interface;"` // refer to test cases

	DataStore map[string]interface{} `json:"dataStore" yaml:"dataStore" gorm:"-"` // store test data that loaded from file or zendata
}
