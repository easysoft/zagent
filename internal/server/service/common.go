package serverService

import (
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
)

type CommonService struct {
	VmCommonService *VmCommonService `inject:""`
	HistoryService  *HistoryService  `inject:""`
}

func (s CommonService) ReturnErr(result *_domain.RpcResp, err error, queueId, vmId uint) {
	result.Fail(err.Error())
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vmId, "", "", "")
	return
}
