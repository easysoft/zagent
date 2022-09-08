package agentZentaoService

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

type ZentaoResponse struct {
	Status string
	Data   string
}

type ZentaoService struct {
}

func NewZentaoService() *ZentaoService {
	return &ZentaoService{}
}

func (s *ZentaoService) GenUrl(server string, path string) string {
	server = s.UpdateUrl(server)
	url := fmt.Sprintf("%s%s", server, path)
	return url
}
func (s *ZentaoService) UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}
