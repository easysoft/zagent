package commDomain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type TestProcessor struct {
	Name string               `json:"name" yaml:"name"`
	Type _const.ProcessorType `json:"type" yaml:"type"`

	Times int `json:"times" yaml:"times"` // for loop type

	Variable string `json:"variable" yaml:"variable"`   // for set variable, switch
	Expr     string `json:"condition" yaml:"condition"` // for variable
	File     string `json:"file" yaml:"file"`           // for variable

	// for loop
	InputVariablePrefix string `json:"inputVariable" yaml:"inputVariable"`
	OutputVariable      string `json:"outVariable" yaml:"outVariable"`
	StartIndex          int    `json:"startIndex" yaml:"startIndex"`
	EndIndex            int    `json:"endIndex" yaml:"endIndex"`

	Condition  string `json:"condition" yaml:"condition"` // for loop, if
	SwitchCase string `json:"variable" yaml:"variable"`   // for switch

	CookieName string `json:"cookieName" yaml:"cookieName"` // for cookieRetrieve

	ParentId uint            `json:"parentId" yaml:"parentId"`
	Children []TestProcessor `json:"children" yaml:"children" gorm:"-"`                        // contains processors
	Cases    []TestCase      `json:"cases" yaml:"cases" gorm:"many2many:r_processor_to_case;"` // refer to test cases

	DataStore map[string]interface{} `json:"dataStore" yaml:"dataStore" gorm:"-"` // store test data that loaded from file or zendata
}
