package kvmService

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

func (s HostService) GetValidForQueue(queue model.Queue) (hostId, backingId, tmplId uint, found bool) {
	backingIdsByBrowser := s.BackingRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVersion)
	backingIds, found := s.BackingRepo.QueryByOs(queue.OsCategory, queue.OsType, queue.OsLang, backingIdsByBrowser)
	if !found {
		return
	}

	busyHostIds := s.getBusyHosts()
	hostId, backingId = s.HostRepo.QueryByBackings(backingIds, busyHostIds)

	tmplId, found = s.TmplRepo.QueryByOs(queue.OsCategory, queue.OsType, queue.OsLang)

	return
}

func (s HostService) getBusyHosts() (ids []uint) {
	ids = s.HostRepo.QueryBusy()

	return
}
