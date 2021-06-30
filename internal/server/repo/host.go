package repo

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/jinzhu/gorm"
	"strconv"
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

func (r HostRepo) QueryByBackings(backingIds []int, hostIds []int) (hostId, backingId int) {
	list := make([]map[string]int, 0)

	sql := fmt.Sprintf(`SELECT r.host_id hostId, r.vm_backing_id backingId
			FROM biz_host_backing_r r 
		    INNER JOIN biz_host host on r.host_id = host.id 

	        WHERE host.status = 'active' 
			AND r.vm_backing_id IN (%s) AND host.id IN (%s) LIMIT 1`,

		strings.Join(_commonUtils.IntToStrArr(backingIds), ","),
		strings.Join(_commonUtils.IntToStrArr(hostIds), ","))

	r.DB.Raw(sql).Find(&list)

	for _, id := range backingIds { // get the most fittest one by backing order
		for _, item := range list {
			if id == item["hostId"] {
				hostId = item["hostId"]
				backingId = item["backingId"]
			}
		}
	}

	return
}

func (r HostRepo) QueryIdle(host int) (ret []map[string]int) {
	sql := `SELECT temp.hostId, temp.vmCount 
			FROM (
				SELECT DISTINCT host.id hostId, IFNULL(sub.num, 0) vmCount
				FROM BizHost host
				LEFT JOIN
					(SELECT hostId, COUNT(1) num
						FROM BizVm
						WHERE status != 'destroy' AND NOT deleted AND NOT disabled
						GROUP BY hostId) sub
					ON host.id = sub.hostId
			) temp
			WHERE temp.vmCount <= ` + strconv.Itoa(commConst.MaxVmOnHost)

	r.DB.Raw(sql).Scan(&ret)
	return
}
