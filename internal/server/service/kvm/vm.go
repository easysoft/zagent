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
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"github.com/mitchellh/mapstructure"
)

type VmService interface {
	Register(vm domain.Vm) (result _domain.RpcResp)
	CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp)
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

	RpcService *commonService.RpcService `inject:""`
}

func NewKvmService() VmService {
	var service VmService

	if serverConf.Config.Adapter.VmPlatform == serverConst.KvmNative {
		service = &KvmNativeService{}
	}

	return service
}

func (s KvmNativeService) Register(vm domain.Vm) (result _domain.RpcResp) {
	err := s.VmRepo.Register(vm)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", vm.MacAddress))
	}
	return
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
		s.QueueRepo.UpdateProgressAndVm(queueId, vm.ID, consts.ProgressLaunchVm)
	} else {
		s.VmRepo.FailToCreate(vm.ID, result.Msg)
		s.QueueRepo.SetQueueStatus(queueId, consts.ProgressCreateVmFail, consts.StatusFail)
	}

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
