package testing

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	rpcService "github.com/easysoft/zagent/internal/server/service/rpc"
	"github.com/mitchellh/mapstructure"
)

type SeleniumService struct {
	BuildRepo *repo.BuildRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`

	RpcService   *rpcService.RpcService      `inject:""`
	QueueService *serverService.QueueService `inject:""`
}

func NewSeleniumService() *SeleniumService {
	return &SeleniumService{}
}

func (s SeleniumService) Start(queue model.Queue) (result _domain.RpcResp) {
	vmId := queue.VmId
	vm := s.VmRepo.GetById(vmId)

	build := model.NewBuild(queue.ID, vmId, queue.BuildType,
		"", queue.Priority, vm.PublicIp, vm.PublicPort)
	s.BuildRepo.Save(&build)

	build = s.BuildRepo.GetBuild(build.ID)

	rpcResult := s.RpcService.SeleniumTest(build)
	if rpcResult.IsSuccess() {
		s.BuildRepo.Start(build)
	} else {
		s.BuildRepo.Delete(build)
	}

	result.Success("")
	return
}

func (s SeleniumService) SaveResult(buildResult _domain.RpcResp, resultPath string) {
	seleniumTestTo := domain.Build{}
	mp := buildResult.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &seleniumTestTo)

	progress := consts.ProgressCompleted
	var result consts.BuildStatus
	if buildResult.IsSuccess() {
		result = consts.StatusPass
	} else {
		result = consts.StatusFail
	}

	s.BuildRepo.SaveResult(seleniumTestTo, resultPath, progress, result, buildResult.Msg)
	s.QueueService.SetQueueResult(seleniumTestTo.QueueId, progress, result)
	if progress == consts.ProgressTimeout {
		s.BuildRepo.SetTimeoutByQueueId(seleniumTestTo.QueueId)
	}
}
