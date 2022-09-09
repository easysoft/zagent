package commonUtils

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/comm/const"
	_const "github.com/easysoft/zv/pkg/const"
	"strings"
)

var ()

func GenAuthorization() (ret string) {
	return fmt.Sprintf("%s %s", _const.Bearer, consts.AuthToken)
}

func GetTokenInAuthorization(value string) (token string) {
	return strings.Replace(value, _const.Bearer+" ", "", -1)
}
