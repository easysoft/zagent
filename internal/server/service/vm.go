package serverService

import (
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
)

type VmService interface {
	CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp)
	DestroyRemote(vmId, queueId uint) (result _domain.RpcResp)

	genVmName(backing model.VmBacking, vmId uint) (name string)
}
