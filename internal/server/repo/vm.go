package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type VmRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewVmRepo() *VmRepo {
	return &VmRepo{}
}

func (r VmRepo) Register(vm commDomain.Vm) (err error) {
	// just update status by mac for exist vm
	r.DB.Model(&vm).Where("mac=?", vm.MacAddress).
		Updates(
			map[string]interface{}{"status": vm.Status, "ip": vm.PublicIp, "port": vm.PublicPort, "workDir": vm.WorkDir,
				"lastRegisterDate": time.Now(), "updatedAt": time.Now()})

	return
}

func (r VmRepo) GetById(id uint) (vm model.Vm) {
	r.DB.Where("ID=?", id).First(&vm)
	return
}
func (r VmRepo) GetByMac(mac string) (vm model.Vm) {
	r.DB.Where("mac=?", mac).First(&vm)
	return
}

func (r VmRepo) Save(po model.Vm) {
	r.DB.Model(&po).Omit("").Create(&po)
	return
}

func (r VmRepo) Launch(vm commDomain.Vm) {
	r.DB.Model(&vm).Where("id=?", vm.Id).
		Updates(
			map[string]interface{}{"status": "launch", "imagePath": vm.ImagePath,
				"defPath": vm.DefPath, "updatedAt": time.Now()})

	return
}

func (r VmRepo) UpdateStatusByNames(vms []string, status commConst.VmStatus) {
	db := r.DB.Model(&model.Vm{}).Where("name = IN (?)", vms)

	if status == "running" {
		db.Where("AND status != 'active'")
	}

	db.Updates(map[string]interface{}{"status": status, "updatedAt": time.Now()})
}

func (r VmRepo) DestroyMissedVmsStatus(vms []string, hostId uint) {
	db := r.DB.Model(&model.Vm{}).Where("hostId=? AND status!=?", hostId, "destroy")

	if len(vms) > 0 {
		db.Where("AND name NOT IN (?)", vms)
	}

	db.Updates(map[string]interface{}{"status": "destroy", "updatedAt": time.Now()})
}

func (r VmRepo) FailToCreate(id uint, msg string) {
	r.DB.Model(&model.Vm{}).
		Where("id=?", id).
		Updates(map[string]interface{}{"msg": commConst.VmFailToCreate, "updatedAt": time.Now()})
}
