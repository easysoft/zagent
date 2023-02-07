package authUtils

import (
	"fmt"
	"net/http"
	"strings"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	_const "github.com/easysoft/zagent/pkg/const"
)

func AddBearTokenIfNeeded(req *http.Request) {
	if strings.Index(req.URL.Path, "api.php") > -1 && consts.AuthToken != "" {
		req.Header.Set(_const.Authorization, GenAuthorization())
	}
}

func GenAuthorization() (ret string) {
	return fmt.Sprintf("%s %s", _const.Bearer, consts.AuthToken)
}

func GetTokenInAuthorization(value string) (token string) {
	return strings.Replace(value, _const.Bearer+" ", "", -1)
}
