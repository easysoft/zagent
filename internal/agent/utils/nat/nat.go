package natHelper

import (
	"errors"
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
	cmd := fmt.Sprintf(`netstat -tln | awk '{print $4}' | grep -o ':51[0-9]\{3\}'`)
	output, err := _shellUtils.ExeSysCmd(cmd)
	if err != nil {
		return
	}

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

func ForwardPort(vmIp string, vmPort int, hostIp string, hostPort int) (err error) {
	/**
	sudo iptables -A INPUT -p tcp --dport 54000 -j ACCEPT
	sudo iptables -t nat -A PREROUTING -d 192.168.0.56 -p tcp -m tcp --dport 54000 -j DNAT --to-destination 192.168.122.79:4000
	sudo iptables -t nat -A POSTROUTING -s 192.168.122.0/255.255.255.0 -d 192.168.122.79 -p tcp -m tcp --dport 4000 -j SNAT --to-source 192.168.122.1

	sudo iptables -t nat  -L -n --line-number
	sudo iptables -D FORWARD 14 -t filter
	sudo ufw disable && sudo ufw enable
	sudo sysctl -a | grep forward

	sudo sysctl -w net.ipv4.ip_forward=1
	sudo modprobe br_netfilter
	sudo sysctl -w net.bridge.bridge-nf-call-iptables=1
	net.bridge.bridge-nf-call-iptables = 1
	sudo lsmod |grep br_netfilter
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
