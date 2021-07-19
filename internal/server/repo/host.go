package repo

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
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
	host.Status = consts.HostActive

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
	list := make([]commDomain.VmHost, 0)

	sql := fmt.Sprintf(`SELECT r.host_id, r.vm_backing_id
			FROM biz_host_backing_r r 
		    INNER JOIN biz_host host on r.host_id = host.id 

	        WHERE host.status = 'active' 
			AND r.vm_backing_id IN (%s) AND host.id NOT IN (%s) LIMIT 1`,

		strings.Join(_commonUtils.UintToStrArr(backingIds), ","),
		strings.Join(_commonUtils.UintToStrArr(busyHostIds), ","))

	r.DB.Raw(sql).Find(&list)

	for _, id := range backingIds { // get the most fittest one by backing order
		for _, item := range list {
			if id == item.HostId {
				hostId = item.HostId
				backingId = item.VmBackingId

				return
			}
		}
	}

	return
}

func (r HostRepo) QueryBusy() (hostIds []uint) {
	hosts := make([]HostResult, 0)
	r.DB.Raw(`SELECT host.id host_id, host.max_vm_num num
					FROM biz_host host
					WHERE NOT host.deleted AND NOT host.disabled`).
		Scan(&hosts)

	for _, host := range hosts {
		items := make([]HostResult, 0)

		r.DB.Raw(`SELECT vm.host_id host_id, COUNT(vm.id) num
					FROM biz_vm vm
					WHERE vm.status == ? AND vm.host_id = ?
					GROUP BY host_id
					HAVING num > ?
					ORDER BY num`,
			consts.VmRunning, host.HostId, host.Num).
			Scan(&items)

		for _, item := range items {
			hostId := item.HostId
			hostIds = append(hostIds, hostId)
		}
	}

	return
}

func (r HostRepo) QueryUnBusy(busyHostIds []uint) (hostId uint) {
	list := make([]commDomain.VmHost, 0)

	sql := fmt.Sprintf(`SELECT id
			FROM biz_host host
	        WHERE host.status = 'active' 
			AND host.id NOT IN (%s) LIMIT 1`,
		strings.Join(_commonUtils.UintToStrArr(busyHostIds), ","))

	r.DB.Raw(sql).Find(&list)

	return
}

type HostResult struct {
	HostId uint
	Num    int
}
