package domain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type Auth struct {
	AuthType     commConst.AuthType `json:"type" yaml:"type"`
	Entities     []Entity           `json:"entities" yaml:"entities"`
	OAuth2Entity OAuth2Entity       `json:"oAuth2Entity" yaml:"oAuth2Entity"`
}

type OAuth2Entity struct {
	TokenName   string                        `json:"tokenName" yaml:"tokenName"`
	GrantType   commConst.OAuth2TypeGrantType `json:"grantType" yaml:"grantType"`
	CallbackUrl string                        `json:"callbackUrl" yaml:"callbackUrl"`

	AuthURL string `json:"authURL" yaml:"authURL"` // url to auth

	AccessTokenURL string `json:"accessTokenURL" yaml:"accessTokenURL"` // url to get the result token

	// param will be used with auth url
	ClientId     string `json:"clientId" yaml:"clientId"`
	ClientSecret string `json:"clientSecret" yaml:"clientSecret"`
	Scope        string `json:"scope" yaml:"scope"`
	State        string `json:"state" yaml:"state"`

	// for grant type ClientCredentials
	UserName string `json:"userName" yaml:"userName"`
	Password string `json:"password" yaml:"password"`

	// for grant type AuthCodeWithPKCE
	CodeChallengeMethod commConst.CodeChallengeMethod `json:"codeChallengeMethod" yaml:"codeChallengeMethod"` // SHA-256 or Plain
	CodeVerifier        string                        `json:"codeVerifier" yaml:"codeVerifier"`               // random string

	ClientAuthentication commConst.OAuth2ClientAuthType
	AuthorizeUseBrowser  bool `json:"authorizeUseBrowser" yaml:"authorizeUseBrowser"` // open browser to auth or not

	HeaderPrefix string `json:"headerPrefix" yaml:"headerPrefix" gorm:"default:Bearer` // prefix added before access token in header
	AccessToken  string `json:"accessToken" yaml:"accessToken"`                        // auth result used to access resource
}
