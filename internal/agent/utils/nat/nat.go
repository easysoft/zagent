package natHelper

import (
	"errors"
	"fmt"
	consts "github.com/easysoft/zv/internal/comm/const"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	PortStart = 51800
	PortEnd   = 51999

	httpConf = `server{
					listen      %d;
					server_name  %s;
					location / {
						proxy_pass   http://%s:%d;
					}
				}`

	streamConf = `upstream %s {
					server %s:%d;    # 源服务
				}
			
				server {
					listen %d;                # 监听代理主机的端口
					proxy_connect_timeout 1h;
					proxy_timeout 1h;
					proxy_pass %s;        # 转向的服务
				}`
)

func GetValidPort() (ret int, err error) {
	cmd := fmt.Sprintf(`netstat -tln | awk '{print $4}' | grep -o ':51[0-9]\{3\}'`)
	output, _ := _shellUtils.ExeSysCmd(cmd)

	list := strings.Split(output, "\n")

	for p := PortStart; p <= PortEnd; p++ {
		str := ":" + strconv.Itoa(p)
		if !_stringUtils.StrInArr(str, list) {
			ret = p
			break
		}
	}

	if ret == 0 {
		err = errors.New("no port left")
	}

	return
}

func ForwardPort(vmIp string, vmPort int, hostPort int, typ consts.NatForwardType) (err error) {
	confPath, _ := getNginxConf()
	name, pth, _ := getNginxHotLoadingConf(confPath, vmIp, vmPort, typ)

	content := "N/A"
	if typ == consts.Http {
		content = fmt.Sprintf(httpConf, hostPort, name, vmIp, vmPort)
	} else {
		upstreamName := fmt.Sprintf("%s-%d", strings.ReplaceAll(vmIp, ".", "-"), vmPort)
		content = fmt.Sprintf(streamConf, upstreamName, vmIp, vmPort, hostPort, upstreamName)
	}

	_fileUtils.WriteFile(pth, content)

	reloadNginx()

	return
}

func RemoveForward(vmIp string, vmPort int) (err error) {
	confPath, _ := getNginxConf()

	dir := filepath.Dir(filepath.Dir(confPath))
	name := fmt.Sprintf("%s:*", vmIp)
	if vmPort > 0 {
		name = fmt.Sprintf("%s:%d", vmIp, vmPort)
	}

	pth := filepath.Join(dir, "conf.*.d", name)

	cmd := fmt.Sprintf("rm -rf %s", pth)
	_shellUtils.ExeSysCmd(cmd)

	reloadNginx()

	return
}

func getNginxConf() (ret string, err error) {
	cmd := fmt.Sprintf(`nginx -t | grep test`)
	out, err := exec.Command("cmd", "/C", cmd).Output()

	regx, _ := regexp.Compile(`file (.+) test`)
	arr := regx.FindStringSubmatch(string(out))
	ret = arr[1]

	return
}

func getNginxHotLoadingConf(confPath, vmIp string, vmPort int, typ consts.NatForwardType) (
	name, ret string, err error) {
	dir := filepath.Dir(filepath.Dir(confPath))
	name = fmt.Sprintf("%s:%d", vmIp, vmPort)
	ret = filepath.Join(dir, fmt.Sprintf("conf.%s.d", typ), name)

	return
}

func reloadNginx() {
	cmd := fmt.Sprintf(`nginx -s reload`)
	_shellUtils.ExeSysCmd(cmd)
}
