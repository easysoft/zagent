package serverService

import (
	"fmt"
	"github.com/easysoft/zv/internal/comm/const"
	virtualboxapi "github.com/easysoft/zv/internal/pkg/vendors/virtualbox/api"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	"github.com/easysoft/zv/internal/server/service/common"
	_domain "github.com/easysoft/zv/pkg/domain"
)

type VmWareCloudVmService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	CommonService   *CommonService               `inject:""`
	VmCommonService *VmCommonService             `inject:""`
	RpcService      *commonService.RemoteService `inject:""`
	HistoryService  *HistoryService              `inject:""`
}

func (s VmWareCloudVmService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	host := s.HostRepo.Get(hostId)
	backing := s.BackingRepo.Get(backingId)
	backing.Name = s.VmCommonService.genTmplName(backing)

	vm := model.Vm{
		HostId: host.ID, HostName: host.Name,
		Status:     consts.VmCreated,
		OsCategory: backing.OsCategory,
		OsType:     backing.OsType,
		OsVersion:  backing.OsVersion,
		OsLang:     backing.OsLang,
		BackingId:  backing.ID,
	}
	s.VmRepo.Save(&vm) // save vm to db, then update name with id
	vm.Name = s.VmCommonService.genVmName(backing, vm.ID)
	s.VmRepo.UpdateVmName(vm)

	req := model.GenVmWareReq(vm.Name, backing.Name, "",
		backing.SuggestCpuCount, backing.SuggestMemorySize,
		host.CloudIamUser, host.CloudIamPassword)
	result = s.RpcService.CreateVmWare(host.Ip, host.Port, req)

	// save to db
	result.Pass("")
	s.VmRepo.UpdateVmCloudInst(vm)

	//vm.VncPort, _ = huaweiCloudService.QueryVnc(vm.CloudInstId, ecsClient)
	s.VmCommonService.SaveVmCreationResult(result.IsSuccess(), result.Msg, queueId, vm.ID, vm.VncPort, vm.VncUrl, "", "")

	return
}

func (s VmWareCloudVmService) DestroyRemote(vmId, queueId uint) (result _domain.RemoteResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	status := consts.VmDestroy

	req := model.GenVmWareReq("", "", vm.CloudInstId,
		0, 0,
		host.CloudIamUser, host.CloudIamPassword)
	result = s.RpcService.DestroyVmWare(host.Ip, host.Port, req)
	if !result.IsSuccess() {
		status = consts.VmDestroyFail
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CloudInstId}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}

func (s VmWareCloudVmService) CreateClient(ip string, port int, account, password string) (
	client *virtualboxapi.VirtualBox, err error) {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client = virtualboxapi.NewVirtualBox(account, password, url, false, "")

	err = client.Logon()

	return
}
