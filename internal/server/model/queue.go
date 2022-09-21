package model

import (
	"github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type Queue struct {
	BaseModel

	Priority int    `json:"priority"`
	Serial   string `json:"serial"`
	VmId     uint   `json:"vmId"`

	BuildType      consts.BuildType   `json:"buildType" example:"selenium"` // Enums consts.BuildType
	OsCategory     consts.OsCategory  `json:"osCategory" example:"windows"` // Enums consts.OsCategory
	OsType         consts.OsType      `json:"osType" example:"win10"`       // Enums consts.OsType
	OsLang         consts.OsLang      `json:"osLang" example:"zh_cn"`       // Enums consts.OsLang
	BrowserType    consts.BrowserType `json:"browserType example:"chrome"`  // Enums consts.BrowserType
	BrowserVersion string             `json:"browserVersion"`
	DockerImage    string             `json:"dockerImage"`

	ScriptUrl   string `json:"scriptUrl"`
	ScmAddress  string `json:"ScmAddress"`
	ScmAccount  string `json:"ScmAccount"`
	ScmPassword string `json:"ScmPassword"`

	AppUrl          string `json:"appUrl"`
	BuildCommands   string `json:"buildCommands"`
	EnvVars         string `json:"envVars"`
	ResultFiles     string `json:"resultFiles"`
	KeepResultFiles bool   `json:"keepResultFiles"`

	Progress consts.BuildProgress `json:"progress" example:"created"` // Enums consts.BuildProgress
	Status   consts.BuildStatus   `json:"status" example:"pass"`      // Enums consts.BuildStatus

	Retry int `json:"retry" gorm:"default:0"`

	TaskName string `json:"taskName"`
	UserName string `json:"userName"`

	ResPendingTime  *time.Time `json:"resPendingTime"`
	ResLaunchedTime *time.Time `json:"resLaunchedTime"`
	RunTime         *time.Time `json:"runTime"`
	ResultTime      *time.Time `json:"resultTime"`
	TimeoutTime     *time.Time `json:"timeoutTime"`

	TaskId  uint `json:"taskId"`
	GroupId uint `json:"groupId"`

	Histories []History `json:"histories" gorm:"polymorphic:Owner;polymorphicValue:queue"`
}

func NewQueue(buildType consts.BuildType, groupId uint, taskId uint, taskPriority int,
	osCategory consts.OsCategory, osType consts.OsType, osLang consts.OsLang, dockerImage string,
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
		OsCategory: osCategory,
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
		DockerImage:    dockerImage,
	}
	return Task
}

func (Queue) TableName() string {
	return "biz_queue"
}
