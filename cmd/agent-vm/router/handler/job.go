package vmHandler

import (
	agentService "github.com/easysoft/zagent/internal/agent/service"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type JobCtrl struct {
	JobService *agentService.JobService `inject:""`
}

func NewJobCtrl() *JobCtrl {
	return &JobCtrl{}
}

func (c *JobCtrl) Add(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	build := commDomain.Build{}
	if err := ctx.ReadJSON(&build); err != nil {
		ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	//size := c.JobService.GetTaskSize()
	//if size == 0 {
	c.JobService.AddTask(build)
	//} else {
	//	reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	//}

	ctx.JSON(_httpUtils.ApiRes(iris.StatusOK, "success to add job", nil))
	return
}
