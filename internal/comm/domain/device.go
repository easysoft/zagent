package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type DeviceInst struct {
	DeviceSpec

	NodeIp           string                  `json:"nodeIp"`
	NodePort         int                     `json:"rpcPort"`
	AppiumPort       int                     `json:"appiumPort"`
	DeviceStatus     commConst.DeviceStatus  `json:"deviceStatus"`
	AppiumStatus     commConst.ServiceStatus `json:"appiumStatus"`
	LastRegisterDate time.Time               `json:"lastRegisterDate"`
}

type DeviceSpec struct {
	Serial           string `json:"serial"`
	Model            string `json:"model"`
	ApiLevel         int    `json:"apiLevel"`
	Version          string `json:"version"`
	Code             string `json:"code"`
	Os               string `json:"os"`
	Kernel           string `json:"kernel"`
	Ram              int    `json:"ram"`
	Rom              int    `json:"rom"`
	Cpu              string `json:"cpu"`
	Battery          int    `json:"battery"`
	Density          int    `json:"density"`
	DeviceIp         string `json:"deviceIp"`
	ResolutionHeight int    `json:"resolutionHeight"`
	ResolutionWidth  int    `json:"resolutionWidth"`
}