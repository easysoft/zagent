package requestUtils

import (
	"fmt"
	"strings"
)

const (
	ApiPath = "api.php/v1/"
)

var (
	zenTaoVersion = ""
	sessionVar    = ""
	sessionId     = ""
	requestFix    = ""
)

func GenUrl(server string, path string) string {
	server = UpdateUrl(server)
	url := fmt.Sprintf("%s%s", server, path)
	return url
}
func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}
