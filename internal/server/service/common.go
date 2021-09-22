package serverService

type CommonService struct {
	VmCommonService *VmCommonService `inject:""`
	HistoryService  *HistoryService  `inject:""`
}

//func (s CommonService) ReturnVmCreateErr(result *_domain.RpcResp, err error, queueId, vmId uint) {
//	result.Fail(err.Error())
//	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vmId, "", "", "")
//	return
//}

//func (s CommonService) ReturnVmDestroyErr(result *_domain.RpcResp, err error, queueId, vmId uint) {
//	result.Fail(err.Error())
//	s.VmCommonService.SaveVmDestroyResult(result.IsSuccess(), result.Msg, queueId, vmId)
//	return
//}
