package handler

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
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

	if build.BuildType == _const.InterfaceScenario {
		c.InterfaceTestService.ExecScenario(&build, &result)
	} else if build.BuildType == _const.InterfaceSet {
		c.InterfaceTestService.ExecSet(&build, &result)
	}

	reply.Success("Success to exec processor.")
	reply.Payload = result

	return nil
}
