package serverService

import (
	"fmt"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	_domain "github.com/easysoft/zv/pkg/domain"
	"time"
)

type DeviceService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (s DeviceService) Register(devices []domain.DeviceInst) (result _domain.RemoteResp) {
	for _, device := range devices {
		device.LastRegisterDate = time.Now()
		err := s.DeviceRepo.Register(device)

		if err != nil {
			result.Fail(fmt.Sprintf("fail to register device %s ", device.Serial))
			break
		}
	}

	result.Pass(fmt.Sprintf("success to register %d devices", len(devices)))
	return
}

func (s DeviceService) IsDeviceReady(device model.Device) bool {
	if device.ID == 0 {
		return false
	}

	deviceStatus := device.DeviceStatus
	appiumStatus := device.DeviceStatus

	registerExpire := time.Now().Unix()-device.LastRegisterDate.Unix() > consts.DeviceRegisterExpireTime

	ret := deviceStatus == consts.DeviceActive && appiumStatus == consts.DeviceActive && !registerExpire
	return ret
}
