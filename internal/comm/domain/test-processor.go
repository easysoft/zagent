package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type TestProcessor struct {
	Name     string               `json:"name" yaml:"name"`
	Comments string               `json:"comments" yaml:"comments"`
	Type     _const.ProcessorType `json:"type" yaml:"type"`

	Variable string `json:"variable,omitempty" yaml:"variable,omitempty"`   // for set variable, switch
	Expr     string `json:"condition,omitempty" yaml:"condition,omitempty"` // for variable
	File     string `json:"file,omitempty" yaml:"file,omitempty"`           // for variable

	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"` // for cookieRetrieve

	ParentId   uint            `json:"parentId" yaml:"parentId"`
	Children   []TestProcessor `json:"children" yaml:"children" gorm:"-"`                                       // contains processors
	Interfaces []TestInterface `json:"interfaces" yaml:"interfaces" gorm:"many2many:r_processor_to_interface;"` // refer to test cases

	DataStore map[string]interface{} `json:"dataStore" yaml:"dataStore" gorm:"-"` // store test data that loaded from file or zendata
}

// logic
type LoopProcessor struct {
	Count    int  `json:"count,omitempty" yaml:"count,omitempty"`
	Infinite bool `json:"infinite,omitempty" yaml:"infinite,omitempty"`
}

type EachProcessor struct {
	InputVariablePrefix string `json:"inputVariable,omitempty" yaml:"inputVariable,omitempty"`
	OutputVariableName  string `json:"outVariableName,omitempty" yaml:"outVariableName,omitempty"`
	StartIndex          int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"` // inclusive
	EndIndex            int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`     // exclusive
}

type WhileProcessor struct {
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`
}

type IfProcessor struct {
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`
}

type SwitchProcessor struct {
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
