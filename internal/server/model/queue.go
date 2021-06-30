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

	Retry int `json:"retry" gorm:"default:0"`

	TaskName string `json:"taskName"`
	UserName string `json:"userName"`

	StartTime   *time.Time `json:"startTime"`
	PendingTime *time.Time `json:"pendingTime"`
	ResultTime  *time.Time `json:"resultTime"`
	TimeoutTime *time.Time `json:"timeoutTime"`

	TaskId  uint `json:"taskId"`
	GroupId uint `json:"groupId"`
}

func NewQueue(buildType commConst.BuildType, groupId uint, taskId uint, taskPriority int,
	osCategory commConst.OsCategory, osType commConst.OsType, osLang commConst.OsLang,
	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles bool, taskName string, userName string,
	serial string, appUrl string, buildCommands string) Queue {

	Task := Queue{
		Progress: commConst.ProgressCreated,
		Status:   commConst.StatusCreated,

		BuildType:  buildType,
		OsLang:     osLang,
		OsPlatform: osCategory,
		OsType:     osType,

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

		Serial:        serial,
		AppUrl:        appUrl,
		BuildCommands: buildCommands,
	}
	return Task
}

func (Queue) TableName() string {
	return "biz_queue"
}
