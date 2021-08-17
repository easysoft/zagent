package repo

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	domain "github.com/easysoft/zagent/internal/comm/domain"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"strings"
)

type HostRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewHostRepo() *HostRepo {
	return &HostRepo{}
}

func (r HostRepo) Register(host model.Host) (po model.Host, err error) {
	err = r.DB.Model(&model.Host{}).Where("ip = ?", host.Ip).First(&po).Error
	if err == gorm.ErrRecordNotFound {
		err = r.DB.Model(&model.Host{}).Omit("").Create(&host).Error
		return
	} else {
		err = r.DB.Model(&model.Host{}).Where("ip = ?", host.Ip).Updates(host).Error
		return
	}
}

func (r HostRepo) Get(id uint) (host model.Host) {
	r.DB.Model(&model.Host{}).Where("id=?", id).First(&host)
	return
}

func (r HostRepo) QueryByBackings(backingIds []uint, busyHostIds []uint) (hostId, backingId uint) {
	vmHostList := make([]domain.VmHost, 0)

	sql := fmt.Sprintf(`SELECT r.host_id, r.vm_backing_id
			FROM biz_host_backing_r r 
		    INNER JOIN biz_host host on r.host_id = host.id 

	        WHERE host.status = '%s' 
			AND r.vm_backing_id IN (%s) AND host.id NOT IN (%s) 
			ORDER BY host.priority
            LIMIT 1`,

		consts.HostReady,
		strings.Join(_commonUtils.UintToStrArr(backingIds), ","),
		strings.Join(_commonUtils.UintToStrArr(busyHostIds), ","))

	r.DB.Raw(sql).Find(&vmHostList)

	for _, id := range backingIds { // get the most fittest one by backing order
		for _, item := range vmHostList {
			if id == item.VmBackingId {
				hostId = item.HostId
				backingId = item.VmBackingId

				return
			}
		}
	}

	return
}

func (r HostRepo) QueryBusy(tp string) (hostIds []uint) {
	hosts := make([]HostResult, 0)
	r.DB.Raw(fmt.Sprintf(
		`SELECT host.id host_id, host.max_vm_num max_num, host.vm_platform vm_platform
					FROM biz_host host
					WHERE host.status = '%s' AND host AND NOT host.deleted AND NOT host.disabled
					ORDER BY host.priority`,
		consts.HostReady)).
		Scan(&hosts)

	for _, host := range hosts {
		if strings.Index(host.Platform.ToString(), "_cloud") > -1 {
			continue
		}

		items := make([]uint, 0)

		r.DB.Raw(`SELECT vm.id
					FROM biz_vm vm
					WHERE vm.status <> ? AND vm.host_id = ?
					ORDER BY num`,
			consts.VmDestroy, host.HostId).
			Scan(&items)

		if len(items) >= host.MaxNum {
			hostIds = append(hostIds, host.HostId)
		}
	}

	return
}

func (r HostRepo) QueryUnBusy(busyHostIds []uint) (hostId uint) {
	list := make([]model.Host, 0)

	whr := r.DB.Model(&model.Host{}).Where(
		"status = ?",
		consts.HostReady).Find(&list)

	if busyHostIds != nil {
		whr.Where(
			"id NOT IN (?)",
			busyHostIds)
	}

	if len(list) > 0 {
		hostId = list[0].ID
	}

	return
}

type HostResult struct {
	HostId   uint
	MaxNum   int
	Platform consts.Platform
}
