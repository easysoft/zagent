package model

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Queue struct {
	BaseModel

	Priority int    `json:"priority"`
	Serial   string `json:"serial"`
	VmId     uint   `json:"vmId"`

	BuildType      commConst.BuildType   `json:"buildType"`
	OsPlatform     commConst.OsCategory  `json:"osPlatform"`
	OsType         commConst.OsType      `json:"osType"`
	OsLang         commConst.OsLang      `json:"osLang"`
	BrowserType    commConst.BrowserType `json:"browserType"`
	BrowserVersion string                `json:"browserVersion"`

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"ScmAddress"`
	ScmAccount  string `json:"ScmAccount"`
	ScmPassword string `json:"ScmPassword"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	Progress commConst.BuildProgress `json:"progress"`
	Status   commConst.BuildStatus   `json:"status"`

	Retry int `json:"retry"`

	TaskName string `json:"taskName"`
	UserName string `json:"userName"`

	StartTime   *time.Time `json:"startTime"`
	PendingTime *time.Time `json:"pendingTime"`
	ResultTime  *time.Time `json:"resultTime"`
	TimeoutTime *time.Time `json:"timeoutTime"`

	TaskId  uint `json:"taskId"`
	GroupId uint `json:"groupId"`
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
	return "biz_queue"
}
