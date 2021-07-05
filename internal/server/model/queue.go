package model

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"time"
)

type Queue struct {
	BaseModel

	Priority int    `json:"priority"`
	Serial   string `json:"serial"`
	VmId     uint   `json:"vmId"`

	BuildType      consts.BuildType   `json:"buildType"`
	OsPlatform     consts.OsCategory  `json:"osPlatform"`
	OsType         consts.OsType      `json:"osType"`
	OsLang         consts.OsLang      `json:"osLang"`
	BrowserType    consts.BrowserType `json:"browserType"`
	BrowserVersion string             `json:"browserVersion"`

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"ScmAddress"`
	ScmAccount  string `json:"ScmAccount"`
	ScmPassword string `json:"ScmPassword"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	Progress consts.BuildProgress `json:"progress"`
	Status   consts.BuildStatus   `json:"status"`

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

func NewQueue(buildType consts.BuildType, groupId uint, taskId uint, taskPriority int,
	osCategory consts.OsCategory, osType consts.OsType, osLang consts.OsLang,
	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles bool, taskName string, userName string,
	serial, appUrl, buildCommands, envVars string,
	browserType consts.BrowserType, browserVersion string,
) Queue {

	Task := Queue{
		Progress: consts.ProgressCreated,
		Status:   consts.StatusCreated,

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
		EnvVars:       envVars,

		BrowserType:    browserType,
		BrowserVersion: browserVersion,
	}
	return Task
}

func (Queue) TableName() string {
	return "biz_queue"
}
