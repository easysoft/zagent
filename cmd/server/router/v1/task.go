package v1

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/jinzhu/copier"
)

type TaskReq struct {
	ID uint `json:"id" extensions:"!x-omitempty"`

	BuildType consts.BuildType `json:"buildType" example:"consts.BuildType"`
	Priority  int              `json:"priority"`
	GroupId   uint             `json:"groupId" swaggerignore:"true"`

	IsDockerNative bool `json:"isDockerNative"`

	Serials      string        `json:"serials"`      // for appium test
	Environments []Environment `json:"environments"` // for selenium test

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"scmAddress"`
	ScmAccount  string `json:"scmAccount"`
	ScmPassword string `json:"scmPassword"`

	BrowserType    consts.BrowserType `json:"browserType" example:"consts.BrowserType"`
	BrowserVersion string             `json:"browserVersion"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`
}

type TaskResp struct {
	Task           model.Task     `json:"data"`
	BuildHistories []BuildHistory `json:"buildHistories"`
}

func (src *TaskReq) ToModel() (po model.Task, err error) {
	copier.Copy(&po, &src)

	return
}
