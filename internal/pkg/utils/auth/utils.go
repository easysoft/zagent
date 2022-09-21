package authUtils

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_const "github.com/easysoft/zv/pkg/const"
	"net/http"
	"strings"
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
