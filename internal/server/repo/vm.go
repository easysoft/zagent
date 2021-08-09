package repo

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type VmRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewVmRepo() *VmRepo {
	return &VmRepo{}
}

func (r VmRepo) Register(vm model.Vm) (err error) {
	// just update status by mac for exist vm
	r.DB.Model(&model.Vm{}).Where("mac_address=?", vm.MacAddress).
		Updates(
			map[string]interface{}{"status": vm.Status, "work_dir": vm.WorkDir,
				"node_ip": vm.NodeIp, "node_port": vm.NodePort,
				"last_register_time": time.Now()})

	return
}

func (r VmRepo) GetById(id uint) (vm model.Vm) {
	r.DB.Model(&model.Vm{}).Where("ID=?", id).First(&vm)
	return
}
func (r VmRepo) GetByMac(mac string) (vm model.Vm) {
	r.DB.Model(&model.Vm{}).Where("mac=?", mac).First(&vm)
	return
}

func (r VmRepo) Save(po *model.Vm) {
	r.DB.Model(&model.Vm{}).Omit("").Create(po)
	return
}

func (r VmRepo) UpdateVmName(vm model.Vm) {
	r.DB.Model(&model.Vm{}).Where("id=?", vm.ID).Update("name", vm.Name)
}
func (r VmRepo) UpdateVmCloudInst(vm model.Vm) {
	r.DB.Model(&model.Vm{}).Where("id=?", vm.ID).
		Updates(map[string]interface{}{
			"could_inst_id": vm.CouldInstId,
			"node_ip":       vm.NodeIp,
			"mac_address":   vm.MacAddress,
		})
}

func (r VmRepo) Launch(vncAddress, imagePath, backingPath string, id uint) {
	r.DB.Model(&model.Vm{}).Where("id=?", id).
		Updates(
			map[string]interface{}{
				"status":       consts.VmLaunch,
				"vnc_port":     vncAddress,
				"image_path":   imagePath,
				"backing_path": backingPath,
			})

	return
}

func (r VmRepo) UpdateStatusByNames(vms []string, status consts.VmStatus) {
	db := r.DB.Model(&model.Vm{}).Where("name IN (?)", vms)

	if status == consts.VmRunning {
		db.Where("status != ?", consts.VmReady) // not to update active vm status
	}

	db.Updates(map[string]interface{}{"status": status})
}

func (r VmRepo) UpdateStatusByCloudInstId(ids []string, status consts.VmStatus) {
	db := r.DB.Model(&model.Vm{}).Where("could_inst_id IN (?)", ids)

	if status == consts.VmRunning {
		db.Where("status != ?", consts.VmReady) // not to update active vm status
	}

	db.Updates(map[string]interface{}{"status": status})
}

func (r VmRepo) DestroyMissedVmsStatus(vms []string, hostId uint) {
	db := r.DB.Model(&model.Vm{}).
		Where("host_id=? AND status!=? "+
			" AND strftime('%s','now') - strftime('%s',created_at) > ?",
			hostId, consts.VmDestroy, consts.AgentCheckInterval)

	if len(vms) > 0 {
		db.Where("name NOT IN (?)", vms)
	}

	db.Updates(map[string]interface{}{"status": consts.VmDestroy})
}

func (r VmRepo) FailToCreate(id uint, msg string) {
	r.DB.Model(&model.Vm{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"status": consts.VmFailCreate, "desc": msg})
}
