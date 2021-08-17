package serverService

import (
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
)

type HostService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	TmplRepo    *repo.TmplRepo    `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s HostService) GetValidForQueueByVm(queue model.Queue) (hostId, backingId, tmplId uint, found bool) {
	backingIdsByBrowser := s.BackingRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVersion)
	backingIds, found := s.BackingRepo.QueryByOs(queue.OsCategory, queue.OsType, queue.OsLang, backingIdsByBrowser)
	if !found {
		return
	}

	busyHostIds := s.getBusyHosts("vm")
	hostId, backingId = s.HostRepo.QueryByBackings(backingIds, busyHostIds)

	tmplId, found = s.TmplRepo.QueryByOs(queue.OsCategory, queue.OsType, queue.OsLang)

	if hostId == 0 || backingId == 0 || tmplId == 0 {
		found = false
	}

	return
}

func (s HostService) GetValidForQueueByContainer(queue model.Queue) (hostId uint, found bool) {
	busyHostIds := s.getBusyHosts("container")
	hostId = s.HostRepo.QueryUnBusy(busyHostIds)

	if hostId != 0 {
		found = true
	}

	return
}

func (s HostService) getBusyHosts(tp string) (ids []uint) {
	ids = s.HostRepo.QueryBusy(tp)

	return
}
