package commDomain

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"time"
)

type BuildTo struct {
	QueueId   uint
	ID        uint             `json:"id,omitempty"`
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`
	NodeIp    string           `json:"nodeIp,omitempty"`
	NodePort  int              `json:"nodePort,omitempty"`

	WorkDir    string `json:"workDir,omitempty"`
	ProjectDir string `json:"projectDir,omitempty"`
	AppPath    string `json:"appPath,omitempty"`

	ScriptUrl   string `json:"scriptUrl,omitempty"`
	ScmAddress  string `json:"scmAddress,omitempty"`
	ScmAccount  string `json:"scmAccount,omitempty"`
	ScmPassword string `json:"scmPassword,omitempty"`

	AppUrl          string         `json:"appUrl,omitempty"`
	BuildCommands   string         `json:"buildCommands,omitempty"`
	ResultFiles     string         `json:"resultFiles,omitempty"`
	KeepResultFiles _domain.MyBool `json:"keepResultFiles,omitempty"`
	ResultPath      string         `json:"resultPath,omitempty"`
	ResultMsg       string         `json:"resultMsg,omitempty"`

	StartTime    time.Time `json:"startTime,omitempty"`
	CompleteTime time.Time `json:"completeTime,omitempty"`

	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`
}
