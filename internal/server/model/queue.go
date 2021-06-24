package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Queue struct {
	BaseModel

	Priority int
	Serial   string
	VmId     uint

	BuildType      commConst.BuildType
	OsPlatform     commConst.OsCategory
	OsType         commConst.OsType
	OsLang         commConst.OsLang
	BrowserType    commConst.BrowserType
	BrowserVersion string

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

	Retry int

	TaskName string
	UserName string

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time
	TimeoutTime time.Time

	TaskId  uint
	GroupId uint
}

func NewQueue() Queue {
	queue := Queue{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}
	return queue
}

func (Queue) TableName() string {
	return "biz_queue"
}
