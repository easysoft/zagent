package main

import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/libvirt/libvirt-go"
)

func main() {
	connStr := ""

	LibvirtConn, err := libvirt.NewConnect(connStr)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	active, err := LibvirtConn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	if !active {
		_logUtils.Errorf("not active")
	}
}
