package hostAgentService

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	downloadUtils "github.com/easysoft/zv/pkg/lib/download"
	"sync"
)

const (
	key = "urls"
)

const (
	keyNotStart   = "not_start"
	keyInProgress = "in_progress"
	keyCompleted  = "completed"
)

var (
	syncMap sync.Map
)

type DownloadService struct {
}

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

func (s *DownloadService) AddTasks(req v1.DownloadReq) (err error) {
	downloadUtils.AddTasks(req.Urls)

	return
}

func (s *DownloadService) CheckTask() (err error) {
	if downloadUtils.IsRunning() {
		return
	}

	downloadUtils.StartTask()
	downloadUtils.CompleteTask()

	return
}
