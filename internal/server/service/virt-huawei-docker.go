package serverService

import (
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/repo"
)

type HuaweiCloudDockerService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService *VmCommonService `inject:""`
	HistoryService  *HistoryService  `inject:""`
}

func (s HuaweiCloudDockerService) CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {

	return
}

func (s HuaweiCloudDockerService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {

	return
}
