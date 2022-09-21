package v1

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/jinzhu/copier"
)

type TaskReq struct {
	ID uint `json:"id" extensions:"!x-omitempty"`

	BuildType consts.BuildType `json:"buildType" example:"selenium"` // Enums consts.BuildType
	Priority  int              `json:"priority"`
	GroupId   uint             `json:"groupId" swaggerignore:"true"`

	IsDockerNative bool `json:"isDockerNative"`

	Serials      string        `json:"serials"`      // for appium test
	Environments []Environment `json:"environments"` // for selenium test

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"scmAddress"`
	ScmAccount  string `json:"scmAccount"`
	ScmPassword string `json:"scmPassword"`

	BrowserType    consts.BrowserType `json:"browserType" example:"chrome"` // Enums consts.BrowserType
	BrowserVersion string             `json:"browserVersion"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`
}

type TaskResp struct {
	Task           model.Task            `json:"task"`
	BuildHistories []domain.BuildHistory `json:"buildHistories"`
}

func (src *TaskReq) ToModel() (po model.Task, err error) {
	copier.Copy(&po, &src)

	return
}
