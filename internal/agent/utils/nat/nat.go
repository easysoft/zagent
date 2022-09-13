package natHelper

import (
	"fmt"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"strconv"
	"strings"
)

const (
	PortStart = 51800
	PortEnd   = 51999
)

func GetValidPort() (ret int, err error) {
	cmd := fmt.Sprintf(`netstat -anp | awk '{print $7}' | grep -o '51[0-9]\{3\}'`)
	output, err := _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}
	arr := strings.Split(output, "\n")

	for p := 51800; p <= 51999; p++ {
		str := strconv.Itoa(p)
		if !_stringUtils.StrInArr(str, arr) {
			ret = p
			break
		}
	}

	return
}

func ForwardPort(vmIp string, vmPort int, hostIp string, hostPort int) (err error) {
	/**
	sudo iptables -A INPUT -p tcp --dport 58086 -j ACCEPT
	sudo iptables -t nat -A PREROUTING -d 192.168.0.56 -p tcp -m tcp --dport 58086 -j DNAT --to-destination 192.168.122.7:8086
	sudo iptables -t nat -A POSTROUTING -s 192.168.122.0/255.255.255.0 -d 192.168.122.7 -p tcp -m tcp --dport 8086 -j SNAT --to-source 192.168.122.1

	sudo iptables -nL -v --line-numbers -t filter
	sudo iptables -D FORWARD 14 -t filter
	sudo ufw disable && sudo ufw enable
	*/

	cmd := fmt.Sprintf(`sudo iptables -A INPUT -p tcp --dport %d -j ACCEPT`, hostPort)
	output, err := _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	//
	cmd = fmt.Sprintf(`sudo iptables -t nat -A PREROUTING -d %s`+
		` -p tcp -m tcp --dport %d -j DNAT --to-destination %s:%d`,
		hostIp, hostPort, vmIp, vmPort)
	output, err = _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	//
	cmd = fmt.Sprintf(`sudo iptables -t nat -A POSTROUTING`+
		` -s 192.168.122.0/255.255.255.0`+
		` -d %s -p tcp -m tcp --dport %d -j SNAT --to-source 192.168.122.1`,
		vmIp, vmPort)
	output, err = _shellUtils.ExeSysCmd(cmd)
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

	// sudo iptables -t nat -L LIBVIRT_FWI -n --line-number | grep %s | awk '{print $1}'
	cmd = fmt.Sprintf(`sudo iptables -t nat -L LIBVIRT_FWI -n --line-number | grep %s | awk '{print $1}'`, vmIp)
	output, err = _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

	arr := strings.Split(output, "\n")

	for _, item := range arr {
		// sudo iptables -nvL
		// iptables -t nat -L LIBVIRT_FWI -D INPUT %s
		cmd = fmt.Sprintf(`iptables -t nat -L LIBVIRT_FWI -D INPUT %s`, item)
		output, err = _shellUtils.ExeSysCmd(cmd)
	}

	return
}
