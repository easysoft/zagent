package serverService

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"time"
)

type DeviceService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`

	QueueService *QueueService `inject:""`
	RpcService   *RpcService   `inject:""`
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (s DeviceService) Register(devices []commDomain.DeviceInst) (result _domain.RpcResp) {
	for _, device := range devices {
		device.LastRegisterDate = time.Now()
		err := s.DeviceRepo.Register(device)

		if err != nil {
			result.Fail(fmt.Sprintf("fail to register device %s ", device.Serial))
			break
		}
	}

	result.Success(fmt.Sprintf("success to register %d devices", len(devices)))
	return
}

func (s DeviceService) IsDeviceReady(device model.Device) bool {
	if device.ID == 0 {
		return false
	}

	deviceStatus := device.DeviceStatus
	appiumStatus := device.DeviceStatus

	registerExpire := time.Now().Unix()-device.LastRegisterDate.Unix() > commConst.RegisterExpireTime*60*1000

	ret := deviceStatus == commConst.DeviceActive && appiumStatus == commConst.DeviceActive && !registerExpire
	return ret
}
