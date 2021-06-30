package repo

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
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

func (r HostRepo) Register(host commDomain.Host) (po model.Host, err error) {
	err = r.DB.Where("ip = ?", host.Ip).First(&po).Error
	if err != gorm.ErrRecordNotFound {
		err = r.DB.Model(&host).Omit("").Create(&host).Error
		return
	} else {
		err = r.DB.Model(&host).Updates(host).Error
		return
	}
}

func (r HostRepo) Get(id uint) (host model.Host) {
	r.DB.Where("id=?", id).First(&host)
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
	list := make([]map[string]uint, 0)
	r.DB.Raw(`SELECT host_id, COUNT(id) num
					FROM biz_vm
					WHERE status != 'destroy' AND NOT deleted AND NOT disabled
					GROUP BY host_id
					ORDER BY num`).
		Scan(&list)

	for _, item := range list {
		if item["num"] > commConst.MaxVmOnHost {
			ret = append(ret, item)
		}
	}

	return
}
