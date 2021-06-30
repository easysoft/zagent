package service

import (
	"crypto/rand"
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	serverConf "github.com/easysoft/zagent/internal/server/cfg"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"strings"
)

type VmService interface {
	Register(vm commDomain.Vm) (result _domain.RpcResp)
	CreateRemote(hostId, backingId, queueId uint) (result _domain.RpcResp)
	genVmName(imageName string) (name string)
	genValidMacAddress() (mac string)
	genRandomMac() (mac string)
}

type KvmNativeService struct {
	QueueRepo   *repo.QueueRepo   `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
	HostRepo    *repo.HostRepo    `inject:""`
	IsoRepo     *repo.IsoRepo     `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`

	RpcService *RpcService `inject:""`
}

func NewKvmService() VmService {
	var service VmService

	if serverConf.Config.Adapter.VmPlatform == serverConst.KvmNative {
		service = &KvmNativeService{}
	}

	return service
}

func (s KvmNativeService) Register(vm commDomain.Vm) (result _domain.RpcResp) {
	err := s.VmRepo.Register(vm)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", vm.MacAddress))
	}
	return
}

func (s KvmNativeService) CreateRemote(hostId, backingId, queueId uint) (result _domain.RpcResp) {
	host := s.HostRepo.Get(hostId)
	backingImage := s.BackingRepo.Get(backingId)
	sysIso := s.IsoRepo.Get(backingImage.SysIsoId)
	sysIsoPath := sysIso.Path

	driverIsoPath := ""
	if backingImage.OsCategory == commConst.Windows {
		driverIso := s.IsoRepo.Get(backingImage.DriverIsoId)
		driverIsoPath = driverIso.Path
	}

	mac := s.genValidMacAddress() // get a unique mac address
	vmName := s.genVmName(backingImage.Name)

	vm := model.Vm{MacAddress: mac, Name: vmName,
		HostId: host.ID, BackingId: backingId,
		DiskSize: backingImage.SuggestDiskSize, MemorySize: backingImage.SuggestMemorySize,
		CdromSys: sysIsoPath, CdromDriver: driverIsoPath, Backing: backingImage.Path}

	s.VmRepo.Save(&vm) // save vm to db

	kvmReq := model.GenKvmReq(vm)
	result = s.RpcService.CreateVm(host.Ip, host.Port, kvmReq)

	if result.IsSuccess() { // success to create vm
		vmInResp := result.Payload.(commDomain.Vm)
		s.VmRepo.Launch(vmInResp) // update vm status, mac address

		s.QueueRepo.UpdateVm(queueId, vm.ID, commConst.ProgressLaunchVm)
	} else {
		s.VmRepo.FailToCreate(vm.ID, result.Msg)

		s.QueueRepo.Pending(queueId)
	}

	return
}

func (s KvmNativeService) genVmName(imageName string) (name string) {
	uuid := strings.Replace(_stringUtils.NewUuid(), "-", "", -1)
	name = strings.Replace(imageName, "backing", uuid, -1)

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
	mac = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return
}
