package handler

import (
	"fmt"
	vmAgentService "github.com/easysoft/zagent/internal/agent-vm/service"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ArithCtrl struct {
	TaskService *vmAgentService.JobService `inject:""`
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
