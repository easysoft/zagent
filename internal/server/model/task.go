package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Task struct {
	BaseModel

	Name         string              `json:"name"`
	Priority     int                 `json:"priority"`
	Serials      string              `json:"serials"`      // for appium test
	Environments string              `json:"environments"` // for selenium test
	BuildType    commConst.BuildType `json:"buildType"`

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"scmAddress"`
	ScmAccount  string `json:"scmAccount"`
	ScmPassword string `json:"scmPassword"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	Progress commConst.BuildProgress `json:"progress"`
	Status   commConst.BuildStatus   `json:"status"`

	StartTime   time.Time `json:"startTime"`
	PendingTime time.Time `json:"pendingTime"`
	ResultTime  time.Time `json:"resultTime"`

	UserName string `json:"userName"`
	UserId   uint   `json:"userId"`
	GroupId  uint   `json:"groupId"`
}

func NewTask() Task {
	plan := Task{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}
	return plan
}

func (Task) TableName() string {
	return "biz_task"
}
