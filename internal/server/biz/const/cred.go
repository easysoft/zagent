package bizConst

import (
	"github.com/pkg/errors"
	"time"
)

const (
	ZXW_SESSION_TOKEN_PREFIX          = "ZST:"
	ZXW_SESSION_BIND_USER_PREFIX      = "ZSBU:"
	ZXW_SESSION_USER_PREFIX           = "ZSU:"
	ZXW_SESSION_USER_MAX_TOKEN_PREFIX = "ZXWUserMaxToken"
)

var (
	ERR_TOKEN_INVALID                  = errors.New("token is invalid!")
	ZXW_SESSION_USER_MAX_TOKEN_DEFAULT = 10
)

const (
	NoneScope uint64 = iota
	AdminScope
)

const (
	NonoAuth int = iota
	AuthPwd
	AuthCode
	AuthThirdparty
)

const (
	LoginTypeWeb int = iota
	LoginTypeApp
	LoginTypeWx
	LoginTypeAlipay
	LoginApplet
)

var (
	RedisSessionTimeoutWeb    = 30 * time.Minute
	RedisSessionTimeoutApp    = 24 * time.Hour
	RedisSessionTimeoutApplet = 7 * 24 * time.Hour
	RedisSessionTimeoutWx     = 5 * 52 * 168 * time.Hour
)

type UserCredentials struct {
	UserId       string `json:"user_id" redis:"user_id"`
	LoginType    int    `json:"login_type" redis:"login_type"`
	AuthType     int    `json:"auth_type" redis:"auth_type"`
	CreationDate int64  `json:"creation_data" redis:"creation_data"`
	ExpiresIn    int    `json:"expires_in" redis:"expires_in"`
	Scope        uint64 `json:"scope" redis:"scope"`
	Token        string `json:"token" redis:"token"`
}
