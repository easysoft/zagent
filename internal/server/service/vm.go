package serverService

import (
	"crypto/rand"
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type VmService interface {
	CreateRemote(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp)
	DestroyRemote(vmId, queueId uint) (result _domain.RpcResp)
	genVmName(backing model.VmBacking, vmId uint) (name string)
}

type VmCommonService struct {
	HostRepo         *repo.HostRepo                  `inject:""`
	QueueRepo        *repo.QueueRepo                 `inject:""`
	VmRepo           *repo.VmRepo                    `inject:""`
	HistoryService   *HistoryService                 `inject:""`
	QueueService     *QueueService                   `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func (s VmCommonService) SaveVmCreationResult(isSuccess bool, result string, queueId uint, vmId uint,
	vncAddress, imagePath, backingPath string) {
	if isSuccess { // success to create vm
		s.VmRepo.Launch(vncAddress, imagePath, backingPath, vmId) // update vm status, mac address
		s.HistoryService.Create(consts.Vm, vmId, queueId, "", consts.VmLaunch.ToString())

		s.QueueRepo.ResLaunched(queueId)
		s.QueueRepo.UpdateVm(queueId, vmId)

		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressResLaunched, "")
	} else {
		s.VmRepo.FailToCreate(vmId, result)
		s.QueueService.SaveResult(queueId, consts.ProgressResFailed, consts.StatusFail)

		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressResFailed, consts.StatusFail.ToString())
	}

	return
}

func (s VmCommonService) genVmName(backing model.VmBacking, vmId uint) (name string) {
	name = fmt.Sprintf("test-%s-%s-%s-%d", backing.OsType, backing.OsVersion, backing.OsLang, vmId)

	return
}

func (s VmCommonService) genValidMacAddress() (mac string) {
	for i := 0; i < 10; i++ {
		mac := s.genRandomMac()
		vm := s.VmRepo.GetByMac(mac)
		if vm.ID == 0 {
			return mac
		}
	}

	return "N/A"
}

func (s VmCommonService) genRandomMac() (mac string) {
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
