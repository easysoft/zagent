package handler

import (
	"fmt"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ArithCtrl struct{}

func NewArithCtrl() *ArithCtrl {
	return &ArithCtrl{}
}

func (c *ArithCtrl) Add(ctx context.Context, args *_domain.ArithArgs, reply *_domain.ArithReply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}
