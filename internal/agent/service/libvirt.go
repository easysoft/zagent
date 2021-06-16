package agentService

import "C"
import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"github.com/libvirt/libvirt-go"
	"github.com/smallnest/rpcx/log"
)

const (
	ConnStrLocal  = "qemu:///system"
	ConnStrRemote = "qemu+ssh://192.168.0.56:22/system?keyfile=~/.ssh/id_rsa"
)

var (
	Conn *libvirt.Connect
)

type LibvirtService struct {
}

func NewLibvirtService() *LibvirtService {
	return &LibvirtService{}
}

func (s *LibvirtService) GetDomain() (dom *libvirt.Domain) {
	s.Connect(ConnStrLocal)
	defer func() {
		if res, _ := Conn.Close(); res != 0 {
			_logUtils.Errorf("close() == %d, expected 0", res)
		}
	}()
	names, err := Conn.ListDefinedDomains()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	if len(names) == 0 {
		_logUtils.Errorf("length of domains shouldn't be 0.")
		return
	}
	dom, err = Conn.LookupDomainByName(names[0])
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	info, err := dom.GetInfo()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	log.Infof("%#v", info)

	defer dom.Free()
	return
}

func (s *LibvirtService) Connect(str string) {
	if Conn != nil {
		active, err := Conn.IsAlive()
		if err != nil {
			_logUtils.Errorf(err.Error())
		}
		if active {
			return
		}
	}

	var err error
	Conn, err = libvirt.NewConnect(str)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	active, err := Conn.IsAlive()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	if !active {
		_logUtils.Errorf("not active")
	}
	return
}
