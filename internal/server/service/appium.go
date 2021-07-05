package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/mitchellh/mapstructure"
)

type AppiumService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`

	QueueService *QueueService `inject:""`
	RpcService   *RpcService   `inject:""`
}

func NewAppiumService() *AppiumService {
	return &AppiumService{}
}

func (s AppiumService) Start(queue model.Queue) (result _domain.RpcResp) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	build := model.NewBuild(queue.ID, uint(0), queue.BuildType,
		serial, queue.Priority, device.NodeIp, device.NodePort)
	s.BuildRepo.Save(&build)

	build = s.BuildRepo.GetBuild(build.ID)
	build.AppiumPort = device.AppiumPort

	rpcResult := s.RpcService.AppiumTest(build)
	if rpcResult.IsSuccess() {
		s.BuildRepo.Start(build)
	} else {
		s.BuildRepo.Delete(build)
	}

	result.Success("")
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
