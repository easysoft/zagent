package vmHandler

import (
	vmAgentService "github.com/easysoft/zagent/internal/agent-vm/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"golang.org/x/net/context"
)

type JobCtrl struct {
	JobService *vmAgentService.JobService `inject:""`
}

func NewJobCtrl() *JobCtrl {
	return &JobCtrl{}
}

func (c *JobCtrl) Add(ctx context.Context, build commDomain.Build, reply *_domain.RpcResp) error {
	//size := c.JobService.GetTaskSize()
	//if size == 0 {
	c.JobService.AddTask(build)
	reply.Pass("Pass to add job.")
	//} else {
	//	reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	//}

	return nil
}
