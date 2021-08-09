package serverService

import (
	"crypto/rand"
	"fmt"
	consts "github.com/easysoft/zagent/internal/comm/const"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type VmCommonService struct {
	VmRepo           *repo.VmRepo                    `inject:""`
	HostRepo         *repo.HostRepo                  `inject:""`
	BackingRepo      *repo.BackingRepo               `inject:""`
	QueueRepo        *repo.QueueRepo                 `inject:""`
	RpcService       *commonService.RpcService       `inject:""`
	HistoryService   *HistoryService                 `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
	QueueService     *QueueService                   `inject:""`
}

func NewKvmService() VmService {
	var service VmService

	if serverConf.Inst.Adapter.VmPlatform == consts.KvmNative {
		service = &KvmNativeService{}
	} else if serverConf.Inst.Adapter.VmPlatform == consts.HuaweiCloud {
		service = &HuaweiCloudService{}
	}

	return service
}

func (s VmCommonService) SaveVmCreationResult(isSuccess bool, result string, queueId uint, vmId uint,
	vncAddress, imagePath, backingPath string) {
	queue := s.QueueRepo.GetQueue(queueId)
	if isSuccess { // success to create vm
		s.VmRepo.Launch(vncAddress, imagePath, backingPath, vmId) // update vm status, mac address
		s.HistoryService.Create(consts.Vm, vmId, queueId, "", consts.VmLaunch.ToString())

		s.QueueRepo.LaunchVm(queueId)
		s.QueueRepo.UpdateVm(queueId, vmId)

		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressLaunchVm, "")
		s.WebSocketService.UpdateTask(queue.TaskId, "success to create vm")
	} else {
		s.VmRepo.FailToCreate(vmId, result)
		s.QueueService.SaveResult(queueId, consts.ProgressCreateVmFail, consts.StatusFail)

		s.HistoryService.Create(consts.Queue, queueId, queueId, consts.ProgressCreateVmFail, consts.StatusFail.ToString())
		s.WebSocketService.UpdateTask(queue.TaskId, "fail to create vm")
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
