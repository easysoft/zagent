package serverService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type AppiumService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`

	WebSocketService *commonService.WebSocketService `inject:""`
	RpcService       *commonService.RpcService       `inject:""`
	QueueService     *QueueService                   `inject:""`
	HistoryService   *HistoryService                 `inject:""`
}

func NewAppiumService() *AppiumService {
	return &AppiumService{}
}

func (s AppiumService) RunRemote(queue model.Queue) (result _domain.RpcResp) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	build := model.NewAppiumBuildPo(queue, device)
	s.BuildRepo.Save(&build)

	s.HistoryService.Create(consts.Build, build.ID, queue.ID, consts.ProgressCreated, consts.StatusCreated.ToString())

	build = s.BuildRepo.GetBuild(build.ID)
	build.AppiumPort = device.AppiumPort

	result = s.RpcService.AppiumTest(build)

	return
}
