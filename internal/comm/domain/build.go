package commDomain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
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

	Progress commConst.BuildProgress `json:"progress,omitempty"`
	Status   commConst.BuildStatus   `json:"status,omitempty"`

	BuildType     commConst.BuildType `json:"buildType,omitempty"`
	TestScenario  TestScenario        `json:"testScenario,omitempty"`
	TestSet       TestSet             `json:"testSet,omitempty"`
	AutomatedTest AutomatedTest       `json:"automatedTest,omitempty"`
}
