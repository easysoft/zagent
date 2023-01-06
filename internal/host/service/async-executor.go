package hostAgentService

import (
	agentModel "github.com/easysoft/zagent/internal/host/model"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"time"
)

type AsyncExecutorService struct {
}

func NewAsyncExecutor() *SnapService {
	s := SnapService{}
	s.TimeStamp = time.Now().Unix()

	return &s
}

func (s *AsyncExecutorService) Exec(po *agentModel.Task, uuidStr string, opt func(*agentModel.Task) (consts.TaskStatus, string)) (
	status consts.TaskStatus, result string) {

	var statusMsg consts.TaskStatus
	var resultMsg string

	ch := make(chan string)

	go func() {
		statusMsg, resultMsg = opt(po)

		ch <- statusMsg.ToString()
	}()

	select {
	case val := <-ch:
		status = consts.TaskStatus(val)

	case <-time.After(consts.CreateSnapTimeout * time.Second):
		status = consts.Timeout
		if uuidStr != "" {
			_shellUtils.KillProcessByUUID(uuidStr)
		}
	}

	result = resultMsg

	return
}
