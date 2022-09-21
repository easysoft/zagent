package v1

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/jinzhu/copier"
)

type VmRegisterReq struct {
	Status consts.VmStatus `json:"status" example:"ready"` // Enums consts.VmStatus
	Secret string          `json:"secret" yaml:"secret"`

	MacAddress string `json:"macAddress" example:"1C:1C:1C:24:F4:BF"`
	Ip         string `json:"ip"`
	WorkDir    string `json:"workDir"`
	Port       int    `json:"port" example:"8086"`
}

func (src *VmRegisterReq) ToModel() (po model.Vm) {
	copier.Copy(&po, &src)

	return
}
