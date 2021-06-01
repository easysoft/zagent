package _domain

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type BuildTo struct {
	WorkDir    string
	ProjectDir string
	AppPath    string

	ID           uint
	Priority     int
	ComputerIp   string
	ComputerPort int

	BuildType _const.BuildType

	QueueId uint

	ScriptUrl   string
	ScmAddress  string
	ScmAccount  string
	ScmPassword string

	AppUrl          string
	BuildCommands   string
	ResultFiles     string
	KeepResultFiles MyBool
	ResultPath      string
	ResultMsg       string

	StartTime    time.Time
	CompleteTime time.Time

	Progress _const.BuildProgress
	Status   _const.BuildStatus
}
