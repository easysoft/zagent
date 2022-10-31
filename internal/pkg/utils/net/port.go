package netUtils

import (
	"fmt"
	"strconv"
	"strings"

	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

func GetUsedPortByKeyword(keyword string, defaultVal int) (port int, err error) {
	cmd := fmt.Sprintf(`ss -tnlp | grep %s | awk '{ print $4 }'`, keyword)
	output, _ := _shellUtils.ExeSysCmd(cmd)
	output = strings.TrimSpace(output)

	if output == "" {
		return port, fmt.Errorf("%s not found", keyword)
	}

	list := strings.Split(output, "\n")
	lastInfo := list[len(list)-1]
	lastInfo = strings.TrimSpace(lastInfo)
	info := strings.Split(lastInfo, ":")
	port, err = strconv.Atoi(info[len(info)-1])

	if port == 0 {
		port = defaultVal
	}

	return
}
