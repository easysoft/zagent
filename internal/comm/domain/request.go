package domain

import _const "github.com/easysoft/zv/internal/pkg/const"

type Request struct {
	Method _const.HttpMethod `json:"method" yaml:"method"`
	Header Header            `json:"header" yaml:"header"`
	URL    URL               `json:"url" yaml:"url"`
	Body   Body              `json:"body" yaml:"body"`
	Auth   Auth              `json:"auth" yaml:"auth"`
}

type Header struct {
	UserAgent string `json:"useAgent" yaml:"useAgent"`
	Accept    string `json:"accept" yaml:"accept"`
}

type URL struct {
	Raw      string   `json:"raw" yaml:"raw"`
	Protocol string   `json:"protocol" yaml:"protocol"`
	Host     string   `json:"host" yaml:"host"`
	Port     int      `json:"port" yaml:"port"`
	Path     string   `json:"path" yaml:"path"`
	Params   []Entity `json:"params" yaml:"params"`
}

type Body struct {
	Mode     string   `json:"mode" yaml:"mode"`
	FormData []Entity `json:"formData" yaml:"formdata"`
	Raw      string   `json:"-" yaml:"-"`
}

type Entity struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
	Type  string `json:"type,omitempty" yaml:"type"`
}
