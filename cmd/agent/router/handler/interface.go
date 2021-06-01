package handler

import (
	"fmt"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type InterfaceCtrl struct{
	InterfaceService *agentService.InterfaceService `inject:""`
	TaskService *agentService.TaskService `inject:""`
}

func NewInterfaceCtrl() *InterfaceCtrl {
	return &InterfaceCtrl{}
}

func (c *InterfaceCtrl) Exec(ctx context.Context, task _domain.BuildTo, reply *_domain.RpcResult) error {
	size := c.TaskService.GetTaskSize()
	if size == 0 {
		c.TaskService.AddTask(task)
		reply.Success("Success to add job.")
	} else {
		reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	}

	return nil
}
