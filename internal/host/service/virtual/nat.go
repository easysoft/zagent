package virtualService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	natHelper "github.com/easysoft/zv/internal/pkg/utils/net"
	"github.com/jinzhu/copier"
	"sync"
)

type NatService struct {
	syncMap sync.Map
}

func NewNatService() *NatService {
	srv := NatService{}

	return &srv
}

func (s *NatService) AddVmPortMap(req v1.VmPortMapReq) (resp v1.VmPortMapResp, err error) {
	copier.CopyWithOption(&resp, req, copier.Option{DeepCopy: true})

	resp.HostPort, resp.AlreadyMapped, _ = natHelper.ForwardPortIfNeeded(req.VmIp, req.VmPort, req.Type)

	return
}

func (s *NatService) RemoveVmPortMap(req v1.VmPortMapReq) (err error) {
	err = natHelper.RemoveForward(req.VmIp, req.VmPort, req.Type)

	return
}
