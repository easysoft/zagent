package commDomain

import (
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type AutomatedTest struct {
	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	AppPath     string `json:"appPath,omitempty"`
	ScriptUrl   string `json:"scriptUrl,omitempty"`
	ScmAddress  string `json:"scmAddress,omitempty"`
	ScmAccount  string `json:"scmAccount,omitempty"`
	ScmPassword string `json:"scmPassword,omitempty"`

	AppUrl          string          `json:"appUrl,omitempty"`
	BuildCommands   string          `json:"buildCommands,omitempty"`
	ResultFiles     string          `json:"resultFiles,omitempty"`
	KeepResultFiles _domain.Boolean `json:"keepResultFiles,omitempty"`
	ResultPath      string          `json:"resultPath,omitempty"`
	ResultMsg       string          `json:"resultMsg,omitempty"`
}
