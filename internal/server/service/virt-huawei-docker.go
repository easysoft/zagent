package serverService

import (
	"encoding/json"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"strings"
)

type HuaweiCloudDockerService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	QueueRepo   *repo.QueueRepo   `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService        *VmCommonService `inject:""`
	HistoryService         *HistoryService  `inject:""`
	HuaweiCloudCommService *vendors.HuaweiCloudCommService
	HuaweiCloudCciService  *vendors.HuaweiCloudCciService `inject:""`
}

func (s HuaweiCloudDockerService) CreateRemote(hostId, queueId uint) (result _domain.RpcResp) {
	queue := s.QueueRepo.GetQueue(queueId)
	host := s.HostRepo.Get(hostId)

	client, _ := s.HuaweiCloudCommService.CreateIamClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	token, _ := s.HuaweiCloudCommService.GetIamToken(client)
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	image := queue.DockerImage
	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()
	region := host.CloudRegion
	namespace := host.CloudNamespace

	resp, success := s.HuaweiCloudCciService.Create(image, jobName, cmd, token, region, namespace)
	if success {
		result.Pass("")
	} else {
		bytes, _ := json.Marshal(resp)
		result.Fail(string(bytes))
	}

	return
}

func (s HuaweiCloudDockerService) DestroyRemote(vmId, queueId uint) {
	vm := s.VmRepo.GetById(vmId)
	jobName := vm.CouldInstId

	host := s.HostRepo.Get(vm.HostId)

	client, err := s.HuaweiCloudCommService.CreateIamClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	token, _ := s.HuaweiCloudCommService.GetIamToken(client)

	var status consts.VmStatus
	if err != nil {
		status = consts.VmFailDestroy
	} else {
		s.HuaweiCloudCciService.Destroy(jobName, token, host.CloudRegion, host.CloudNamespace)
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CouldInstId}, status)
	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
