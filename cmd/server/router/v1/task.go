package v1

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/jinzhu/copier"
)

type Task struct {
	BuildType consts.BuildType `json:"buildType" enums:"ztf,selenium,appium,unittest,interface"`
	Priority  int              `json:"priority"`
	GroupId   uint             `json:"groupId" swaggerignore:"true"`

	IsDockerNative bool `json:"isDockerNative"`

	Serials      string        `json:"serials"`                                // for appium test
	Environments []Environment `json:"environments" gorm:"foreignKey:task_id"` // for selenium test

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"scmAddress"`
	ScmAccount  string `json:"scmAccount"`
	ScmPassword string `json:"scmPassword"`

	BrowserType    consts.BrowserType `json:"browserType" enums:"chrome,firefox,edge,ie"`
	BrowserVersion string             `json:"browserVersion"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`
}

func (src *Task) ToModel() (po model.Task, err error) {
	copier.Copy(&po, &src)

	return
}
