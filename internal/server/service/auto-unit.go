package serverService

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	_domain "github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	commonService "github.com/easysoft/zv/internal/server/service/common"
)

type UnitService struct {
	BuildRepo *repo.BuildRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`

	QueueService     *QueueService                   `inject:""`
	RpcService       *commonService.RpcService       `inject:""`
	FacadeService    *FacadeService                  `inject:""`
	HistoryService   *HistoryService                 `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewUnitService() *UnitService {
	return &UnitService{}
}

func (s UnitService) RunRemote(queue model.Queue, host model.Host) (result _domain.RpcResp) {
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
