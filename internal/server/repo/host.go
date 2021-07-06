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

func (r HostRepo) QueryBusy() (ret []map[string]uint) {
	hosts := make([]map[string]uint, 0)
	r.DB.Raw(`SELECT host.id hostId, host.max_vm_num maxVmNum
					FROM biz_host
					AND NOT deleted AND NOT disabled`).
		Scan(&hosts)

	for _, host := range hosts {
		//items := make([]map[string]uint, 0)
		r.DB.Raw(`SELECT host_id hostId, COUNT(id) vmCount
					FROM biz_vm
					WHERE status == ? AND host_id = ?
					GROUP BY hostId
					HAVING vmCount > ?
					ORDER BY vmCount`,
			consts.VmRunning, host["hostId"], host["maxVmNum"]-1).
			Scan(&ret)

		//for _, item := range items {
		//	if item["vmCount"] >= consts.MaxVmOnHost {
		//		ret = append(ret, item)
		//	}
		//}
	}

	return
}
