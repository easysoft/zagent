package kvmService

import (
	"crypto/rand"
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"github.com/mitchellh/mapstructure"
)

type VmService interface {
	CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp)
	DestroyRemote(vmId, queueId uint) (result _domain.RpcResp)

	genVmName(backing model.VmBacking, vmId uint) (name string)
	genValidMacAddress() (mac string)
	genRandomMac() (mac string)
}

type KvmNativeService struct {
	QueueRepo   *repo.QueueRepo   `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
	HostRepo    *repo.HostRepo    `inject:""`
	IsoRepo     *repo.IsoRepo     `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	TmplRepo    *repo.TmplRepo    `inject:""`

	QueueService   *serverService.QueueService   `inject:""`
	HistoryService *serverService.HistoryService `inject:""`
	RpcService     *commonService.RpcService     `inject:""`
}

func NewKvmService() VmService {
	var service VmService

	if serverConf.Inst.Adapter.VmPlatform == serverConst.KvmNative {
		service = &KvmNativeService{}
	}

	return service
}

func (s KvmNativeService) CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	host := s.HostRepo.Get(hostId)
	backing := s.BackingRepo.Get(backingId)
	sysIso := s.IsoRepo.Get(backing.SysIsoId)
	sysIsoPath := sysIso.Path

	driverIsoPath := ""
	if backing.OsCategory == consts.Windows {
		driverIso := s.IsoRepo.Get(backing.DriverIsoId)
		driverIsoPath = driverIso.Path
	}

	tmpl := s.TmplRepo.Get(tmplId)

	macAddress := s.genValidMacAddress() // get a unique mac address

	vm := model.Vm{
		MacAddress: macAddress,
		BackingId:  backing.ID, BackingPath: backing.Path,
		TmplId: tmpl.ID, TmplName: tmpl.Name,

		OsCategory: backing.OsCategory,
		OsType:     backing.OsType,
		OsVersion:  backing.OsVersion,
		OsLang:     backing.OsLang,

		HostId: host.ID, HostName: host.Name,
		DiskSize: backing.SuggestDiskSize, MemorySize: backing.SuggestMemorySize,
		CdromSys: sysIsoPath, CdromDriver: driverIsoPath,
		Status: consts.VmCreated,
	}

	if backing.SuggestDiskSize == 0 {
		if vm.OsCategory == consts.Windows {
			vm.DiskSize = consts.DiskSizeWindows
		} else if vm.OsCategory == consts.Linux {
			vm.DiskSize = consts.DiskSizeLinux
		} else {
			vm.DiskSize = consts.DiskSizeDefault
		}
	}

	s.VmRepo.Save(&vm) // save vm to db
	vm.Name = s.genVmName(backing, vm.ID)
	s.VmRepo.UpdateVmName(vm)

	kvmReq := model.GenKvmReq(vm)
	result = s.RpcService.CreateVm(host.Ip, host.Port, kvmReq)

	if result.IsSuccess() { // success to create vm
		mp := result.Payload.(map[string]interface{})
		vmInResp := domain.Vm{}
		mapstructure.Decode(mp, &vmInResp)

		s.VmRepo.Launch(vmInResp, vm.ID) // update vm status, mac address
		s.HistoryService.Create(consts.Vm, vm.ID, queueId, "", consts.VmLaunch.ToString())

		s.QueueRepo.UpdateProgressAndVm(queueId, vm.ID, consts.ProgressLaunchVm)
		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressLaunchVm, "")
	} else {
		s.VmRepo.FailToCreate(vm.ID, result.Msg)
		s.QueueService.SaveResult(queueId, consts.ProgressCreateVmFail, consts.StatusFail)
		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressCreateVmFail, consts.StatusFail.ToString())
	}

	return
}

func (s KvmNativeService) DestroyRemote(vmId, queueId uint) (result _domain.RpcResp) {
	vm := s.VmRepo.GetById(vmId)
	host := s.HostRepo.Get(vm.HostId)

	req := domain.KvmReq{VmUniqueName: vm.Name}

	result = s.RpcService.DestroyVm(host.Ip, host.Port, req)
	var status consts.VmStatus
	if result.IsSuccess() {
		status = consts.VmDestroy
	} else {
		status = consts.VmFailDestroy
	}
	s.VmRepo.UpdateStatusByNames([]string{vm.Name}, status)

	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}

func (s KvmNativeService) genVmName(backing model.VmBacking, vmId uint) (name string) {
	name = fmt.Sprintf("test-%s-%s-%s-%d", backing.OsType, backing.OsVersion, backing.OsLang, vmId)

	return
}

func (s KvmNativeService) genValidMacAddress() (mac string) {
	for i := 0; i < 10; i++ {
		mac := s.genRandomMac()
		vm := s.VmRepo.GetByMac(mac)
		if vm.ID == 0 {
			return mac
		}
	}

	return "N/A"
}

func (s KvmNativeService) genRandomMac() (mac string) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	buf[0] |= 2
	mac = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", 0xfa, 0x92, buf[2], buf[3], buf[4], buf[5])
	return
}
