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
	Task := Queue{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,
	}
	return Task
}
func NewTaskDetail(serial string, buildType commConst.BuildType, groupId uint, taskId uint, taskPriority int,
	osPlatform commConst.OsCategory, osType commConst.OsType,
	osLang commConst.OsLang, browserType commConst.BrowserType, browserVersion string,
	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles bool, taskName string, userName string,
	appUrl string, buildCommands string) Queue {
	Task := Queue{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,

		Serial:      serial,
		BuildType:   buildType,
		OsLang:      osLang,
		OsPlatform:  osPlatform,
		OsType:      osType,
		BrowserType: browserType,

		GroupId:         groupId,
		TaskId:          taskId,
		Priority:        taskPriority,
		ScriptUrl:       scriptUrl,
		ScmAddress:      scmAddress,
		ScmAccount:      scmAccount,
		ScmPassword:     scmPassword,
		ResultFiles:     resultFiles,
		KeepResultFiles: keepResultFiles,
		TaskName:        taskName,
		UserName:        userName,

		AppUrl:        appUrl,
		BuildCommands: buildCommands,
	}
	return Task
}

func (Queue) TableName() string {
	return "biz_Task"
}
