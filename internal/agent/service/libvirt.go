package agentService

import "C"
import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"github.com/libvirt/libvirt-go"
)

const (
	ConnStrLocal = "qemu:///system"
)

var (
	ConnLocal *libvirt.Connect
)

type LibvirtService struct {
}

func NewLibvirtService() *LibvirtService {
	return &LibvirtService{}
}

func (s *LibvirtService) GetDomain() (dom *libvirt.Domain) {
	conn := s.GetConn(ConnStrLocal)
	defer func() {
		if res, _ := conn.Close(); res != 0 {
			_logUtils.Errorf("Close() == %d, expected 0", res)
		}
	}()
	ids, err := conn.ListDomains()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	if len(ids) == 0 {
		_logUtils.Errorf("Length of ListDomains shouldn't be zero")
		return
	}
	dom, err = conn.LookupDomainById(ids[0])
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	info, err := dom.GetInfo()
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	_logUtils.Infof("%#v", info)

	defer dom.Free()

	return
}

func (s *LibvirtService) GetConn(str string) *libvirt.Connect {
	if ConnLocal != nil {
		active, err := ConnLocal.IsAlive()
		if err != nil {
			_logUtils.Errorf(err.Error())
		}
		if active {
			return ConnLocal
		}
	}

	ConnLocal, err := libvirt.NewConnect(str)
	if err != nil {
		_logUtils.Errorf(err.Error())
	}

	active, err := ConnLocal.IsAlive()
	if active {
		return ConnLocal
	}

	return nil
}
