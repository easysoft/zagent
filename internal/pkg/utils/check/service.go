package checkUtils

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	consts "github.com/easysoft/zv/internal/pkg/const"
	natHelper "github.com/easysoft/zv/internal/pkg/utils/net"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	"github.com/libvirt/libvirt-go"
)

func CheckKvm() (status consts.HostServiceStatus, err error) {
	status = consts.HostServiceNotAvailable

	defer func() {
		err1 := recover()
		if err1 != nil {
			err = fmt.Errorf("%v", err1)
		}
	}()

	connStr := "qemu:///system"
	LibvirtConn, err := libvirt.NewConnect(connStr)

	if err == nil {
		active, err := LibvirtConn.IsAlive()
		if err == nil && active {
			status = consts.HostServiceReady
			return status, err
		}
	}

	out, _ := _shellUtils.ExeShell("which libvirtd") // ps -ef | grep libvirt | grep -v grep | grep -v dnsmasq
	if strings.Contains(out, "libvirtd") {
		status = consts.HostServiceNotInstall
		err = errors.New("not installed")
	}

	return
}

func CheckNovnc() (status consts.HostServiceStatus, err error) {
	status = consts.HostServiceNotAvailable

	port, _ := natHelper.GetUsedPortByKeyword("agent", agentConf.Inst.NodePort)
	address := fmt.Sprintf("127.0.0.1:%v/novnc/index.html", port)
	html, err := _httpUtils.Get(address)

	if err != nil {
		_logUtils.Infof("request to %s error: %s", address, err)
	} else if !strings.Contains(string(html), `<title>noVNC</title>`) {
		_logUtils.Infof("request to %s error: %s", address, "html not contain novnc")
	} else {
		status = consts.HostServiceReady
		return
	}

	pth := filepath.Join(consts.NovncDir, "index.html")
	found := _fileUtils.FileExist(pth)

	if !found {
		status = consts.HostServiceNotInstall
	}

	return
}

func CheckWebsockify() (status consts.HostServiceStatus, err error) {
	status = consts.HostServiceNotAvailable

	timeout := time.Second

	port, _ := natHelper.GetUsedPortByKeyword("websockify", agentConf.Inst.WebsockifyPort)
	address := net.JoinHostPort(consts.Localhost, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		_logUtils.Infof("tcp connect to %s error: %s", address, err)
	}

	if conn != nil {
		status = consts.HostServiceReady
		defer conn.Close()
	} else {
		pth := filepath.Join(consts.WebsockifyDir, "run")
		found := _fileUtils.FileExist(pth)

		if !found {
			status = consts.HostServiceNotInstall
		}
	}

	return
}
