package testing

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
	"github.com/mitchellh/mapstructure"
)

type AppiumService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`

	QueueService *serverService.QueueService `inject:""`
	RpcService   *commonService.RpcService   `inject:""`
}

func NewAppiumService() *AppiumService {
	return &AppiumService{}
}

func (s AppiumService) Start(queue model.Queue) (result _domain.RpcResp) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	build := model.NewAppiumBuildPo(queue, device)
	s.BuildRepo.Save(&build)

	build = s.BuildRepo.GetBuild(build.ID)
	build.AppiumPort = device.AppiumPort

	result = s.RpcService.AppiumTest(build)
	if result.IsSuccess() {
		s.BuildRepo.Start(build)
	} else {
		s.BuildRepo.Delete(build)
	}

	return
}

func (s AppiumService) SaveResult(buildResult _domain.RpcResp, resultPath string) {
	appiumTestTo := domain.Build{}
	mp := buildResult.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &appiumTestTo)

	progress := consts.ProgressCompleted
	var status consts.BuildStatus
	if buildResult.IsSuccess() {
		status = consts.StatusPass
	} else {
		status = consts.StatusFail
	}

	s.BuildRepo.SaveResult(appiumTestTo, resultPath, progress, status, buildResult.Msg)
	s.QueueService.SetQueueResult(appiumTestTo.QueueId, progress, status)
	if progress == consts.ProgressTimeout {
		s.BuildRepo.SetTimeoutByQueueId(appiumTestTo.QueueId)
	}
}
