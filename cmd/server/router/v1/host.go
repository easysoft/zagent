package v1

import (
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/jinzhu/copier"
)

type HostRegisterReq struct {
	Status consts.HostStatus `json:"status" example:"online"` // Enums consts.HostStatus

	Vms []VmInHostReq `json:"vms"`
}

type VmInHostReq struct {
	Name   string          `json:"name"`
	Status consts.VmStatus `json:"status" example:"running"` // Enums consts.VmStatus
}

func (src *HostRegisterReq) ToModel() (po model.Host) {
	copier.Copy(&po, &src)

	if po.Name == "" {
		po.Name = po.Ip
	}

	for _, v := range src.Vms {
		vm := model.Vm{}
		copier.Copy(&vm, &v)

		po.Vms = append(po.Vms, vm)
	}

	return
}
