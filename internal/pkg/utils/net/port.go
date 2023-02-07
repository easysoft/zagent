package netUtils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

// find port by keyword
func GetUsedPortByKeyword(keyword string, defaultVal int) (port int, err error) {
	if _commonUtils.IsWin() {
		return GetWindowsUsedPortByKeyword(keyword, defaultVal)
	}
	cmd := fmt.Sprintf(`ss -tnlp | grep %s | awk '{ print $4 }'`, keyword)
	output, _ := _shellUtils.ExeSysCmd(cmd)
	output = strings.TrimSpace(output)

	if output != "" {
		list := strings.Split(output, "\n")
		lastInfo := list[len(list)-1]
		lastInfo = strings.TrimSpace(lastInfo)
		info := strings.Split(lastInfo, ":")
		port, err = strconv.Atoi(info[len(info)-1])
	}

	if port == 0 {
		port, err = GetUsedPortByPs(keyword, defaultVal)
	}

	if port == 0 {
		port = defaultVal
	}

	return
}

func GetWindowsUsedPortByKeyword(keyword string, defaultVal int) (port int, err error) {
	cmd := fmt.Sprintf(`tasklist|findstr %s`, keyword)
	output, _ := _shellUtils.ExeSysCmd(cmd)

	output = strings.TrimSpace(output)
	reg := regexp.MustCompile(`( +)`)
	output = reg.ReplaceAllString(output, " ")
	cols := strings.Split(output, " ")
	if len(cols) < 2 {
		return
	}

	pid := cols[1]
	pidNum, _ := strconv.Atoi(pid)
	if pidNum == 0 {
		return
	}

	cmd = fmt.Sprintf(`netstat -aon|findstr %s`, pid)
	output, _ = _shellUtils.ExeSysCmd(cmd)

	output = strings.TrimSpace(output)
	output = reg.ReplaceAllString(output, " ")
	cols = strings.Split(output, " ")
	if len(cols) < 2 {
		return
	}

	address := strings.Split(cols[1], ":")
	port, err = strconv.Atoi(address[1])
	return
}

// find port by ps command
func GetUsedPortByPs(keyword string, defaultVal int) (port int, err error) {
	cmd := fmt.Sprintf(`ps aux | grep "%s" |grep -v grep | tr -s ' ' | cut -d' ' -f2,11-20`, keyword)
	output, _ := _shellUtils.ExeSysCmd(cmd)
	output = strings.TrimSpace(output)

	if output == "" {
		return port, fmt.Errorf("%s not found", keyword)
	}

	list := strings.Split(output, "\n")
	for _, l := range list {
		l = strings.TrimSpace(l)
		if !strings.Contains(l, keyword) {
			continue
		}
		info := strings.Split(l, " ")

		cmd := fmt.Sprintf(`ss -tnlp | grep pid=%s, | awk '{ print $4 }'`, strings.TrimSpace(info[0]))
		output, _ := _shellUtils.ExeSysCmd(cmd)
		output = strings.TrimSpace(output)

		if output == "" {
			continue
		}

		list2 := strings.Split(output, "\n")
		for _, l2 := range list2 {
			l2 = strings.TrimSpace(l2)
			info = strings.Split(l2, ":")
			port2, _ := strconv.Atoi(info[len(info)-1])

			if port == 0 || port > port2 {
				port = port2
			}

		}

	}

	if port == 0 {
		port = defaultVal
	}

	return
}
