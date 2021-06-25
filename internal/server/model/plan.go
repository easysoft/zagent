package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Plan struct {
	BaseModel

	Name         string
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

	UserName string
	UserId   uint
	GroupId  uint
}

func NewPlan() Plan {
	plan := Plan{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}
	return plan
}

func (Plan) TableName() string {
	return "biz_plan"
}
