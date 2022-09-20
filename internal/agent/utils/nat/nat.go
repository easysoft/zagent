package natHelper

import (
	"errors"
	"fmt"
	consts "github.com/easysoft/zv/internal/comm/const"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
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
	cmd := fmt.Sprintf(`sudo nginx -t | grep test`)
	out, err := _shellUtils.ExeSysCmd(cmd)

	regx, _ := regexp.Compile(`file (.+) test`)
	arr := regx.FindStringSubmatch(out)
	confPath := arr[1]

	dir := filepath.Dir(filepath.Dir(confPath))
	name := fmt.Sprintf("%d:%d@%s", hostPort, vmPort, vmIp)
	pth := filepath.Join(dir, fmt.Sprintf("conf.%s.d", typ), name)

	content := "N/A"
	if typ == consts.Http {
		content = fmt.Sprintf(httpConf, hostPort, name, vmIp, vmPort)
	} else {
		upstreamName := fmt.Sprintf("%s-%d", strings.ReplaceAll(vmIp, ".", "-"), vmPort)
		content = fmt.Sprintf(streamConf, upstreamName, vmIp, vmPort, hostPort, upstreamName)
	}

	_fileUtils.WriteFile(pth, content)

	// reload nginx
	cmd = fmt.Sprintf(`sudo nginx -s reload`)
	output, err := _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	_logUtils.Info(output)

	return
}

func RemoveForwardByVm(vmIp string, vmPort int, hostIp string, hostPort int) (err error) {
	// sudo iptables -t nat -D INPUT -p tcp --dport 58086 -j ACCEPT
	cmd := fmt.Sprintf(`sudo iptables -t nat -D INPUT -p tcp --dport %d -j ACCEPT`, hostPort)
	output, err := _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	_logUtils.Info(output)

	// sudo iptables -t nat  -L -n --line-number | grep 192.168.122.79 | awk '{print $1}'
	cmd = fmt.Sprintf(`sudo iptables -t nat  -L -n --line-number | grep %s | awk '{print $1}'`, vmIp)
	output, err = _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	arr := strings.Split(output, "\n")

	for _, item := range arr {
		// sudo iptables -nvL
		// sudo iptables -t nat -D POSTROUTING 2
		// sudo iptables -t filter -D LIBVIRT_FWI 2
		// sudo iptables -t filter -D LIBVIRT_FWO 2

		cmd = fmt.Sprintf(`iptables -t nat -D INPUT %s`, item)
		output, err = _shellUtils.ExeSysCmd(cmd)
	}

	return
}
