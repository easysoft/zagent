package serverService

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	serverUitls "github.com/easysoft/zv/internal/server/utils/lib"
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

	busyHostIds := s.getBusyHosts()
	hostId, backingId = s.HostRepo.QueryByBackings(backingIds, busyHostIds)

	platform := s.HostRepo.Get(hostId).Capabilities.ToString()

	if serverUitls.IsCloud(platform) {
		tmplId = 1 // not use tmpl for cloud
	} else {
		tmplId, found = s.TmplRepo.QueryByOs(queue.OsCategory, queue.OsType, queue.OsLang)
	}

	if hostId == 0 || backingId == 0 || tmplId == 0 {
		found = false
	}

	return
}

func (s HostService) GetValidForQueueByDocker(queue model.Queue) (hostId uint, found bool) {
	busyHostIds := s.getBusyHosts()

	isNative := queue.DockerImage == ""
	hostId = s.HostRepo.QueryUnBusy(busyHostIds, consts.PlatformDocker, isNative)

	if hostId != 0 {
		found = true
	}

	return
}

func (s HostService) getBusyHosts() (ids []uint) {
	ids = s.HostRepo.QueryBusy()

	return
}
