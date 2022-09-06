package serverService

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	commonService "github.com/easysoft/zv/internal/server/service/common"
	_domain "github.com/easysoft/zv/pkg/domain"
)

type UnitService struct {
	BuildRepo *repo.BuildRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`

	QueueService     *QueueService                   `inject:""`
	RpcService       *commonService.RemoteService    `inject:""`
	FacadeService    *FacadeService                  `inject:""`
	HistoryService   *HistoryService                 `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewUnitService() *UnitService {
	return &UnitService{}
}

func (s UnitService) RunRemote(queue model.Queue, host model.Host) (result _domain.RemoteResp) {
	build := model.NewUnitBuildPo(queue, host)
	s.BuildRepo.Save(&build)

	s.HistoryService.Create(consts.Build, build.ID, queue.ID, consts.ProgressCreated, consts.StatusCreated.ToString())

	build = s.BuildRepo.GetBuild(build.ID)

	if queue.DockerImage == "" {
		result = s.RpcService.UnitTest(build)
	} else {
		result = s.FacadeService.Create(host.ID, 0, 0, queue.ID)
	}

	return
}
