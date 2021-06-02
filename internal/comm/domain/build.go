package commDomain

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type Build struct {
	QueueId  uint
	ID       uint   `json:"id,omitempty"`
	Priority int    `json:"priority,omitempty"`
	NodeIp   string `json:"nodeIp,omitempty"`
	NodePort int    `json:"nodePort,omitempty"`

	WorkDir    string `json:"workDir,omitempty"`
	ProjectDir string `json:"projectDir,omitempty"`

	StartTime    time.Time `json:"startTime,omitempty"`
	CompleteTime time.Time `json:"completeTime,omitempty"`

	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`

	BuildType     _const.BuildType `json:"buildType,omitempty"`
	TestScenario  `json:"testScenario,omitempty"`
	TestSet       `json:"testSet,omitempty"`
	AutomatedTest `json:"automatedTest,omitempty"`
}
