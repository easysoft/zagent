package testing

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	serverService "github.com/easysoft/zagent/internal/server/service"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type UnitService struct {
	BuildRepo *repo.BuildRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`

	RpcService       *commonService.RpcService       `inject:""`
	QueueService     *serverService.QueueService     `inject:""`
	HistoryService   *serverService.HistoryService   `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewUnitService() *UnitService {
	return &UnitService{}
}

func (s UnitService) RemoteRun(queue model.Queue, host model.Host) (result _domain.RpcResp) {
	build := model.NewUnitBuildPo(queue, host)
	s.BuildRepo.Save(&build)

	s.HistoryService.Create(consts.Build, build.ID, queue.ID, consts.ProgressCreated, consts.StatusCreated.ToString())

	build = s.BuildRepo.GetBuild(build.ID)

	result = s.RpcService.UnitTest(build)

	return
}
