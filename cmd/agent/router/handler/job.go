package handler

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type JobCtrl struct {
	JobService *agentService.JobService `inject:""`
}

func NewJobCtrl() *JobCtrl {
	return &JobCtrl{}
}

func (c *JobCtrl) Add(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	//size := c.JobService.GetTaskSize()
	//if size == 0 {
	c.JobService.AddTask(build)
	reply.Success("Pass to add job.")
	//} else {
	//	reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	//}

	return nil
}
