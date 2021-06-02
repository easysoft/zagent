package commDomain

import "time"

type Cookie struct {
	Name     string    `json:"name" yaml:"name"`
	Value    string    `json:"value" yaml:"value"`
	Domain   string    `json:"domain" yaml:"domain"`
	Path     string    `json:"path" yaml:"path"`
	Expires  time.Time `json:"expires" yaml:"expires"`
	HttpOnly bool      `json:"httpOnly" yaml:"httpOnly"`
	Secure   bool      `json:"secure" yaml:"secure"`
	MaxAge   int       `json:"maxAge" yaml:"maxAge"`
	Raw      string    `json:"raw" yaml:"raw"`
}
