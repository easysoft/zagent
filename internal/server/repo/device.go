package repo

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
)

type DeviceRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewDeviceRepo() *DeviceRepo {
	return &DeviceRepo{}
}

func (r DeviceRepo) Register(device commDomain.DeviceInst) (err error) {
	code := 1
	tx := r.DB.Begin()
	defer Defer(tx, &code)

	var po model.Device
	r.DB.Where("serial = ?", device.Serial).First(&po)

	if po.ID == 0 {
		err := r.DB.Model(&device).Omit("Ip").Create(&device).Error
		return err
	} else {
		r.DB.Model(&device).Updates(device)
		return nil
	}
}

func (r DeviceRepo) GetBySerial(serial string) (device model.Device) {
	r.DB.Where("serial=?", serial).First(&device)
	return
}
