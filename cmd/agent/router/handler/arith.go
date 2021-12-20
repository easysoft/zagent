package handler

import (
	"fmt"
	agentService "github.com/easysoft/zv/internal/agent/service"
	_domain "github.com/easysoft/zv/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ArithCtrl struct {
	TaskService *agentService.JobService `inject:""`
}

func NewArithCtrl() *ArithCtrl {
	return &ArithCtrl{}
}

func (c *ArithCtrl) Add(ctx context.Context, args *_domain.ArithArgs, reply *_domain.ArithReply) error {
	size := c.TaskService.GetTaskSize()
	fmt.Printf("size: %d\n", size)

	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}
