package repo

import (
	commDomain "github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
)

type DeviceRepo struct {
	BaseRepo
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
	r.DB.Model(&model.Device{}).Where("serial = ?", device.Serial).First(&po)

	if po.ID == 0 {
		err := r.DB.Model(&model.Device{}).Omit("Ip").Create(&device).Error
		return err
	} else {
		r.DB.Model(&model.Device{}).Updates(device)
		return nil
	}
}

func (r DeviceRepo) GetBySerial(serial string) (device model.Device) {
	r.DB.Model(&model.Device{}).Where("serial=?", serial).First(&device)
	return
}
