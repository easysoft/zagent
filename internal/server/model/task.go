package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Task struct {
	BaseModel

	Priority     int
	Serials      string // for appium test
	Environments string // for selenium test
	BuildType    commConst.BuildType

	ScriptUrl   string
	ScmAddress  string
	ScmAccount  string
	ScmPassword string

	AppUrl          string
	BuildCommands   string
	ResultFiles     string
	KeepResultFiles bool

	Progress commConst.BuildProgress
	Status   commConst.BuildStatus

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time

	TaskName string
	UserName string

	GroupId uint
}

func NewTask() Task {
	task := Task{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}
	return task
}

func (Task) TableName() string {
	return "biz_task"
}
