package netUtils

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
)

const (
	httpConf = `
                server{
					listen      %d;
					location / {
						proxy_pass   http://%s:%d;
					}
				}`

	streamConf = `
				server {
					listen %d;   
					proxy_connect_timeout 1h;
					proxy_timeout 1h;
					proxy_pass %s:%d;
				}`
)

func GetValidPort() (hostPort int, err error) {
	cmd := fmt.Sprintf(`netstat -tln | awk '{print $4}' | grep -o ':51[0-9]\{3\}'`)
	output, _ := _shellUtils.ExeSysCmd(cmd)

	list := strings.Split(output, "\n")

	for p := consts.NatPortStart; p <= consts.NatPortEnd; p++ {
		str := ":" + strconv.Itoa(p)
		if !_stringUtils.StrInArr(str, list) {
			hostPort = p
			break
		}
	}

	if hostPort == 0 {
		err = errors.New("no port left")
	}

	return
}

func ForwardPortIfNeeded(vmIp string, vmPort int, typ consts.NatForwardType) (hostPort int, alreadyMapped bool, err error) {
	hostPort, err = GetValidPort()
	if err != nil {
		return
	}

	mappedHostPort := getMappedInfo(vmIp, vmPort, typ)
	if mappedHostPort != 0 {
		hostPort = mappedHostPort
		alreadyMapped = true
		return
	}

	_, pth, err := getNginxHotLoadingConf(vmIp, vmPort, hostPort, typ)
	if err != nil {
		return
	}

	content := "N/A"
	if typ == consts.Http {
		content = fmt.Sprintf(httpConf, hostPort, vmIp, vmPort)
	} else {
		content = fmt.Sprintf(streamConf, hostPort, vmIp, vmPort)
	}

	_fileUtils.WriteFile(pth, content)

	reloadNginx()

	return
}

func getMappedInfo(vmIp string, vmPort int, typ consts.NatForwardType) (mappedPort int) {
	homeDir, _ := _fileUtils.GetUserHome()
	dir := filepath.Join(homeDir, "zagent", "nginx", fmt.Sprintf("conf.%s.d", typ))

	cmd := fmt.Sprintf("ls -al %s | grep %s:%d | grep -v grep", dir, vmIp, vmPort)
	out, _ := _shellUtils.ExeSysCmd(cmd)

	regx, _ := regexp.Compile(fmt.Sprintf(`%s:%d@([0-9]+).conf`, vmIp, vmPort))
	arr := regx.FindStringSubmatch(out)
	if len(arr) > 1 {
		portStr := arr[1]
		mappedPort, _ = strconv.Atoi(portStr)
	}

	return
}

func RemoveForward(vmIp string, vmPort int, typ consts.NatForwardType) (err error) {
	homeDir, _ := _fileUtils.GetUserHome()
	dir := filepath.Join(homeDir, "zagent", "nginx", fmt.Sprintf("conf.%s.d", typ))

	name := fmt.Sprintf("%s:*@*.conf", vmIp)
	if vmPort > 0 {
		name = fmt.Sprintf("%s:%d@*.conf", vmIp, vmPort)
	}

	pth := filepath.Join(dir, name)

	cmd := fmt.Sprintf("rm -rf %s", pth)
	_shellUtils.ExeSysCmd(cmd)

	reloadNginx()

	return
}
func RemoveForwardByPort(vmPort int, typ consts.NatForwardType) (err error) {
	homeDir, _ := _fileUtils.GetUserHome()
	dir := filepath.Join(homeDir, "zagent", "nginx")
	pth := filepath.Join(dir, fmt.Sprintf("conf.%s.d/*:%d@*.conf", typ, vmPort))

	cmd := fmt.Sprintf("rm -rf %s", pth)
	_, err = _shellUtils.ExeSysCmd(cmd)

	return
}

//func getNginxConf() (ret string, err error) {
//	cmd := fmt.Sprintf(`nginx -t 2>&1 | grep test`)
//	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
//
//	regx, _ := regexp.Compile(`file (.+) test`)
//	arr := regx.FindStringSubmatch(string(out))
//
//	if len(arr) > 1 {
//		ret = arr[1]
//	} else {
//		err = errors.New("not found")
//	}
//
//	return
//}

func getNginxHotLoadingConf(vmIp string, vmPort int, hostPort int, typ consts.NatForwardType) (
	name, ret string, err error) {

	homeDir, _ := _fileUtils.GetUserHome()
	dir := filepath.Join(homeDir, "zagent", "nginx")
	name = fmt.Sprintf("%s:%d@%d", vmIp, vmPort, hostPort)

	ret = filepath.Join(dir, fmt.Sprintf("conf.%s.d", typ), name+".conf")

	return
}

func reloadNginx() {
	cmd := fmt.Sprintf(`nginx -s reload`)
	_shellUtils.ExeSysCmd(cmd)
}
