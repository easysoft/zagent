package handler

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"golang.org/x/net/context"
)

type TaskCtrl struct {
	TaskService          *agentService.TaskService          `inject:""`
	InterfaceTestService *agentService.InterfaceTestService `inject:""`
}

func NewTaskCtrl() *TaskCtrl {
	return &TaskCtrl{}
}

func (c *TaskCtrl) Add(ctx context.Context, task commDomain.Build, reply *_domain.RpcResp) error {
	size := c.TaskService.GetTaskSize()
	if size == 0 {
		c.TaskService.AddTask(task)
		reply.Success("Success to add job.")
	} else {
		reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	}

	return nil
}

func (c *TaskCtrl) Exec(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	_logUtils.Infof("%v", build)

	result := commDomain.TestResult{}

	if build.BuildType == commConst.InterfaceScenario {
		result = c.InterfaceTestService.ExecScenario(&build)
	} else if build.BuildType == commConst.InterfaceSet {
		c.InterfaceTestService.ExecSet(&build, &result)
	}

	reply.Success("Success to exec processor.")
	reply.Payload = result

	return nil
}
