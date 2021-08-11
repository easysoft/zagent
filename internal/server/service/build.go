package serverService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	"github.com/easysoft/zagent/internal/server/repo"
	commonService "github.com/easysoft/zagent/internal/server/service/common"
)

type BuildService struct {
	BuildRepo *repo.BuildRepo `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`

	QueueService     *QueueService                   `inject:""`
	HistoryService   *HistoryService                 `inject:""`
	WebSocketService *commonService.WebSocketService `inject:""`
}

func NewBuildService() *BuildService {
	return &BuildService{}
}

func (s BuildService) SaveResult(build domain.Build) (count int) {
	s.BuildRepo.SaveResult(build)

	po := s.BuildRepo.GetBuild(build.ID)
	s.HistoryService.Create(consts.Build, po.ID, po.QueueId, po.Progress, po.Status.ToString())

	s.QueueService.SaveResult(po.QueueId, po.Progress, po.Status)

	queue := s.QueueRepo.GetQueue(po.QueueId)
	s.WebSocketService.UpdateTask(queue.TaskId, "CheckAndCall Test")

	return
}
