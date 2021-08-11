package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Task struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	BuildType    consts.BuildType `json:"buildType"`
	Priority     int              `json:"priority"`
	Serials      string           `json:"serials"`                                // for appium test
	Environments []Environment    `json:"environments" gorm:"foreignKey:task_id"` // for selenium test

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"scmAddress"`
	ScmAccount  string `json:"scmAccount"`
	ScmPassword string `json:"scmPassword"`

	BrowserType    consts.BrowserType `json:"browserType"`
	BrowserVersion string             `json:"browserVersion"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	Progress consts.BuildProgress `json:"progress"`
	Status   consts.BuildStatus   `json:"status"`

	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	PendingTime *time.Time `json:"pendingTime"`

	UserName string `json:"userName"`
	UserId   uint   `json:"userId"`
	GroupId  uint   `json:"groupId"`

	Queues []Queue `json:"queues" gorm:"-"`

	Histories []History `json:"histories" gorm:"polymorphic:Owner;polymorphicValue:task"`
}

func NewTask() Task {
	plan := Task{
		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,
	}
	return plan
}

func (Task) TableName() string {
	return "biz_task"
}
